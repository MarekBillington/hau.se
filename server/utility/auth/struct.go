package auth

type register struct {
	Email    string `json"email" binding:"required"`
	Password string `json"password" binding:"required"`
}

type refresh struct {
	Token string `json"token" binding:"required"`
}
