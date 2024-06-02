package inscriptions

import (
	dtos "backend/dtos/inscriptions"
	inscriptionsService "backend/services/inscriptions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller estructura para manejar las solicitudes de inscripciones
type Controller struct {
	db *gorm.DB
}

// NewController crea una nueva instancia del controlador de inscripciones
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

// CreateInscription maneja la creación de una nueva inscripción
func (ctrl *Controller) CreateInscription(context *gin.Context) {
	var req dtos.CreateInscriptionRequestDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Formato invalido de la request: " + err.Error()})
		return
	}
	inscription, err := inscriptionsService.CreateInscription(ctrl.db, req)
	if err != nil {
		switch err.Error() {
		case "course not found":
			context.JSON(http.StatusBadRequest, gin.H{"error": "El curso no existe"})
		case "user not found":
			context.JSON(http.StatusBadRequest, gin.H{"error": "El usuario no existe"})
		case "inscription already exists":
			context.JSON(http.StatusConflict, gin.H{"error": "El usuario ya está inscrito en este curso"})
		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la inscripción: " + err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, inscription)
}

// GetInscriptions maneja la obtención de todas las inscripciones
func (ctrl *Controller) GetInscriptions(context *gin.Context) {
	inscriptions, err := inscriptionsService.GetInscriptions(ctrl.db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al mostrar las inscripciones: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, inscriptions)
}

// GetInscriptionsByUser maneja la obtención de todas las inscripciones de un usuario
func (ctrl *Controller) GetInscriptionsByUser(context *gin.Context) {
	userID, err := strconv.ParseUint(context.Param("userID"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	courses, err := inscriptionsService.GetInscriptionsByUser(ctrl.db, uint(userID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los cursos del usuario: " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, courses)
}
