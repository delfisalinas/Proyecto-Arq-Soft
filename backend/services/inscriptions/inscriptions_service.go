package inscriptions

import (
	courseDtos "backend/DTOs/courses" // Importar los DTOs de cursos
	dtos "backend/DTOs/inscriptions"
	coursesDomain "backend/domain/courses"
	"backend/domain/inscriptions"
	usersDomain "backend/domain/users"
	"errors"

	"gorm.io/gorm"
)

// CreateInscription crea una nueva inscripción en la base de datos
func CreateInscription(db *gorm.DB, req dtos.CreateInscriptionRequestDTO) (dtos.InscriptionResponseDTO, error) {
	// Verificar si el curso existe
	var course coursesDomain.Course
	if err := db.First(&course, req.CourseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.InscriptionResponseDTO{}, errors.New("course not found")
		}
		return dtos.InscriptionResponseDTO{}, err
	}

	// Verificar si el usuario existe
	var user usersDomain.User
	if err := db.First(&user, req.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.InscriptionResponseDTO{}, errors.New("user not found")
		}
		return dtos.InscriptionResponseDTO{}, err
	}

	// Verificar si ya existe una inscripción con el mismo usuario y curso
	var existingInscription inscriptions.Inscription
	if err := db.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).First(&existingInscription).Error; err == nil {
		return dtos.InscriptionResponseDTO{}, errors.New("inscription already exists")
	}

	// Crear una instancia de inscripción con los datos recibidos
	inscription := inscriptions.Inscription{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}

	// Intentar guardar la inscripción en la base de datos
	if err := db.Create(&inscription).Error; err != nil {
		return dtos.InscriptionResponseDTO{}, err
	}

	return dtos.InscriptionResponseDTO{
		ID:       inscription.ID,
		UserID:   inscription.UserID,
		CourseID: inscription.CourseID,
	}, nil
}

// GetInscriptions obtiene todas las inscripciones de la base de datos
func GetInscriptions(db *gorm.DB) ([]dtos.InscriptionResponseDTO, error) {
	var inscriptions []inscriptions.Inscription
	if err := db.Find(&inscriptions).Error; err != nil {
		return nil, err
	}

	var response []dtos.InscriptionResponseDTO
	for _, inscription := range inscriptions {
		response = append(response, dtos.InscriptionResponseDTO{
			ID:       inscription.ID,
			UserID:   inscription.UserID,
			CourseID: inscription.CourseID,
		})
	}
	return response, nil
}

// GetInscriptionsByUser obtiene todas las inscripciones de un usuario en la base de datos
func GetInscriptionsByUser(db *gorm.DB, userID uint) ([]courseDtos.CourseResponseDTO, error) {
	var inscriptions []inscriptions.Inscription
	if err := db.Where("user_id = ?", userID).Find(&inscriptions).Error; err != nil {
		return nil, err
	}

	var courses []coursesDomain.Course
	for _, inscription := range inscriptions {
		var course coursesDomain.Course
		if err := db.First(&course, inscription.CourseID).Error; err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	var response []courseDtos.CourseResponseDTO
	for _, course := range courses {
		response = append(response, courseDtos.CourseResponseDTO{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Category:     course.Category,
			Duration:     course.Duration,
			InstructorID: course.InstructorID,
		})
	}
	return response, nil
}
