package users

import (
	"github.com/dgrijalva/jwt-go"
)

// Estructura para la solicitud de Login
type LoginRequest struct {
	Username string `gorm:"type:varchar(191)"`
	Password string `gorm:"type:longtext"`
}

// Estructura para la solicitud de registro
type RegisterRequest struct {
	Username string `gorm:"type:varchar(191)" binding:"required"`
	Email    string `gorm:"type:varchar(191)" binding:"required,email"`
	Password string `gorm:"type:longtext" binding:"required"`
	UserType string `gorm:"default:'alumno'"` // "alumno" o "administrador"
}

// Estructura de usuario, define los atributos esenciales de un usuario y cómo se deben almacenar en la base de datos.
type User struct {
	ID       uint   `gorm:"primaryKey;AUTO_INCREMENT"` // Campo ID, clave primaria
	Username string `gorm:"type:varchar(191);unique"`  // Campo Username, único
	Email    string `gorm:"type:varchar(191);unique"`  // Campo Email, único
	Password string `gorm:"type:longtext"`             // Campo Password (HASHEADA VERIFICAR)
	UserType string `gorm:"type:longtext"`             // Campo UserType
}

type Claims struct { //Define los datos que se incluirán en un token JWT(JasonWebToken) para verificar la identidad del usuario
	Username string `gorm:"type:varchar(191)"`
	UserType string `gorm:"type:longtext"`
	jwt.StandardClaims
}
