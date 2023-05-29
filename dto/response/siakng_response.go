package response

import (
	"projects-subscribeme-backend/models"
	"strings"

	"github.com/jinzhu/copier"
)

type ClassScheduleResponse struct {
	ScheduleId  string              `json:"schedule_id"`
	Day         string              `json:"day"`
	StartTime   string              `json:"start_time"`
	EndTime     string              `json:"end_time"`
	ClassDetail ClassDetailResponse `json:"class_detail"`
	Room        RoomResponse        `json:"room"`
}

type ClassDetailResponse struct {
	ClassName   string                `json:"class_name"`
	ClassCode   string                `json:"class_code"`
	Course      CourseResponse        `json:"course,omitempty"`
	Lecturers   []LecturersResponse   `json:"lectures"`
	ListStudent []ListStudentResponse `json:"list_student,omitempty"`
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

func NewClassDetailResponses(models []models.ClassDetail) *[]ClassDetailResponse {
	var response []ClassDetailResponse

	copier.Copy(&response, models)

	return &response
}
