package models

import (
	"time"
)

type Campaigns struct {
	Id              uint64    `json:"id" xorm:"bigserial pk autoincr"`
	MessageTemplate string    `json:"message_template" xorm:"varchar(512) notnull"`
	Enable          bool      `json:"enable" xorm:"bool default true comment('is this campaign available')"`
	ScheduleTime    time.Time `json:"schedule_time" xorm:"notnull"`
	ScheduleState   uint8     `json:"scheddule_state" xorm:"smallint default 0 comment('type of the schedule state, 0 - waiting, 1 - processed')"`
	CreateDate      time.Time `json:"created_time" xorm:"created"`
	LastUpdate      time.Time `json:"updated_time" xorm:"updated"`
}

type RecipientCSVPath struct {
	Id           uint64 `json:"-" xorm:"bigserial pk autoincr"`
	CampaignId   uint64 `json:"campaign_id" xorm:"notnull index"`
	CSVPath      string `json:"csv_path" xorm:"varchar(255)"`
	Url          string `json:"url" xorm:"varchar(255)"`
	UploadMethod uint8  `json:"upload_method" xorm:"smallint notnull comment(0 - upload by csv, 1 - upload by url)"`
}

func (t *RecipientCSVPath) TableName() string {
	return "recipient_csv_path"
}

type Recipients struct {
	Id            uint64    `json:"-" xorm:"bigserial pk autoincr"`
	CampaignId    uint64    `json:"campaign_id" xorm:"notnull unique(cp)"`
	PhoneNumber   string    `json:"phone_number" xorm:"varchar(24) notnull unique(cp)"`
	RecipientName string    `json:"recipient_name" xorm:"varchar(64)"`
	ScheduleTime  time.Time `json:"schedule_time" xorm:"notnull"` // redundancy, denormalization
	CreateDate    time.Time `json:"created_time" xorm:"created"`
	LastUpdate    time.Time `json:"updated_time" xorm:"updated"`
}

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
