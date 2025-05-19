package schemas

type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
}

type SignupRequest struct {
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
