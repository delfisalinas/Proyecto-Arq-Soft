package router

import (
	"backend/controllers/users"
	//"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	// Ruta para inicio de sesión
	engine.POST("/users/login", users.Login)

	// Ruta para registro de nuevos usuarios
	engine.POST("/users/register", users.Register)

	/*
		// Rutas protegidas por middleware de autenticación
		auth := engine.Group("/auth")
		auth.Use(middleware.AuthMiddleware()) // Aplica el middleware de autenticación
	*/
}
