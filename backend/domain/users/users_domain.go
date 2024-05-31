package users

import (
	"github.com/dgrijalva/jwt-go"
)

// Estructura para la solicitud de Login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Estructura para la respuesta de Login
type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

// Estructura para la solicitud de registro
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type" binding:"required"` // "alumno" o "administrador"
}

// Estructura para la respuesta de registro
type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Error    string `json:"error,omitempty"`
}

// Estructura de usuario, define los atributos esenciales de un usuario y cómo se deben almacenar en la base de datos.
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`   // Campo ID, clave primaria
	Username string `json:"username" gorm:"unique"` // Campo Username, único
	Email    string `json:"email" gorm:"unique"`    // Campo Email, único
	Password string `json:"password"`               // Campo Password (HASHEADA VERIFICAR)
	UserType string `json:"user_type"`              // Campo UserType
}

type Claims struct { //Define los datos que se incluirán en un token JWT(JasonWebToken) para verificar la identidad del usuario
	Username string `json:"username"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}
