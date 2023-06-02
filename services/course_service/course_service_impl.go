package course_service

import (
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/repositories/course_repository"
	"projects-subscribeme-backend/repositories/user_repository"
)

type courseService struct {
	repository     course_repository.CourseRepository
	userRepository user_repository.UserRepository
}

func NewCourseService(repository course_repository.CourseRepository, userRepository user_repository.UserRepository) CourseService {
	return &courseService{repository: repository, userRepository: userRepository}
}

func (s *courseService) GetCoursesByUsername(claims *helper.JWTClaim) (*[]response.CourseMoodleResponse, error) {
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

	return response.NewCourseMoodleResponses(*courseMoodle), nil

}

func (s *courseService) SubscribeCourse(claims *helper.JWTClaim, payload payload.ChooseCourse) (*response.CourseSceleResponse, error) {
	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	courseScele := models.CourseScele{
		CourseSceleID:   payload.Id,
		CourseSceleName: payload.Name,
	}

	course, err := s.repository.CreateCourse(courseScele, user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	//Scele Assignment to Event type Assignment
	var data = map[string]interface{}{}

	data["course_id"] = courseScele.CourseSceleID

	assignment, err := helper.GetMoodleData[models.ListCourses](constant.GetAssignmentFromCourseID, data)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}
	log.Println("MASUK")
	log.Println(assignment)

	return response.NewCourseSceleResponse(course), nil

}
