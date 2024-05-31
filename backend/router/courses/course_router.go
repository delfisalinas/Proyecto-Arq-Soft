package router

import (
	coursesController "backend/controllers/courses"

	"github.com/gin-gonic/gin"
)

// MapCourseUrls mapea las rutas de cursos a sus respectivos controladores
func MapCourseUrls(engine *gin.Engine, ctrl *coursesController.Controller) {
	// Ruta para crear un nuevo curso
	engine.POST("/courses", ctrl.CreateCourse)
	// Ruta para actualizar un curso existente
	engine.PUT("/courses/:id", ctrl.UpdateCourse)
	// Ruta para eliminar un curso existente
	engine.DELETE("/courses/:id", ctrl.DeleteCourse)
}
