package payload

import "time"

type ReminderPayload struct {
	SetTime time.Time `json:"set_time"`
	EventID string    `json:"event_id"`
}
