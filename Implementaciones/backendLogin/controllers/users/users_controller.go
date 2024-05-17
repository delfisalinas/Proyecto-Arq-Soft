package users

import (
	usersService "backend/services/users" //le estoy asignando un alias a la carpeta users

	usersDomain "backend/domain/users" //le estoy asignando un alias a la carpeta users

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var loginRequest usersDomain.LoginRequest
	context.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	context.JSON(http.StatusOK, response)

}
