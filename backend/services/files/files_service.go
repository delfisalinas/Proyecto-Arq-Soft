package services

import (
	dtos "backend/DTOs/files"
	"backend/domain/files"
	"backend/domain/inscriptions"
	"errors"
	"log"

	"gorm.io/gorm"
)

// CreateFile crea un nuevo archivo en la base de datos
func CreateFile(db *gorm.DB, req dtos.CreateFileRequestDTO) (dtos.FileResponseDTO, error) {
	file := files.File{
		Name:     req.Name,
		Content:  []byte(req.Content),
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

// GetFilesByCourseID obtiene los archivos de un curso por su ID
func GetFilesByCourseID(db *gorm.DB, courseID uint) ([]dtos.FileResponseDTO, error) {
	var files []files.File
	if err := db.Where("course_id = ?", courseID).Find(&files).Error; err != nil {
		log.Printf("Error fetching files from DB: %v", err)
		return nil, err
	}

	log.Printf("Files found in DB: %v", files)
	var dtosFiles []dtos.FileResponseDTO
	for _, file := range files {
		dtosFiles = append(dtosFiles, dtos.FileResponseDTO{
			ID:       file.ID,
			Name:     file.Name,
			Content:  file.Content,
			UserID:   file.UserID,
			CourseID: file.CourseID,
		})
	}

	log.Printf("Files converted to DTOs: %v", dtosFiles)
	return dtosFiles, nil
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
