package initializers

import (
	"context"
	"projects-subscribeme-backend/helper"
	"time"
)

var eventListeners = helper.Listeners{
	"ReminderClassWillStarted":       helper.ReminderClassWillStarted,
	"ReminderAssignmentSetDeadline":  helper.ReminderAssignmentSetDeadline,
	"ReminderQuizSetDeadline":        helper.ReminderQuizSetDeadline,
	"UpdateAllAssignmentAndQuizData": helper.UpdateAllAssignmentAndQuizData,
	"ReminderAbsenceCanBeDone":       helper.ReminderAbsenceCanBeDone,
	"ReminderAbsenceWillOver":        helper.ReminderAbsenceWillOver,
	"ReminderQuizWillOver":           helper.ReminderQuizWillOver,
}

var SchedulerEvent helper.Scheduler

func SetupScheduler() {
	scheduler := helper.NewScheduler(DB, eventListeners)

	scheduler.StartCron()

	scheduler.CheckEventsInInterval(context.Background(), 5*time.Second)
	// scheduler.Schedule("UpdateAllAssignmentAndQuizData", "", time.Now().Add(5*time.Second), "", "")
	// scheduler.ScheduleCron("UpdateAllAssignmentAndQuizData", "", "* * * * *", "", "", time.Now().Add(30*time.Second))
	// scheduler.ScheduleCron("UpdateAllAssignmentAndQuizData", "", "45 1 * * *", "", "")

}
