package payload

type FcmPayload struct {
	FcmToken string `json:"fcm_token" binding:"required"`
}
