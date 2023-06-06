package course_service

import (
	"encoding/json"
	"errors"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/dto/payload"
	"projects-subscribeme-backend/dto/response"
	"projects-subscribeme-backend/helper"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/repositories/course_repository"
	"projects-subscribeme-backend/repositories/user_repository"
	"sort"
	"time"
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

	userCourses, err := s.repository.GetUserCourseByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	return response.NewCourseMoodleResponses(*courseMoodle, userCourses), nil

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

	

	if len(assignment.Courses[0].Assignments) != 0 {
		for _, v := range assignment.Courses[0].Assignments {
			event := models.ClassEvent{
				CourseSceleID:  course.ID.String(),
				Type:           constant.AssignmentType,
				Date:           time.Unix(v.DueDate, 0),
				EventName:      v.Name,
				CourseModuleID: v.ID,
			}

			event, err := s.repository.FirstOrCreateEvent(event)
			if err != nil {
				log.Println(string("\033[31m"), err.Error())
				return nil, err
			}
		}
	}

	//Scele Quiz to Event type Assignment
	quiz, err := helper.GetMoodleData[models.ListQuizzez](constant.GetQuizFromCourseID, data)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}


	if len(quiz.CourseQuizzez) != 0 {
		for _, v := range quiz.CourseQuizzez {
			event := models.ClassEvent{
				CourseSceleID: course.ID.String(),
				Type:          constant.QuizType,
				Date:          time.Unix(v.TimeOpen, 0),
				EventName:     v.Name,
				CourseModuleID: v.ID,
			}

			event, err := s.repository.FirstOrCreateEvent(event)
			if err != nil {
				log.Println("ERROR ", err)
				log.Println(string("\033[31m"), err.Error())
			}
		}
	}

	//Mapping User Event
	events, err := s.repository.GetEventByCourseId(course.ID.String())
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	for _, v := range events {
		userEvent := models.UserEvent{
			UserID:   user.ID.String(),
			EventID:  v.ID.String(),
			CourseID: v.CourseSceleID,
			IsDone:   false,
		}

		userEvent, err := s.repository.CreateUserEvent(userEvent)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
		}
	}

	return response.NewCourseSceleResponse(course), nil

}

func (s *courseService) UnsibscribeCourse(claims *helper.JWTClaim, payload payload.ChooseCourse) (*response.CourseSceleResponse, error) {
	course, err := s.repository.GetCourseByCourseSceleId(payload.Id)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	//Delete User Event
	err = s.repository.DeleteUserEventByUserIdAndCourseId(user.ID.String(), course.ID.String())
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	//Delete User Course

	err = s.repository.DeletUserCourseByUserAndCourse(user, course)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewCourseSceleResponse(course), nil

}

func (s *courseService) GetUserCourseByUsername(claims *helper.JWTClaim) (*[]response.CourseSceleResponse, error) {
	userCourse, err := s.repository.GetUserCourseByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	return response.NewCourseSceleResponses(userCourse), nil

}

func (s *courseService) GetUserEventByCourseId(claims *helper.JWTClaim, courseId string) (*[]response.UserEventResponse, error) {
	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	userEvent, err := s.repository.GetUserEventByCourseId(courseId, user.ID.String())
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	sort.Slice(userEvent, func(i, j int) bool {
		return userEvent[i].ClassEvent.Date.After(userEvent[j].ClassEvent.Date)
	})

	return response.NewUserEventResponses(userEvent), nil
}

func (s *courseService) GetDeadlineTodayByUserId(claims *helper.JWTClaim) (*[]response.UserEventResponse, error) {
	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	userEvent, err := s.repository.GetDeadlineTodayByUserId(user.ID.String())
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	return response.NewUserEventResponses(userEvent), nil
}

func (s *courseService) GetDeadline7DaysAheadByUserId(claims *helper.JWTClaim) (*[]response.UserEventResponse, error) {
	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	userEvent, err := s.repository.GetDeadline7DaysAheadByUserId(user.ID.String())
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
	}

	return response.NewUserEventResponses(userEvent), nil
}

func (s *courseService) SetDeadlineReminder(claims *helper.JWTClaim, payload payload.ReminderPayload) (bool, error) {
	if payload.SetTime.Before(time.Now()) {
		return false, errors.New("400")
	}

	event, err := s.repository.GetEventByEventId(payload.EventID)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	if event.Date.Before(payload.SetTime) {
		return false, errors.New("409")
	}

	user, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return false, err
	}

	helper.SchedulerEvent.Schedule("ReminderEventSetDeadline", string(jsonBytes), payload.SetTime, user.ID.String(), event.ID.String())

	return true, nil

}
