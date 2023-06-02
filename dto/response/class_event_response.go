package response

import (
	"projects-subscribeme-backend/constant"
	"time"
)

type ClassEventResponse struct {
	ID            string             `json:"id"`
	CourseSceleID string             `json:"course_scele"`
	Type          constant.EventEnum `json:"type"`
	Date          time.Time          `json:"date"`
	EventName     string             `json:"event_name"`
}
