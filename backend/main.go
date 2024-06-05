package main

import (
	courseController "backend/controllers/courses"
	inscriptionController "backend/controllers/inscriptions"
	usersController "backend/controllers/users"
	domainInscriptions "backend/domain/inscriptions"
	domainUsers "backend/domain/users"
	routerCourses "backend/router/courses"
	routerInscriptions "backend/router/inscriptions"
	routerUsers "backend/router/users"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "gestion_de_cursos_arqsoft"

	// Conectar a la base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos", err)
	}

	// Migrar los modelos a la base de datos
	if err := db.AutoMigrate(&domainUsers.User{}, &domainInscriptions.Inscription{}); err != nil {
		log.Fatalf("No se pudo migrar a la base de datos", err)
	}

	// Inicializar el enrutador de Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Inicializar los controladores
	usersController := usersController.NewController(db)
	courseController := courseController.NewController(db)
	inscriptionController := inscriptionController.NewController(db)

	// Pasar el controlador al enrutador
	routerUsers.MapUrls(r, usersController)
	routerCourses.MapCourseUrls(r, courseController)
	routerInscriptions.MapInscriptionUrls(r, inscriptionController)

	// Ejecutar el servidor
	if err := r.Run(); err != nil {
		log.Fatalf("Error de correr el servidor", err)
	}
}
