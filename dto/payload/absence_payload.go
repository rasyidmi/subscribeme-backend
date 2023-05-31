package payload

import "time"

type ClassAbsenceSessionPayload struct {
	StartTime  time.Time `json:"start_time"`
	Duration   float64   `json:"duration" binding:"required"`
	ClassCode  string    `json:"class_code" binding:"required"`
	IsGeofence bool      `json:"is_geofence"`
	GeoRadius  float64   `json:"geo_radius"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}

type AbsencePayload struct {
	ClassAbsenceSessionId string  `json:"class_session_id" binding:"required"`
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	DeviceCode            string  `json:"device_code" binding:"required"`
}
