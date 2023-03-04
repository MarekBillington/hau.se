package dto

type Register struct {
	Email     string `json"email" binding:"required"`
	Password  string `json"password" binding:"required"`
	FirstName string `json"password" binding:"required"`
	LastName  string `json"password" binding:"required"`
}
