package payload

type ClassCodePayload struct {
	ClassCode string `json:"class_code" binding:"required"`
}
