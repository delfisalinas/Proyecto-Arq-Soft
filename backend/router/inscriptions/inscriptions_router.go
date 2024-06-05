package inscriptions

import (
	inscriptionsController "backend/controllers/inscriptions"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// MapInscriptionUrls mapea las rutas de inscripciones a sus respectivos controladores
func MapInscriptionUrls(engine *gin.Engine, ctrl *inscriptionsController.Controller) {
	// Agrupar las rutas protegidas por el middleware
	auth := engine.Group("/inscriptions")
	auth.Use(middleware.AuthMiddleware()) // Aplica el middleware de autenticación

	// Ruta para crear una nueva inscripción
	engine.POST("/inscriptions", ctrl.CreateInscription)
	// Ruta para obtener todas las inscripciones
	engine.GET("/inscriptions", ctrl.GetInscriptions)
	// Ruta para obtener los cursos de un usuario
	engine.GET("/users/:userID/courses", ctrl.GetInscriptionsByUser)
}
