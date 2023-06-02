package course_service

import (
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
)

type cuorseService struct {
}

func NewCourseService() CourseService {
	return &cuorseService{}
}

func (s *cuorseService) GetCoursesByUsername(claims *helper.JWTClaim) (*[]response.CourseMoodleResponse, error) {
	var data = map[string]interface{}{}

	data["username"] = claims.Username

	userDetail, err := helper.GetMoodleData[models.Moodle](constant.GetUserDetailByUsername, data)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	indexUserDetail := *userDetail

	id := indexUserDetail.User[0].ID

	data["user_id"] = id

	courseMoodle, err := helper.GetMoodleData[[]models.CourseMoodle](constant.GetCourseByUserid, data)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewCourseMoodleResponse(*courseMoodle), nil

}
