package services

import (
	dtos "backend/DTOs/files"
	"backend/domain/files"
	"backend/domain/inscriptions"
	"errors"

	"gorm.io/gorm"
)

// CreateFile crea un nuevo archivo en la base de datos
func CreateFile(db *gorm.DB, req dtos.CreateFileRequestDTO) (dtos.FileResponseDTO, error) {
	// Verificar si el usuario está inscrito en el curso
	if !IsUserEnrolled(db, req.UserID, req.CourseID) {
		return dtos.FileResponseDTO{}, ErrUserNotEnrolled
	}

	file := files.File{
		Name:     req.Name,
		Content:  req.Content,
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}

	// Intentar guardar el archivo en la base de datos
	if err := db.Create(&file).Error; err != nil {
		return dtos.FileResponseDTO{}, err
	}

	return dtos.FileResponseDTO{
		ID:       file.ID,
		Name:     file.Name,
		Content:  file.Content,
		UserID:   file.UserID,
		CourseID: file.CourseID,
	}, nil
}

// IsUserEnrolled verifica si un usuario está inscrito en un curso
func IsUserEnrolled(db *gorm.DB, userID uint, courseID uint) bool {
	var inscription inscriptions.Inscription
	if err := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&inscription).Error; err != nil {
		return false
	}
	return true
}

// Error de usuario no inscrito
var ErrUserNotEnrolled = errors.New("user is not enrolled in the course")
