package router

import (
	filesController "backend/controllers/files"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// MapFileUrls mapea las rutas de archivos a sus respectivos controladores
func MapFileUrls(engine *gin.Engine, ctrl *filesController.Controller) {
	// Agrupar las rutas protegidas por el middleware
	auth := engine.Group("/files")
	auth.Use(middleware.AuthMiddleware()) // Aplica el middleware de autenticación

	// Ruta para cargar un nuevo archivo
	engine.POST("/files", ctrl.CreateFile)

	// Ruta para obtener archivos por ID del curso
	auth.GET("/course/:courseID", ctrl.GetFilesByCourseID)
}
