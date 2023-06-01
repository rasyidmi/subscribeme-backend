package response

import (
	"log"
	"projects-subscribeme-backend/models"
	"strings"

	"github.com/jinzhu/copier"
)

type ClassScheduleResponse struct {
	ScheduleId string       `json:"schedule_id,omitempty"`
	Day        string       `json:"day,omitempty"`
	StartTime  string       `json:"start_time,omitempty"`
	EndTime    string       `json:"end_time,omitempty"`
	Room       RoomResponse `json:"room,omitempty"`
}

type ClassDetailResponse struct {
	ClassName     string                  `json:"class_name"`
	ClassCode     string                  `json:"class_code"`
	Course        CourseResponse          `json:"course,omitempty"`
	ClassSchedule []ClassScheduleResponse `json:"class_schedule_response,omitempty"`
	Lecturers     []LecturersResponse     `json:"lectures"`
	ListStudent   []ListStudentResponse   `json:"list_student,omitempty"`
}

type CourseResponse struct {
	CourseCode string `json:"course_code,omitempty"`
	CourseName string `json:"course_name,omitempty"`
	CourseSKS  int    `json:"total_sks,omitempty"`
}

type LecturersResponse struct {
	Name string `json:"name"`
}

type RoomResponse struct {
	RoomName string `json:"room_name"`
}

type ListStudentResponse struct {
	Student []StudentResponse `json:"students"`
}

type StudentResponse struct {
	Name string `json:"name"`
	Npm  string `json:"npm"`
}

func NewClassScheduleResponse(models models.ClassSchedule) *ClassScheduleResponse {
	var response ClassScheduleResponse

	copier.Copy(&response, models)

	replacer := strings.NewReplacer("http://api-kp.cs.ui.ac.id/siakngcs/jadwal/", "", "/", "")

	scheduleId := replacer.Replace(models.ScheduleUrl)

	response.ScheduleId = scheduleId

	return &response
}

func NewClassScheduleResponses(models []models.ClassSchedule) *[]ClassScheduleResponse {
	var response []ClassScheduleResponse

	copier.Copy(&response, models)

	for i := 0; i < len(response); i++ {
		replacer := strings.NewReplacer("http://api-kp.cs.ui.ac.id/siakngcs/jadwal/", "", "/", "")

		scheduleId := replacer.Replace(models[i].ScheduleUrl)

		response[i].ScheduleId = scheduleId
	}

	return &response
}

func NewClassDetailResponse(models models.ClassDetail) *ClassDetailResponse {
	var response ClassDetailResponse
	copier.Copy(&response, models)

	return &response
}

func NewClassDetailResponses(models []models.ClassSchedule) *[]ClassDetailResponse {
	var responses []ClassDetailResponse

	classMap := make(map[string]int)

	for _, v := range models {
		iterator := 0
		var response ClassDetailResponse
		log.Println(v.ClassDetail.ClassName)

		val, ok := classMap[v.ClassDetail.ClassCode]
		if ok {
			log.Println("MASUK OK")
			classSchedule := NewClassScheduleResponse(v)
			responses[val].ClassSchedule = append(responses[val].ClassSchedule, *classSchedule)
		} else {
			log.Println("MASUK ELSE")
			classMap[v.ClassDetail.ClassCode] = iterator
			classSchedule := NewClassScheduleResponse(v)
			response = *NewClassDetailResponse(v.ClassDetail)
			response.ClassSchedule = append(response.ClassSchedule, *classSchedule)
			responses = append(responses, response)
			iterator++
		}
		log.Println(responses)
	}

	return &responses
}

func NewClassParticipantResponses(models []models.ClassDetail) *[]ClassDetailResponse {
	var response []ClassDetailResponse

	copier.Copy(&response, models)

	return &response
}
