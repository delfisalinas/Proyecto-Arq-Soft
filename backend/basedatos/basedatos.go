package basedatos

import (
	domainComments "backend/domain/comments"
	domainCourses "backend/domain/courses"
	domainInscriptions "backend/domain/inscriptions"
	domainUsers "backend/domain/users"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Inicializar() (*gorm.DB, error) {
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
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	// Migrar los modelos a la base de datos
	if err := db.AutoMigrate(
		&domainUsers.User{},
		&domainCourses.Course{},
		&domainInscriptions.Inscription{},
		&domainComments.Comment{}); err != nil {
		log.Fatalf("No se pudo migrar a la base de datos: %v", err)
		return nil, err
	}

	return db, nil
}
