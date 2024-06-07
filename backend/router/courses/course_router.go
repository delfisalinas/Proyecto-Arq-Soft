package router

import (
	coursesController "backend/controllers/courses"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// MapCourseUrls mapea las rutas de cursos a sus respectivos controladores
func MapCourseUrls(engine *gin.Engine, ctrl *coursesController.Controller) {
	// Agrupar las rutas protegidas por el middleware
	auth := engine.Group("/courses")
	auth.Use(middleware.AuthMiddleware()) // Aplica el middleware de autenticación

	// Ruta para crear un nuevo curso, solo accesible por administradores
	auth.POST("", middleware.AdminMiddleware(), ctrl.CreateCourse)
	// Ruta para actualizar un curso existente, solo accesible por administradores
	auth.PUT("/:id", middleware.AdminMiddleware(), ctrl.UpdateCourse)
	// Ruta para eliminar un curso existente, solo accesible por administradores
	auth.DELETE("/:id", middleware.AdminMiddleware(), ctrl.DeleteCourse)

	// Rutas públicas
	// Ruta para obtener todos los cursos
	engine.GET("/courses", ctrl.GetCourses)
	// Ruta para obtener un curso por su ID
	engine.GET("/courses/:id", ctrl.GetCourseByID)
	// Ruta para buscar cursos por nombre o categoría
	engine.GET("/search/courses", ctrl.SearchCourses) //http://localhost:8080/search/courses?q=<query> en query va nombre o cat

	engine.GET("/user/:user_id/courses", ctrl.GetCoursesByUserID)
}
