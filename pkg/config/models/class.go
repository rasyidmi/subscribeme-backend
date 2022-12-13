package models

type Class struct {
	ID        int      `gorm:"primaryKey"`
	SubjectID int      `gorm:"not null"`
	Title     string   `gorm:"not null"`
	Events    []*Event `gorm:"many2many:class_events;constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Users     []User   `gorm:"many2many:class_students;constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
}

type ClassResponse struct {
	ID     int
	Title  string
	Events []*EventResponse
}
