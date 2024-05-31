package courses

// Course representa la estructura de un curso
type Course struct {
	ID           uint   `json:"id" gorm:"primaryKey"` // Identificador único del curso
	Name         string `json:"name"`                 // Nombre del curso
	Description  string `json:"description"`          // Descripción del curso
	Category     string `json:"category"`             // Categoría del curso
	Duration     string `json:"duration"`             // Duración del curso
	InstructorID uint   `json:"instructor_id"`        // ID del instructor que imparte el curso
}

// CreateCourseRequest representa la solicitud para crear un curso
type CreateCourseRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Duration     string `json:"duration" binding:"required"`
	InstructorID uint   `json:"instructor_id" binding:"required"`
}

// UpdateCourseRequest representa la solicitud para actualizar un curso
type UpdateCourseRequest struct {
	ID           uint   `json:"id" binding:"required"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID uint   `json:"instructor_id"`
}
