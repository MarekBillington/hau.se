package auth

type login struct {
	Email    string `json"email" binding:"required"`
	Password string `json"password" binding:"required"`
}

type register struct {
	Email     string `json"email" binding:"required"`
	Password  string `json"password" binding:"required"`
	FirstName string `json"password" binding:"required"`
	LastName  string `json"password" binding:"required"`
}

type refresh struct {
	Token string `json"token" binding:"required"`
}
