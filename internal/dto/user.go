package dto

// CreateUserRequest defines the payload for creating a user
// @Description Payload for creating a new user
type CreateUserRequest struct {
	Login    string `json:"login" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=8"`
	RoleId   uint   `json:"roleId" binding:"required"`
}
