package main

import (
	"backend/basedatos"
	commentController "backend/controllers/comments"
	courseController "backend/controllers/courses"
	filesController "backend/controllers/files"
	inscriptionController "backend/controllers/inscriptions"
	usersController "backend/controllers/users"
	routerComments "backend/router/comments"
	routerCourses "backend/router/courses"
	routerFiles "backend/router/files"
	routerInscriptions "backend/router/inscriptions"
	routerUsers "backend/router/users"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	db, err := basedatos.Inicializar()
	if err != nil {
		log.Fatalf("No se pudo inicializar la base de datos: %v", err)
	}

	// Inicializar el enrutador de Gin
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Inicializar los controladores
	usersController := usersController.NewController(db)
	courseController := courseController.NewController(db)
	inscriptionController := inscriptionController.NewController(db)
	commentController := commentController.NewController(db)
	filesController := filesController.NewController(db)

	// Pasar el controlador al enrutador
	routerUsers.MapUrls(r, usersController)
	routerCourses.MapCourseUrls(r, courseController)
	routerInscriptions.MapInscriptionUrls(r, inscriptionController)
	routerComments.MapCommentUrls(r, commentController)
	routerFiles.MapFileUrls(r, filesController)

	// Ejecutar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}
