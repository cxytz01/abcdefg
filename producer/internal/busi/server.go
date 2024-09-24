package busi

import (
	"context"
	"producer/pkg/models"
	"producer/pkg/utils"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Ctx context.Context
	Cf  utils.TomlConfig
}

func NewServer(ctx context.Context) *Server {
	return &Server{Ctx: ctx}
}

func (s *Server) initconfig() {
	if err := utils.InitConfFile(Flags.Config, &s.Cf); err != nil {
		log.Fatalf("Load configuration file err: %v", err)
	}

	utils.EngineGroup = utils.NewEngineGroup(s.Ctx, &[]utils.EngineInfo{{utils.DBProducer, s.Cf.Producer.Pg, models.Tables}})
	utils.KafkaWriter = utils.NewKafkaWriter(s.Ctx, s.Cf.Producer.Kafka)

}

func (s *Server) setLogTimeformat() {
	timeFormater := new(log.TextFormatter)
	timeFormater.FullTimestamp = true
	log.SetFormatter(timeFormater)
}

func (s *Server) Start() {
	s.initconfig()
	s.setLogTimeformat()

	HttpServerStart(s.Cf.Producer.Addr, s.Cf.Producer.CSVStore)
}
