package main

import (
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080") //para el frontend se suele usar el puerto 3000
} //usar postman para probar el backend
