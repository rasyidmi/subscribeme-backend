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
	ClassSchedule []ClassScheduleResponse `json:"class_schedule,omitempty"`
	Lecturers     []LecturersResponse     `json:"lectures"`
	ListStudent   []ListStudentResponse   `json:"list_student,omitempty"`
}

type CourseResponse struct {
	CourseCode     string `json:"course_code,omitempty"`
	CourseName     string `json:"course_name,omitempty"`
	CourseSKS      int    `json:"total_sks,omitempty"`
	CurriculumCode string `json:"curriculum_code,omitempty"`
}

type CourseResponseDosen struct {
	CourseCode          string                `json:"course_code,omitempty"`
	CourseName          string                `json:"course_name,omitempty"`
	CourseSKS           int                   `json:"total_sks,omitempty"`
	CurriculumCode      string                `json:"curriculum_code,omitempty"`
	ClassDetailResponse []ClassDetailResponse `json:"class_detail,omitempty"`
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
	iterator := 0
	for _, v := range models {

		var response ClassDetailResponse

		val, ok := classMap[v.ClassDetail.ClassCode]
		if ok {
			classSchedule := NewClassScheduleResponse(v)
			responses[val].ClassSchedule = append(responses[val].ClassSchedule, *classSchedule)
		} else {
			classMap[v.ClassDetail.ClassCode] = iterator
			classSchedule := NewClassScheduleResponse(v)
			response = *NewClassDetailResponse(v.ClassDetail)
			response.ClassSchedule = append(response.ClassSchedule, *classSchedule)
			responses = append(responses, response)
			iterator++
		}
	}

	return &responses
}

func NewCourceResponseDosen(models []models.ClassSchedule) *[]CourseResponseDosen {
	var responses []CourseResponseDosen

	courseMap := make(map[string]int)
	iterator := 0
	iterator2 := 0

	for _, v := range models {
		var response CourseResponseDosen

		val, ok := courseMap[v.ClassDetail.Course.CourseCode]

		if ok {
			classDetail := ClassDetailResponse{
				ClassName: v.ClassDetail.ClassName,
				ClassCode: v.ClassDetail.ClassCode,
			}
			copier.Copy(&classDetail.Lecturers, v.ClassDetail.Lecturers)
			log.Println("Masuk ", responses[val].ClassDetailResponse[iterator2].ClassCode)
			if responses[val].ClassDetailResponse[iterator2].ClassCode != v.ClassDetail.ClassCode {
				responses[val].ClassDetailResponse = append(responses[val].ClassDetailResponse, classDetail)
				iterator2++
			}

		} else {
			courseMap[v.ClassDetail.Course.CourseCode] = iterator
			response.CourseCode = v.ClassDetail.Course.CourseCode
			response.CourseName = v.ClassDetail.Course.CourseName
			response.CourseSKS = v.ClassDetail.Course.CourseSKS
			response.CurriculumCode = v.ClassDetail.Course.CurriculumCode

			classDetail := ClassDetailResponse{
				ClassName: v.ClassDetail.ClassName,
				ClassCode: v.ClassDetail.ClassCode,
			}
			copier.Copy(&classDetail.Lecturers, v.ClassDetail.Lecturers)
			
			response.ClassDetailResponse = append(response.ClassDetailResponse, classDetail)
			iterator++
			iterator2 = 0
			responses = append(responses, response)

		}
		

	}
	return &responses
}

func NewClassParticipantResponses(models []models.ClassDetail) *[]ListStudentResponse {
	var response []ClassDetailResponse

	copier.Copy(&response, models)

	return &response[0].ListStudent
}
