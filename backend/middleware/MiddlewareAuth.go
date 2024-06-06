package middleware

import (
	"backend/domain/users"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Definimos una clave secreta que se usará para firmar y verificar los tokens JWT.
var jwtKey = []byte("my_secret_key")

// AuthMiddleware devuelve una función middleware que se usa para proteger las rutas.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtiene el encabezado de autorización de la solicitud HTTP.
		authHeader := c.GetHeader("Authorization")

		// Si el encabezado de autorización está vacío, responde con un error 401.
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort() // Abortamos el manejo de la solicitud.
			return
		}

		// Elimina el prefijo "Bearer " del encabezado de autorización para obtener el token.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Creamos una estructura Claims para almacenar los claims del token.
		claims := &users.Claims{}

		// Analiza el token usando la clave secreta y los claims definidos.
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // Devuelve la clave secreta para verificar el token.
		})

		// Si hay un error en la validación o el token no es válido, responde con un error 401.
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // Abortamos el manejo de la solicitud.
			return
		}

		// Si el token es válido, almacenamos los valores de los claims en el contexto.
		c.Set("username", claims.Username)
		c.Set("user_type", claims.UserType)

		// Llamamos a la siguiente función en la cadena del middleware.
		c.Next()
	}
}

// AdminMiddleware devuelve una función middleware que se usa para proteger rutas que requieren privilegios de administrador.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("user_type")

		if !exists || userType != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
			c.Abort() // Abortamos el manejo de la solicitud.
			return
		}

		c.Next()
	}
}
