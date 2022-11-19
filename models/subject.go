package models

type Subject struct {
	ID      int     `gorm:"primaryKey"`
	Title   string  `gorm:"not null;unique"`
	Term    int     `gorm:"not null"`
	Major   string  `gorm:"not null"`
	Classes []Class `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
}

type SubjectRequest struct {
	Title   string              `json:"title" binding:"required"`
	Term    int                 `json:"term" binding:"required"`
	Major   string              `json:"major" binding:"required"`
	Classes []map[string]string `json:"classes"`
}

type SubjectResponse struct {
	ID      int
	Title   string
	Term    int
	Major   string
	Classes []map[string]string
}
