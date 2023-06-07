package absensi_repository

import (
	"projects-subscribeme-backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *absensiRepository) UpdateAbsenceSession(absenceSession models.ClassAbsenceSession, id string) (models.ClassAbsenceSession, error) {
	err := r.db.Model(&models.ClassAbsenceSession{}).Where("id = ?", id).Updates(map[string]interface{}{"start_time": absenceSession.StartTime, "end_time": absenceSession.EndTime, "is_geofence": absenceSession.IsGeofence, "geo_radius": absenceSession.GeoRadius, "latitude": absenceSession.Latitude, "longitude": absenceSession.Longitude}).Error
	if err != nil {
		return models.ClassAbsenceSession{}, err
	}

	err = r.db.First(&absenceSession, "id = ?", id).Error

	return absenceSession, err
}

func (r *absensiRepository) CreateAbsence(absence []models.Absence) ([]models.Absence, error) {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "class_absence_session_id"}, {Name: "student_npm"}},
		DoUpdates: clause.AssignmentColumns([]string{"latitude", "longitude", "class_date"}),
	}).Create(&absence).Error

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
	err := r.db.Raw("select * from class_absence_sessions cas where cas.end_time >= now() and cas.start_time <= now() and cas.class_code = ?", classCode).Scan(&classAbsence).Error
	return classAbsence, err
}

func (r *absensiRepository) GetAbsenceSessionById(id string) (models.ClassAbsenceSession, error) {
	var absenceSession models.ClassAbsenceSession

	err := r.db.First(&absenceSession, "id = ?", id).Error

	return absenceSession, err
}

func (r *absensiRepository) GetAbsenceSessionByClassCode(classCode string) ([]models.ClassAbsenceSession, error) {
	var absenceSession []models.ClassAbsenceSession

	err := r.db.Preload("Absence").Find(&absenceSession, "class_code = ?", classCode).Error

	return absenceSession, err

}

func (r *absensiRepository) GetAbsenceByClassCodeAndNpm(classCode string, npm string) ([]models.Absence, error) {
	var absences []models.Absence

	err := r.db.Raw("select a.* from absences a JOIN class_absence_sessions cas ON a.class_absence_session_id = cas.id AND cas.class_code = ? AND a.student_npm = ? ORDER By class_date desc", classCode, npm).Scan(&absences).Error

	return absences, err
}

func (r *absensiRepository) GetAbsenceSessionByAbsenceSessionId(id string) ([]models.ClassAbsenceSession, error) {
	var absenceSessions []models.ClassAbsenceSession

	err := r.db.Preload("Absence").First(&absenceSessions, "id = ?", id).Error

	return absenceSessions, err
}
