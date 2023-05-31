package absensi_repository

import (
	"projects-subscribeme-backend/models"
)

type AbsensiRepository interface {
	CreateAbsenceSession(absenceSession models.ClassAbsenceSession) (models.ClassAbsenceSession, error)
	CreateAbsence(absence []models.Absence) ([]models.Absence, error)
	UpdateAbsence(absence models.Absence, npm string, id string) (models.Absence, error)

	GetIsOpenAbsenceSessionByClassCodeAndEndTime(classCode string) (models.ClassAbsenceSession, error)

	GetAbsenceSessionByClassCode(classCode string) ([]models.ClassAbsenceSession, error)
	GetAbsenceSessionById(id string) (models.ClassAbsenceSession, error)

	GetAbsenceByAbsenceSessionId(id string) ([]models.Absence, error)
	GetAbsenceByClassCodeAndNpm(classCode string, npm string) ([]models.Absence, error)
}
