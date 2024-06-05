package router

import (
	"backend/controllers/users"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine, usersController *users.Controller) {
	// Rutas protegidas por middleware de autenticación
	auth := engine.Group("/auth")
	auth.Use(middleware.AuthMiddleware()) // Aplica el middleware de autenticación

	// Ruta para inicio de sesión
	engine.POST("/users/login", usersController.Login)

	// Ruta para registro de nuevos usuarios
	engine.POST("/users/register", usersController.Register)
}
