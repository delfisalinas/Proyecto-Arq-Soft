package router

import (
	coursesController "backend/controllers/courses"

	"github.com/gin-gonic/gin"
)

// MapCourseUrls mapea las rutas de cursos a sus respectivos controladores
func MapCourseUrls(engine *gin.Engine, ctrl *coursesController.Controller) {
	// Ruta para crear un nuevo curso
	engine.POST("/courses", ctrl.CreateCourse)
	// Ruta para obtener todos los cursos
	engine.GET("/courses", ctrl.GetCourses)
	// Ruta para obtener un curso por su ID
	engine.GET("/courses/:id", ctrl.GetCourseByID)
	// Ruta para buscar cursos por nombre o categoría
	engine.GET("/search/courses", ctrl.SearchCourses) //http://localhost:8080/search/courses?q=<query> en query va nombre o cat
	// Ruta para actualizar un curso existente
	engine.PUT("/courses/:id", ctrl.UpdateCourse)
	// Ruta para eliminar un curso existente
	engine.DELETE("/courses/:id", ctrl.DeleteCourse)
}
