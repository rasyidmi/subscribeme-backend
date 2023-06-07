package initializers

import (
	"context"
	"projects-subscribeme-backend/helper"
	"time"
)

var eventListeners = helper.Listeners{
	"ReminderEventSetDeadline":       helper.ReminderEventSetDeadline,
	"UpdateAllAssignmentAndQuizData": helper.UpdateAllAssignmentAndQuizData,
	"TestCron":                       helper.TestCron,
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