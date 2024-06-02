package courses

import (
	"backend/domain/courses"
	dtos "backend/dtos/courses"
	"errors"

	"gorm.io/gorm"
)

// CreateCourse crea un nuevo curso en la base de datos
func CreateCourse(db *gorm.DB, req dtos.CreateCourseRequestDTO) (dtos.CourseResponseDTO, error) {
	// Verificar si ya existe un curso con el mismo nombre y el mismo instructor
	var existingCourse courses.Course
	if err := db.Where("name = ? AND instructor_id = ?", req.Name, req.InstructorID).First(&existingCourse).Error; err == nil {
		return dtos.CourseResponseDTO{}, errors.New("course already exists")
	}

	// Crear una instancia del curso con los datos recibidos
	course := courses.Course{
		Name:         req.Name,
		Description:  req.Description,
		Category:     req.Category,
		Duration:     req.Duration,
		InstructorID: req.InstructorID,
	}
	// Intentar guardar el curso en la base de datos
	if err := db.Create(&course).Error; err != nil {
		return dtos.CourseResponseDTO{}, err
	}
	return dtos.CourseResponseDTO{
		ID:           course.ID,
		Name:         course.Name,
		Description:  course.Description,
		Category:     course.Category,
		Duration:     course.Duration,
		InstructorID: course.InstructorID,
	}, nil
}

// GetCourses obtiene todos los cursos de la base de datos
func GetCourses(db *gorm.DB) ([]dtos.CourseResponseDTO, error) {
	var courses []courses.Course
	// Buscar todos los cursos en la base de datos
	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}

	// Convertir los cursos encontrados a DTOs
	var coursesDTO []dtos.CourseResponseDTO
	for _, course := range courses {
		coursesDTO = append(coursesDTO, dtos.CourseResponseDTO{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Category:     course.Category,
			Duration:     course.Duration,
			InstructorID: course.InstructorID,
		})
	}
	return coursesDTO, nil
}

// GetCourseByID obtiene un curso por ID de la base de datos
func GetCourseByID(db *gorm.DB, courseID string) (dtos.CourseResponseDTO, error) {
	var course courses.Course
	// Buscar el curso por ID en la base de datos
	if err := db.First(&course, courseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.CourseResponseDTO{}, errors.New("course not found")
		}
		return dtos.CourseResponseDTO{}, err
	}
	return dtos.CourseResponseDTO{
		ID:           course.ID,
		Name:         course.Name,
		Description:  course.Description,
		Category:     course.Category,
		Duration:     course.Duration,
		InstructorID: course.InstructorID,
	}, nil
}

// UpdateCourse actualiza un curso existente en la base de datos
func UpdateCourse(db *gorm.DB, req dtos.UpdateCourseRequestDTO) (dtos.CourseResponseDTO, error) {
	var course courses.Course
	// Buscar el curso por ID en la base de datos
	if err := db.First(&course, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.CourseResponseDTO{}, errors.New("course not found")
		}
		return dtos.CourseResponseDTO{}, err
	}

	// Actualizar los campos del curso si se proporcionaron nuevos valores
	if req.Name != "" {
		course.Name = req.Name
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	if req.Category != "" {
		course.Category = req.Category
	}
	if req.Duration != "" {
		course.Duration = req.Duration
	}
	if req.InstructorID != 0 {
		course.InstructorID = req.InstructorID
	}

	// Guardar los cambios en la base de datos
	if err := db.Save(&course).Error; err != nil {
		return dtos.CourseResponseDTO{}, err
	}
	return dtos.CourseResponseDTO{
		ID:           course.ID,
		Name:         course.Name,
		Description:  course.Description,
		Category:     course.Category,
		Duration:     course.Duration,
		InstructorID: course.InstructorID,
	}, nil
}

// DeleteCourse elimina un curso de la base de datos
func DeleteCourse(db *gorm.DB, courseID string) error {
	// Intentar eliminar el curso por ID
	if err := db.Delete(&courses.Course{}, courseID).Error; err != nil {
		return err
	}
	return nil
}
