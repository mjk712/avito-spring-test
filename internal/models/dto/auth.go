package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // "employee" или "moderator"
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type DummyLoginRequest struct {
	Role string `json:"role"`
}

type DummyLoginResponse struct {
	Token string `json:"token"`
}
