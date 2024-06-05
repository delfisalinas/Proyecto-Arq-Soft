package users

import (
	dtos "backend/dtos/users"
	usersService "backend/services/users" // asigno un alias a la carpeta users del servicio

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller estructura que contiene la base de datos
type Controller struct {
	db *gorm.DB
}

// NewController función que inicializa el controlador con la base de datos
func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

// Login función que maneja la solicitud de inicio de sesión
func (ctrl *Controller) Login(context *gin.Context) {
	var loginRequest dtos.LoginRequestDTO

	// Recibe el cuerpo JSON de la solicitud y maneja errores
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}

	// Llama al servicio de autenticación de usuarios con los datos de inicio de sesión
	response, err := usersService.Login(ctrl.db, loginRequest)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta del servicio como JSON con un estado HTTP 200 (OK)
	context.JSON(http.StatusOK, response)
}

// Register función que maneja la solicitud de registro
func (ctrl *Controller) Register(context *gin.Context) {
	var registerRequest dtos.RegisterRequestDTO

	// Recibe el cuerpo JSON de la solicitud y maneja errores
	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!!"})
		return
	}

	// Llama al servicio de registro de usuarios con los datos de registro
	response, err := usersService.Register(ctrl.db, registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta del servicio como JSON con un estado HTTP 201 (Created)
	context.JSON(http.StatusCreated, response)
}
