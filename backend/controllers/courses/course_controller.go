package courses

import (
	dtos "backend/DTOs/courses"               // Importar los DTOs de cursos
	coursesService "backend/services/courses" // Importar el servicio de cursos
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller estructura para manejar las solicitudes de cursos
type Controller struct {
	db *gorm.DB // Conexión a la base de datos
}

// NewController crea una nueva instancia del controlador de cursos
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

// CreateCourse maneja la creación de un nuevo curso
func (ctrl *Controller) CreateCourse(context *gin.Context) {
	var req dtos.CreateCourseRequestDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	course, err := coursesService.CreateCourse(ctrl.db, req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, course)
}

// GetCourses maneja la obtención de todos los cursos
func (ctrl *Controller) GetCourses(context *gin.Context) {
	courses, err := coursesService.GetCourses(ctrl.db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list courses: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, courses)
}

// GetCourseByID maneja la obtención de un curso por su ID
func (ctrl *Controller) GetCourseByID(context *gin.Context) {
	courseID := context.Param("id")
	course, err := coursesService.GetCourseByID(ctrl.db, courseID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get course: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":            course.ID,
		"name":          course.Name,
		"description":   course.Description,
		"category":      course.Category,
		"duration":      course.Duration,
		"instructor_id": course.InstructorID,
	})
}

// UpdateCourse maneja la actualización de un curso existente
func (ctrl *Controller) UpdateCourse(context *gin.Context) {
	var req dtos.UpdateCourseRequestDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	course, err := coursesService.UpdateCourse(ctrl.db, req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, course)
}

// DeleteCourse maneja la eliminación de un curso
func (ctrl *Controller) DeleteCourse(context *gin.Context) {
	courseID := context.Param("id")
	if err := coursesService.DeleteCourse(ctrl.db, courseID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course: " + err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

// SearchCourses maneja la búsqueda de cursos por nombre o categoría
func (ctrl *Controller) SearchCourses(context *gin.Context) {
	query := context.Query("q")
	courses, err := coursesService.SearchCourses(ctrl.db, query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search courses: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, courses)
}

func (ctrl *Controller) GetCoursesByUserID(context *gin.Context) {
	userIDParam := context.Param("user_id")
	fmt.Printf("Received userIDParam: %s\n", userIDParam)
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	courses, err := coursesService.GetCoursesByUserID(ctrl.db, uint(userID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user courses: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, courses)
}
