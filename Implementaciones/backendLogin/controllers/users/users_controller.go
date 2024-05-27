package users

import (
	usersService "backend/services/users" //le estoy asignando un alias a la carpeta users

	usersDomain "backend/domain/users" //le estoy asignando un alias a la carpeta users

	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var loginRequest usersDomain.LoginRequest

	// Recibe el cuerpo JSON de la solicitud y maneja errores
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Llama al servicio de autenticación de usuarios con los datos de inicio de sesión
	response := usersService.Login(loginRequest)

	// Verificar si la respuesta contiene un error
	if response.Error != "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": response.Error})
		return
	}

	// Devolver la respuesta del servicio como JSON con un estado HTTP 200 (OK)
	context.JSON(http.StatusOK, response)
}

var db *gorm.DB //variable global de la base de datos (VALIDAR!!!!)

func Register(context *gin.Context) {
	var registerRequest usersDomain.RegisterRequest

	// Recibe el cuerpo JSON de la solicitud y maneja errores
	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Llama al servicio de registro de usuarios con los datos de registro
	response, err := usersService.Register(db, registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta del servicio como JSON con un estado HTTP 201 (Created)
	context.JSON(http.StatusCreated, response)
}
