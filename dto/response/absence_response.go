package response

import (
	"projects-subscribeme-backend/models"
	"time"

	"github.com/jinzhu/copier"
)

type ClassAbsenceSessionResponse struct {
	ID          string    `json:"id"`
	TeacherName string    `json:"teacher_name"`
	ClassCode   string    `json:"class_code"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	IsGeofence  bool      `json:"is_geofence"`
	GeoRadius   float64   `json:"geo_radius"`
	Latitude    float64   `json:"latitude,omitempty"`
	Longitude   float64   `json:"longitude,omitempty"`
}

type AbsenceResponse struct {
	ClassAbsenceSessionID string    `json:"class_absence_session_id"`
	StudentName           string    `json:"student_name"`
	StudentNpm            string    `json:"student_npm"`
	Latitude              float64   `json:"latitude"`
	Longitude             float64   `json:"longitude"`
	DeviceCode            string    `json:"device_code"`
	PresentTime           time.Time `json:"present_time"`
	Present               bool      `json:"present"`
}

func NewClassAbsenceSessionResponse(model models.ClassAbsenceSession) *ClassAbsenceSessionResponse {
	response := &ClassAbsenceSessionResponse{
		ID:          model.ID.String(),
		TeacherName: model.TeacherName,
		ClassCode:   model.ClassCode,
		StartTime:   model.StartTime,
		EndTime:     model.EndTime,
	}

	if model.IsGeofence {
		response.IsGeofence = true
		response.GeoRadius = model.GeoRadius
		response.Latitude = model.Latitude
		response.Longitude = model.Longitude
	}

	return response
}

func NewClassAbsenceSessionResponses(model []models.ClassAbsenceSession) *[]ClassAbsenceSessionResponse {
	var responses []ClassAbsenceSessionResponse

	for _, v := range model {
		response := NewClassAbsenceSessionResponse(v)
		responses = append(responses, *response)

	}

	return &responses
}

func NewAbsenceResponse(model models.Absence) *AbsenceResponse {
	var response AbsenceResponse
	copier.Copy(&response, model)

	return &response

}

func NewAbsenceResponses(model []models.Absence) *[]AbsenceResponse {
	var responses []AbsenceResponse
	for _, v := range model {
		responses = append(responses, *NewAbsenceResponse(v))
	}

	return &responses

}
