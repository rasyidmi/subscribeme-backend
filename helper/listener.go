package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"
	"time"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

func ReminderClassWillStarted(data map[string]interface{}) error {
	// Convert the map to JSON
	var classSchedule models.ClassSchedule

	mapstructure.Decode(data, &classSchedule)

	sendData := make(map[string]interface{})

	sendData["title"] = "Dalam 15 Menit Kelas Akan Dimulai"
	sendData["body"] = fmt.Sprintf("Kelas %s akan dimulai pada jam %s", classSchedule.ClassDetail.ClassName, classSchedule.StartTime)

	return nil

}

func ReminderAbsenceWillOver(data string, eventId uint, db *gorm.DB) error {
	var classAbsence models.ClassAbsenceSession

	err := json.Unmarshal([]byte(data), &classAbsence)

	if err != nil {
		fmt.Println(err)
	}

	//Get job
	var job models.Job
	err = db.Preload("User").First(&job, eventId).Error
	if err != nil {
		log.Println(err)
		return err
	}

	var dataMap = map[string]interface{}{}

	dataMap["code"] = classAbsence.ClassCode

	classDetail, err := GetSiakngData[models.ClassDetail](constant.GetClassByCode, dataMap)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return err
	}

	sendData := make(map[string]string)

	sendData["title"] = fmt.Sprintf("Cepat lakukan absensu pada kelas %s", classDetail.ClassName)
	sendData["body"] = fmt.Sprintf("Absensi kelas %s akan ditutup pada %s", classDetail.ClassName, classAbsence.EndTime.Format("2 January 2006 - 15:04"))

	err = SendPushNotification(sendData)
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Delete(&models.Job{}, eventId).Error
	if err != nil {
		log.Print("💀 error: ", err)
	}

	return nil
}

func ReminderAbsenceCanBeDone(data string, eventId uint, db *gorm.DB) error {
	var classAbsence models.ClassAbsenceSession

	err := json.Unmarshal([]byte(data), &classAbsence)

	if err != nil {
		fmt.Println(err)
	}

	//Get job
	var job models.Job
	err = db.Preload("User").First(&job, eventId).Error
	if err != nil {
		log.Println(err)
		return err
	}

	var dataMap = map[string]interface{}{}

	dataMap["code"] = classAbsence.ClassCode

	classDetail, err := GetSiakngData[models.ClassDetail](constant.GetClassByCode, dataMap)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return err
	}

	sendData := make(map[string]string)

	sendData["title"] = fmt.Sprintf("Absensi kelas %s sudah dibuka", classDetail.ClassName)
	sendData["body"] = fmt.Sprintf("Absensi kelas %s akan ditutup pada %s", classDetail.ClassName, classAbsence.EndTime.Format("2 January 2006 - 15:04"))

	err = SendPushNotification(sendData)
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Delete(&models.Job{}, eventId).Error
	if err != nil {
		log.Print("💀 error: ", err)
	}

	return nil

}

func ReminderAssignmentSetDeadline(data string, eventId uint, db *gorm.DB) error {
	var classEvent models.ClassEvent

	err := json.Unmarshal([]byte(data), &classEvent)

	if err != nil {
		fmt.Println(err)
	}

	//Get job
	var job models.Job
	err = db.Preload("User").First(&job, eventId).Error
	if err != nil {
		log.Println(err)
		return err
	}

	diff := classEvent.Date.Sub(job.RunAt)

	sendData := make(map[string]string)

	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60

	message := "Tugas Anda tersisa"
	if days > 0 {
		message += fmt.Sprintf(" %d hari", days)
	}
	if hours > 0 {
		message += fmt.Sprintf(" %d jam", hours)
	}
	if minutes > 0 {
		message += fmt.Sprintf(" %d menit", minutes)
	}
	message += " lagi"

	sendData["title"] = message
	sendData["body"] = fmt.Sprintf("Tugas %s akan berakhir pada %s", classEvent.EventName, classEvent.Date.Format("2 January 2006 - 15:04"))
	sendData["token"] = job.User.FcmToken

	err = SendPushNotification(sendData)
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Delete(&models.Job{}, eventId).Error
	if err != nil {
		log.Print("💀 error: ", err)
	}

	return nil
}

func ReminderQuizSetDeadline(data string, eventId uint, db *gorm.DB) error {
	var classEvent models.ClassEvent

	err := json.Unmarshal([]byte(data), &classEvent)

	if err != nil {
		fmt.Println(err)
	}

	//Get job
	var job models.Job
	err = db.Preload("User").First(&job, eventId).Error
	if err != nil {
		log.Println(err)
		return err
	}

	diff := classEvent.Date.Sub(job.RunAt)

	sendData := make(map[string]string)

	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60

	message := "Kuis akan dimulai "
	if days > 0 {
		message += fmt.Sprintf(" %d hari", days)
	}
	if hours > 0 {
		message += fmt.Sprintf(" %d jam", hours)
	}
	if minutes > 0 {
		message += fmt.Sprintf(" %d menit", minutes)
	}
	message += " lagi"

	sendData["title"] = message
	sendData["body"] = fmt.Sprintf("Kuis %s akan dimulai pada %s", classEvent.EventName, classEvent.Date.Format("2 January 2006 - 15:04"))
	sendData["token"] = job.User.FcmToken

	err = SendPushNotification(sendData)
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Delete(&models.Job{}, eventId).Error
	if err != nil {
		log.Print("💀 error: ", err)
	}

	return nil
}

