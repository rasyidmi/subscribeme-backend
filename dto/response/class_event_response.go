package response

import (
	"projects-subscribeme-backend/constant"
	"time"
)

type ClassEventResponse struct {
	ID          string              `json:"id"`
	Type        constant.EventEnum  `json:"type"`
	Date        time.Time           `json:"date"`
	EventName   string              `json:"event_name"`
	CourseScele CourseSceleResponse `json:"course_scele"`
}
