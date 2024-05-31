package main

import (
	usersController "backend/controllers/users"
	"backend/domain/users"
	"backend/router"
	"fmt"
	"log"

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
	if err := db.AutoMigrate(&users.User{}); err != nil {
		log.Fatalf("No se pudo migrar a la base de datos", err)
	}

	// Inicializar el enrutador de Gin
	r := gin.Default()

	// Inicializar el controlador de usuarios
	controller := usersController.NewController(db)

	// Pasar el controlador al enrutador
	router.MapUrls(r, controller)

	// Ejecutar el servidor
	if err := r.Run(); err != nil {
		log.Fatalf("Error de correr el servidor", err)
	}
}
