package basedatos

import (
	"time"

	domainComments "backend/domain/comments"
	domainCourses "backend/domain/courses"
	files "backend/domain/files"
	domainInscriptions "backend/domain/inscriptions"
	domainUsers "backend/domain/users"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(retries int) (*gorm.DB, error) {
	dbUser := "root"
	dbPass := "rootpassword"
	dbHost := "mysql"
	dbPort := "3306"
	dbName := "gestion_de_cursos_arqsoft"

	var db *gorm.DB
	var err error

	for i := 0; i < retries; i++ {
		// Intentar conectar a la base de datos
		db, err = connect(dbUser, dbPass, dbHost, dbPort, dbName)
		if err == nil {
			break // Conexión exitosa, salir del bucle
		}
		log.Printf("Error al conectar con la base de datos (intento %d/%d): %v", i+1, retries, err)
		time.Sleep(5 * time.Second) // Esperar antes de intentar nuevamente
	}

	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar con la base de datos después de %d intentos: %v", retries, err)
	}

	// Migrar los modelos a la base de datos
	if err := db.AutoMigrate(
		&domainUsers.User{},
		&domainCourses.Course{},
		&domainInscriptions.Inscription{},
		&files.File{},
		&domainComments.Comment{}); err != nil {
		return nil, fmt.Errorf("no se pudo migrar a la base de datos: %v", err)
	}

	return db, nil
}

func connect(dbUser, dbPass, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Inicializar() (*gorm.DB, error) {
	const maxRetries = 5 // Número máximo de intentos de conexión
	return ConnectDatabase(maxRetries)
}
