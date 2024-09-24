package core

import (
	"time"

	"github.com/fxamacker/cbor/v2"
)

type dispatchKafkaMessage struct {
	Id            uint64
	CampaignId    uint64
	RecipienId    uint64
	PhoneNumber   string
	RecipientName string
	ScheduleTime  time.Time
	Message       string
}

func (m dispatchKafkaMessage) marshal() ([]byte, error) {
	return cbor.Marshal(m)
}
