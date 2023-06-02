package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"
)

func GetMoodleData[T []models.CourseMoodle | models.ListCourses | models.ListQuizzez | models.Moodle](api constant.MoodleEnum, data map[string]interface{}) (*T, error) {
	client := &http.Client{}

	var req *http.Request

	token := "090a3215b2b7c80afc5cb77ff3b86009"

	if api == constant.GetCourseByUserid {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetCourseByUserid.String(), token, data["user_id"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetUserDetailByUsername {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetUserDetailByUsername.String(), token, data["username"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetAssignmentFromCourseID {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetAssignmentFromCourseID.String(), token, data["course_id"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetQuizFromCourseID {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetQuizFromCourseID.String(), token, data["course_id"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else {
		return nil, errors.New("ERROR : API NOT FOUND")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &responseObject, nil
}
