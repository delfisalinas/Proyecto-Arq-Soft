package comments

import (
	dtos "backend/dtos/comments"
	commentsService "backend/services/comments"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller estructura para manejar las solicitudes de comentarios
type Controller struct {
	db *gorm.DB
}

// NewController crea una nueva instancia del controlador de comentarios
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

// CreateComment maneja la creación de un nuevo comentario
func (ctrl *Controller) CreateComment(context *gin.Context) {
	var req dtos.CreateCommentRequestDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Formato invalido de la request: " + err.Error()})
		return
	}
	comment, err := commentsService.CreateComment(ctrl.db, req)
	if err != nil {
		switch err.Error() {
		case "user not found":
			context.JSON(http.StatusBadRequest, gin.H{"error": "El usuario no existe"})
		case "course not found":
			context.JSON(http.StatusBadRequest, gin.H{"error": "El curso no existe"})
		case "user not enrolled in the course":
			context.JSON(http.StatusForbidden, gin.H{"error": "El usuario no está inscripto en el curso"})
		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el comentario: " + err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, comment)
}

// GetCommentsByCourse maneja la obtención de todos los comentarios de un curso
func (ctrl *Controller) GetCommentsByCourse(context *gin.Context) {
	courseID, err := strconv.ParseUint(context.Param("courseID"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	comments, err := commentsService.GetCommentsByCourse(ctrl.db, uint(courseID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los comentarios del curso: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, comments)
}
