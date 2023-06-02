package response

import (
	"projects-subscribeme-backend/constant"
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
)

type UserResponse struct {
	ID       string                `json:"id,omitempty"`
	Username string                `json:"username"`
	Role     constant.UserRoleEnum `json:"role"`
	FcmToken string                `json:"fcm_token"`
}

type UserEventResponse struct {
	ID         string             `json:"id"`
	UserID     string             `json:"user_id"`
	EventID    string             `json:"event_id"`
	CourseID   string             `json:"course_id"`
	IsDone     bool               `json:"is_done"`
	ClassEvent ClassEventResponse `json:"class_event"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	IsUserExists *bool  `json:"is_user_exists,omitempty"`
}

func NewUserResponse(user models.User) *UserResponse {
	response := &UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
		FcmToken: user.FcmToken,
	}

	return response
}

func NewLoginResponse(token string, isUserExists *bool) *LoginResponse {
	response := &LoginResponse{
		Token:        token,
		IsUserExists: isUserExists,
	}

	return response
}

func NewUserEventResponses(models []models.UserEvent) *[]UserEventResponse {
	var responses []UserEventResponse

	copier.Copy(&responses, models)

	return &responses

}
