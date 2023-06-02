package absensi_repository

import (
	"projects-subscribeme-backend/models"

	"gorm.io/gorm"
)

type absensiRepository struct {
	db *gorm.DB
}

func NewAbsenceRepository(db *gorm.DB) AbsensiRepository {
	return &absensiRepository{db: db}
}

func (r *absensiRepository) CreateAbsenceSession(absenceSession models.ClassAbsenceSession) (models.ClassAbsenceSession, error) {
	err := r.db.Create(&absenceSession).Error

	return absenceSession, err
}

func (r *absensiRepository) CreateAbsence(absence []models.Absence) ([]models.Absence, error) {
	err := r.db.Create(&absence).Error

	return absence, err

}

func (r *absensiRepository) UpdateAbsence(absence models.Absence, npm string, id string) (models.Absence, error) {
	err := r.db.Where("student_npm = ? AND class_absence_session_id = ?", npm, id).Updates(absence).Error

	if err != nil {
		return models.Absence{}, err
	}

	err = r.db.First(&absence, "student_npm = ? AND class_absence_session_id = ?", npm, id).Error
	if err != nil {
		return models.Absence{}, err
	}

	return absence, err

}

func (r *absensiRepository) GetIsOpenAbsenceSessionByClassCodeAndEndTime(classCode string) (models.ClassAbsenceSession, error) {
	var classAbsence models.ClassAbsenceSession
	err := r.db.Raw("select * from class_absence_sessions cas where cas.end_time >= now() and cas.class_code = ?", classCode).Scan(&classAbsence).Error
	return classAbsence, err
}

func (r *absensiRepository) GetAbsenceSessionById(id string) (models.ClassAbsenceSession, error) {
	var absenceSession models.ClassAbsenceSession

	err := r.db.First(&absenceSession, "id = ?", id).Error

	return absenceSession, err
}

func (r *absensiRepository) GetAbsenceSessionByClassCode(classCode string) ([]models.ClassAbsenceSession, error) {
	var absenceSession []models.ClassAbsenceSession

	err := r.db.Find(&absenceSession, "class_code = ?", classCode).Error

	return absenceSession, err

}

func (r *absensiRepository) GetAbsenceByClassCodeAndNpm(classCode string, npm string) ([]models.Absence, error) {
	var absences []models.Absence

	err := r.db.Raw("select a.* from absences a JOIN class_absence_sessions cas ON a.class_absence_session_id = cas.id AND cas.class_code = ? AND a.student_npm = ?", classCode, npm).Order("class_date DESC").Scan(&absences).Error

	return absences, err
}

func (r *absensiRepository) GetAbsenceSessionByAbsenceSessionId(id string) ([]models.ClassAbsenceSession, error) {
	var absenceSessions []models.ClassAbsenceSession

	err := r.db.Preload("Absence").First(&absenceSessions, "id = ?", id).Error

	return absenceSessions, err
}