func UpdateAllAssignmentAndQuizData(data string, eventId uint, db *gorm.DB) error {
	//Find Course
	//time.Unix(v.DueDate, 0)
	var courses []models.CourseScele

	err := db.Preload("User").Find(&courses).Error
	if err != nil {
		log.Print("💀 error: ", err)
	}

	for _, course := range courses {
		// Update Assignments

		// Get Class Event By Course
		var classEvent []models.ClassEvent

		err = db.Find(&classEvent, "course_scele_id = ?", course.ID).Error
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
			return err
		}

		//Update Assignments
		var data = map[string]interface{}{}

		data["course_id"] = course.CourseSceleID

		courseFromScele, err := GetMoodleData[models.ListCourses](constant.GetAssignmentFromCourseID, data)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
			return err
		}

		numbers := make([]int64, 0)

		for _, ce := range classEvent {
			numbers = append(numbers, ce.CourseModuleID)
		}

		assignment := courseFromScele.Courses[0].Assignments

		for _, am := range assignment {
			timeModified := time.Unix(am.TimeModified, 0)

			oneDayAgo := time.Now().AddDate(0, 0, -1)
			if timeModified.After(oneDayAgo) && timeModified.Before(time.Now()) {
				if slices.Contains(numbers, am.ID) {
					// Update course
					newClassEvent := models.ClassEvent{
						Date:      time.Unix(am.DueDate, 0),
						EventName: am.Name,
					}
					err := db.Model(&models.ClassEvent{}).Omit("CourseScele").Where("course_module_id = ?", am.ID).Updates(newClassEvent).Error
					if err != nil {
						log.Println(string("\033[31m"), err.Error())
						return err
					}
					break
				} else {
					newClassEvent := models.ClassEvent{
						Date:           time.Unix(am.DueDate, 0),
						EventName:      am.Name,
						CourseModuleID: am.ID,
						Type:           constant.AssignmentType,
						CourseSceleID:  course.ID.String(),
					}
					err := db.Create(&newClassEvent).Error
					if err != nil {
						log.Println(string("\033[31m"), err.Error())
						return err
					}

					//Input Ke User Event
					for _, user := range course.User {
						userEvent := models.UserEvent{
							UserID:   user.ID.String(),
							EventID:  newClassEvent.ID.String(),
							CourseID: course.ID.String(),
							IsDone:   false,
						}

						err := db.Create(&userEvent).Error
						if err != nil {
							log.Println(string("\033[31m"), err.Error())
							return err
						}
					}
				}
			}

		}

		//Update Quiz

		listQuizzez, err := GetMoodleData[models.ListQuizzez](constant.GetQuizFromCourseID, data)
		if err != nil {
			log.Println(string("\033[31m"), err.Error())
			return err
		}

		quizzez := listQuizzez.CourseQuizzez

		for _, quiz := range quizzez {
			timeModified := time.Unix(quiz.TimeModified, 0)
			oneDayAgo := time.Now().AddDate(0, 0, -1)
			if timeModified.After(oneDayAgo) && timeModified.Before(time.Now()) {
				if slices.Contains(numbers, quiz.ID) {
					newClassEvent := models.ClassEvent{
						Date:      time.Unix(quiz.TimeOpen, 0),
						EventName: quiz.Name,
					}
					err := db.Model(&models.ClassEvent{}).Omit("CourseScele").Where("course_module_id = ?", quiz.ID).Updates(newClassEvent).Error
					if err != nil {
						log.Println(string("\033[31m"), err.Error())
						return err
					}
				} else {
					newClassEvent := models.ClassEvent{
						Date:           time.Unix(quiz.TimeOpen, 0),
						EventName:      quiz.Name,
						CourseModuleID: quiz.ID,
						Type:           constant.AssignmentType,
						CourseSceleID:  course.ID.String(),
					}
					err := db.Create(&newClassEvent).Error
					if err != nil {
						log.Println(string("\033[31m"), err.Error())
						return err
					}

					//Input Ke User Event
					for _, user := range course.User {
						userEvent := models.UserEvent{
							UserID:   user.ID.String(),
							EventID:  newClassEvent.ID.String(),
							CourseID: course.ID.String(),
							IsDone:   false,
						}

						err := db.Create(&userEvent).Error
						if err != nil {
							log.Println(string("\033[31m"), err.Error())
							return err
						}
					}
				}
			}

		}

	}

	return nil
}

func TestCron(data string, eventId uint, db *gorm.DB) error {
	log.Println("Masuk Cron")
	return nil
}
