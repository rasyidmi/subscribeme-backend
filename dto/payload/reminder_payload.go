package payload

import (
	"projects-subscribeme-backend/constant"
	"time"
)

type ReminderPayload struct {
	SetTime time.Time          `json:"set_time"`
	EventID string             `json:"event_id"`
	Type    constant.EventEnum `json:"type"`
}
