package core

import (
	"bytes"
	"context"
	"encoding/csv"
	"html/template"
	"net/http"
	"producer/pkg/models"
	"producer/pkg/utils"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

func DispatchToKafka(ctx context.Context) *utils.BuErrorResponse {
	campaigns := make([]models.Campaigns, 0)
	if err := utils.EngineGroup[utils.DBProducer].Where("enable = ? and schedule_state = ? and schedule_time <= ?", true, 0, time.Now()).Asc("schedule_time").Find(&campaigns); err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}

	for _, campaign := range campaigns {
		messages := make([]models.Messages, 0)
		if err := utils.EngineGroup[utils.DBProducer].Where("campaign_id = ?", campaign.Id).Find(&messages); err != nil {
			return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
		}

		log.Infof("push campaign [id: %v] to kafka", campaign.Id)

		// kafka
		kafkaMessages := make([]kafka.Message, len(messages))
		for index, content := range messages {
			dkm := dispatchKafkaMessage{
				Id:            content.Id,
				CampaignId:    content.CampaignId,
				RecipienId:    content.RecipienId,
				PhoneNumber:   content.PhoneNumber,
				RecipientName: content.RecipientName,
				ScheduleTime:  content.ScheduleTime,
				Message:       content.Message,
			}
			v, _ := dkm.marshal()
			kafkaMessages = append(kafkaMessages, kafka.Message{Key: []byte(strconv.Itoa(index)), Value: v})
		}
		if err := utils.KafkaWriter.WriteMessages(ctx, kafkaMessages...); err != nil {
			log.Errorf("dispatch to kafka err: %v", err)
			return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
		}

		log.Infof("%v messages were successfully pushed to kafka", len(messages))

		campaign.ScheduleState = 1
		if _, err := utils.EngineGroup[utils.DBProducer].Where("id = ?", campaign.Id).Update(&campaign); err != nil {
			return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
		}
	}

	return nil
}

func CreateCampaign(ctx context.Context, r *CreateCampaignParam, f *FileUploaded) (interface{}, *utils.BuErrorResponse) {
	campaign := models.Campaigns{
		MessageTemplate: r.MessageTemplate,
		Enable:          true,
		ScheduleTime:    r.ScheduleTime,
		ScheduleState:   0,
	}

	var cs CampaignSession
	cs.init()
	defer cs.close()

	if err := cs.insertCampaign(&campaign, f.filePath); err != nil {
		return nil, &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}

	if err := cs.insertRecipientsAndMessages(f.content, r.MessageTemplate, campaign.Id, r.ScheduleTime); err != nil {
		return nil, &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}
	cs.commit()

	f.store()

	return &campaign, nil
}

type CampaignSession struct {
	s *xorm.Session
}

func (s *CampaignSession) init() {
	s.s = utils.EngineGroup[utils.DBProducer].NewSession()
	if err := s.s.Begin(); err != nil {
		panic(err)
	}
}

func (s *CampaignSession) commit() {
	s.s.Commit()
}

func (s *CampaignSession) close() {
	s.s.Close()
}

func (s *CampaignSession) insertCampaign(campaign *models.Campaigns, filePath string) error {
	if _, err := s.s.Insert(campaign); err != nil {
		log.Errorf("insert campaign, execute sql error: %v", err)
		return err
	}

	recipientCSVPath := models.RecipientCSVPath{
		CampaignId: campaign.Id,
		CSVPath:    filePath,
	}

	if _, err := s.s.Insert(&recipientCSVPath); err != nil {
		log.Errorf("insert campaign, execute sql error: %v", err)
		return err
	}

	return nil
}

func (s *CampaignSession) insertRecipientsAndMessages(csvbyte []byte, messageTemplate string, campaignID uint64, scheduleTime time.Time) error {
	tmpl, err := template.New("check").Parse(messageTemplate)
	if err != nil {
		log.Errorf("template syntax checkout error: %v", err)
		return err
	}

	reader := csv.NewReader(bytes.NewReader(csvbyte))

	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("failed to parse CSV: %v", err)
		return err
	}

	for _, record := range records[1:] {
		data := map[string]interface{}{
			"phone": record[0],
			"name":  record[1],
		}

		var output bytes.Buffer
		if err = tmpl.Execute(&output, data); err != nil {
			log.Errorf("tmpl execute err: %v", err)
			return err
		}

		recipient := models.Recipients{
			CampaignId:    campaignID,
			PhoneNumber:   record[0],
			RecipientName: record[1],
			ScheduleTime:  scheduleTime,
		}

		if _, err := s.s.Insert(&recipient); err != nil {
			log.Errorf("insert recipient, execute sql error: %v", err)
			return err
		}

		message := models.Messages{
			CampaignId:    campaignID,
			RecipienId:    recipient.Id,
			PhoneNumber:   record[0],
			RecipientName: record[1],
			ScheduleTime:  scheduleTime,
			Message:       output.String(),
		}

		if _, err := s.s.Insert(&message); err != nil {
			log.Errorf("insert message, execute sql error: %v", err)
			return err
		}
	}

	return nil
}
