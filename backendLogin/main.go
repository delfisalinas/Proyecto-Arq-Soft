package main

import (
	usersController "backend/controllers/users"
	"backend/domain/users"
	"backend/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Conectar a la base de datos
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrar los modelos a la base de datos
	db.AutoMigrate(&users.User{})

	// Inicializar el enrutador de Gin
	r := gin.Default()

	// Inicializar el controlador de usuarios
	controller := usersController.NewController(db)

	// Pasar el controlador al enrutador
	router.MapUrls(r, controller)

	// Ejecutar el servidor
	r.Run()
}

/* OLD CODE
	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080") //para el frontend se suele usar el puerto 3000
} //usar postman para probar el backend
*/
