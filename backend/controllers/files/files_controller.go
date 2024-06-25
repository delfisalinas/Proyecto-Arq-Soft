package controllers

import (
	files "backend/DTOs/files"
	services "backend/services/files"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller estructura para manejar las solicitudes de archivos
type Controller struct {
	db *gorm.DB // Conexión a la base de datos
}

// NewController crea una nueva instancia del controlador de archivos
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

// CreateFile maneja la carga de un nuevo archivo
func (ctrl *Controller) CreateFile(context *gin.Context) {
	var req files.CreateFileRequestDTO
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	// Decodificar el contenido del archivo desde base64
	decodedContent, err := base64.StdEncoding.DecodeString(string(req.Content))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode file content: " + err.Error()})
		return
	}

	req.Content = decodedContent

	file, err := services.CreateFile(ctrl.db, req)
	if err != nil {
		if err == services.ErrUserNotEnrolled {
			context.JSON(http.StatusForbidden, gin.H{"error": "User is not enrolled in the course"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file: " + err.Error()})
		}
		return
	}

	context.JSON(http.StatusOK, file)
}
