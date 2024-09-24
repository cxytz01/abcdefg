package models

import (
	"time"
)

type Messages struct {
	Id            uint64    `json:"-" xorm:"bigserial pk autoincr"`
	CampaignId    uint64    `json:"campaign_id" xorm:"notnull index"`
	RecipienId    uint64    `json:"recipient_id" xorm:"notnull index"`
	PhoneNumber   string    `json:"phone_number" xorm:"varchar(24) notnull index"`
	RecipientName string    `json:"recipient_name" xorm:"varchar(64)"`
	ScheduleTime  time.Time `json:"schedule_time" xorm:"notnull"`
	DeliveredTime time.Time `json:"delivered_time"`
	Message       string    `json:"message_template" xorm:"varchar(640) notnull"`
	CreateDate    time.Time `json:"created_time" xorm:"created"`
	LastUpdate    time.Time `json:"updated_time" xorm:"updated"`
}
