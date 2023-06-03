package helper

import (
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"
)

func ReminderClassWillStarted(data map[string]interface{}) error {
	_, err := GetSiakngData[[]models.ClassSchedule](constant.GetClassDetailByNpmMahasiswa, data)

	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return err
	}

	return nil

}
