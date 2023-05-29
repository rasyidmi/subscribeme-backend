package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects-subscribeme-backend/config"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"
)

func GetSiakngData[T models.ClassSchedule | []models.ClassSchedule | []models.ClassDetail](api constant.SiakNGEnum, data map[string]interface{}) (*T, error) {
	siakng := config.LoadSiakNGConfig()

	client := &http.Client{}

	var req *http.Request

	if api == constant.GetClassScheduleByNpmMahasiswa {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetClassScheduleByNpmMahasiswa.String(), data["npm"], data["tahun"], data["term"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetClassScheduleByYearAndTerm {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetClassScheduleByYearAndTerm.String(), data["year"], data["term"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetClassScheduleByScheduleId {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetClassScheduleByScheduleId.String(), data["schedule_id"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else if api == constant.GetClassParticipantByClassCode {
		reqs, err := http.NewRequest("GET", fmt.Sprintf(constant.GetClassParticipantByClassCode.String(), data["kd_kls"]), nil)
		if err != nil {
			return nil, err
		}
		req = reqs
	} else {
		return nil, errors.New("ERROR : API NOT FOUND")
	}

	req.SetBasicAuth(siakng.Username, siakng.Password)

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
		return nil, err
	}

	return &responseObject, nil

}
