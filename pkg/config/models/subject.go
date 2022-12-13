package models

type Subject struct {
	ID      int     `gorm:"primaryKey"`
	Title   string  `gorm:"not null;unique"`
	Term    int     `gorm:"not null"`
	Major   string  `gorm:"not null"`
	Classes []Class `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
}
