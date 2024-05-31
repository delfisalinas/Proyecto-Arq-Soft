package dtos

// LoginRequestDTO representa la solicitud de login.
type LoginRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponseDTO representa la respuesta de login.
type LoginResponseDTO struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

// RegisterRequestDTO representa la solicitud de registro.
type RegisterRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type"` // "alumno" o "administrador"
}

// RegisterResponseDTO representa la respuesta de registro.
type RegisterResponseDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Error    string `json:"error,omitempty"`
}
