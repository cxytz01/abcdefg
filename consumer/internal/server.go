package internal

import (
	"context"
	"os"
	"time"

	"consumer/pkg/models"
	"consumer/pkg/utils"

	log "github.com/sirupsen/logrus"
)

type Task struct {
	Ctx context.Context
	Cf  utils.TomlConfig
}

func NewServer(ctx context.Context) *Task {
	return &Task{Ctx: ctx}
}

func (s *Task) initconfig() {
	if err := utils.InitConfFile(Flags.Config, &s.Cf); err != nil {
		log.Fatalf("Load configuration file err: %v", err)
	}

	utils.EngineGroup = utils.NewEngineGroup(s.Ctx, &[]utils.EngineInfo{{utils.DBConsumer, s.Cf.Consumer.Pg, nil}})
	utils.KafkaReader = utils.NewKafkaReader(s.Ctx, s.Cf.Consumer.Kafka)
}

func (s *Task) Start() {
	s.initconfig()

	for {
		if message, err := utils.KafkaReader.ReadMessage(context.Background()); err != nil {
			log.Errorf("kafka ReadMessage err: %v", err)
			break
		} else {
			var dkm dispatchKafkaMessage
			if err := dkm.unmarshal(message.Value); err != nil {
				log.Errorf("invalid message: [topic: %v, partition: %v, offset: %v, key: %v], unmarshal err.", message.Topic, message.Partition, message.Offset, message.Key)
				continue
			}

			// send to thirdparty

			mm := models.Messages{
				Id:            dkm.Id,
				DeliveredTime: time.Now(),
			}

			log.Infof("PID: %v fetch message: [partition: %v, id: %v, campaignId: %v, recipientId: %v, phoneNumber: %v]", os.Getpid(), message.Partition, dkm.Id, dkm.CampaignId, dkm.RecipienId, dkm.PhoneNumber)
			if _, err := utils.EngineGroup[utils.DBConsumer].Where("id = ?", dkm.Id).Update(&mm); err != nil {
				panic(err)
			}
		}
	}
}
