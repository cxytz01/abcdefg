package core

import (
	"time"
)

type CampaignParam struct {
	MessageTemplate string    `form:"message_template" json:"message_template"`
	ScheduleTime    time.Time `form:"schedule_time" json:"schedule_time" format:"date-time" example:"2024-09-24T16:57:00+08:00 defined by RFC3339, section5.6"`
}

type CreateCampaignParam struct {
	CampaignParam `json:",inline"`
}

type UpdateCampaignParam struct {
	CampaignParam `json:",inline"`
	Enable        bool  `form:"enable" json:"enable"`
	ScheduleState uint8 `form:"schedule_state" json:"schedule_state"`
}
