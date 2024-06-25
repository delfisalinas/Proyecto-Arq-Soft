package comments

import (
	commentsController "backend/controllers/comments"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// MapCommentUrls mapea las rutas de comentarios a sus respectivos controladores
func MapCommentUrls(engine *gin.Engine, ctrl *commentsController.Controller) {
	// Agrupar las rutas protegidas por el middleware
	auth := engine.Group("/comments")
	auth.Use(middleware.AuthMiddleware())

	// Ruta para crear un nuevo comentario
	auth.POST("/", ctrl.CreateComment)
	// Ruta para obtener todos los comentarios de un curso
	auth.GET("/courses/:courseID", ctrl.GetCommentsByCourse)
}
