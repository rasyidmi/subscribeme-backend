package helper

import (
	"context"
	"log"
	"projects-subscribeme-backend/models"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var SchedulerEvent Scheduler

// Scheduler data structure
type Scheduler struct {
	db          *gorm.DB
	listeners   Listeners
	cron        *cron.Cron
	cronEntries map[string]cron.EntryID
}

// Listeners has attached event listeners
type Listeners map[string]ListenFunc

// ListenFunc function that listens to events
type ListenFunc func(string, uint, *gorm.DB) error

// Event structure
type Event struct {
	ID      uint
	Name    string
	Payload string
	Cron    string
}

// NewScheduler creates a new scheduler
func NewScheduler(db *gorm.DB, listeners Listeners) Scheduler {

	return Scheduler{
		db:          db,
		listeners:   listeners,
		cron:        cron.New(),
		cronEntries: map[string]cron.EntryID{},
	}

}

// AddListener adds the listener function to Listeners
func (s Scheduler) AddListener(event string, listenFunc ListenFunc) {
	s.listeners[event] = listenFunc
}

// callListeners calls the event listener of provided event
func (s Scheduler) callListeners(event Event) {
	eventFn, ok := s.listeners[event.Name]
	if ok {
		go eventFn(event.Payload, event.ID, s.db)

	} else {
		log.Print("💀 error: couldn't find event listeners attached to ", event.Name)
	}

}

// CheckEventsInInterval checks the event in given interval
func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				events := s.checkDueEvents()
				for _, e := range events {
					s.callListeners(e)
				}
			}

		}
	}()
}

// checkDueEvents checks and returns due events
func (s Scheduler) checkDueEvents() []Event {
	events := []Event{}
	rows, err := s.db.Model(&models.Job{}).Where("run_at <= ? AND cron IS NULL", time.Now()).Select("id", "name", "payload").Rows()
	defer rows.Close()
	if err != nil {
		log.Print("💀 error: ", err)
		return events
	}

	for rows.Next() {
		evt := Event{}
		rows.Scan(&evt.ID, &evt.Name, &evt.Payload)
		events = append(events, evt)
	}
	return events
}

// Schedule sechedules the provided events
func (s Scheduler) Schedule(event string, payload string, runAt time.Time, userId string, eventId string) {
	log.Print("🚀 Scheduling event ", event, " to run at ", runAt)
	job := &models.Job{
		Name:    event,
		Payload: payload,
		RunAt:   runAt,
		UserID:  userId,
		Cron:    nil,
	}

	if eventId != "" {
		job.EventID = eventId
	}

	err := s.db.FirstOrCreate(&job, models.Job{UserID: userId, Name: event, Payload: payload, RunAt: runAt}).Error
	if err != nil {
		log.Print("schedule insert error: ", err)
	}
}

// ScheduleCron schedules a cron job
func (s Scheduler) ScheduleCron(event string, payload string, cron string, userId string, eventId string) {
	log.Print("🚀 Scheduling event ", event, " with cron string ", cron)
	entryID, ok := s.cronEntries[event]
	var jobs models.Job
	if ok {
		s.cron.Remove(entryID)
		err := s.db.Model(&jobs).Where("name = ? AND run_at >= ? AND cron IS NOT NULL", time.Now(), event).Updates(&models.Job{Cron: &cron, Payload: payload}).Error
		if err != nil {
			log.Print("schedule cron update error: ", err)
		}
	} else {
		job := &models.Job{
			Name:    event,
			Payload: payload,
			RunAt:   time.Now(),
			Cron:    &cron,
			UserID:  userId,
		}

		if eventId != "" {
			job.EventID = eventId
		}

		err := s.db.Create(&job).Error
		if err != nil {
			log.Print("schedule cron insert error: ", err)
		}

		jobs = *job
	}

	eventFn, ok := s.listeners[event]
	if ok {
		entryID, err := s.cron.AddFunc(cron, func() { eventFn(payload, jobs.ID, s.db) })
		s.cronEntries[event] = entryID
		if err != nil {
			log.Print("💀 error: ", err)
		}
	}
}

// attachCronJobs attaches cron jobs
func (s Scheduler) attachCronJobs() {
	log.Printf("Attaching cron jobs")
	log.Println(time.Now())
	rows, err := s.db.Model(&models.Job{}).Where("run_at <= ? AND cron IS NOT NULL", time.Now()).Select("id", "name", "payload").Rows()
	defer rows.Close()
	if err != nil {
		log.Print("💀 error: ", err)
	}
	for rows.Next() {
		evt := Event{}
		rows.Scan(&evt.ID, &evt.Name, &evt.Payload, &evt.Cron)
		eventFn, ok := s.listeners[evt.Name]
		if ok {
			entryID, err := s.cron.AddFunc(evt.Cron, func() { eventFn(evt.Payload, 1, nil) })
			s.cronEntries[evt.Name] = entryID

			if err != nil {
				log.Print("💀 error: ", err)
			}
		}
	}
}

// StartCron starts cron job
func (s Scheduler) StartCron() func() {
	s.attachCronJobs()
	s.cron.Start()

	SchedulerEvent = s

	return func() {
		s.cron.Stop()
	}
}
