package dto

type Refresh struct {
	Token string `json"token" binding:"required"`
}
