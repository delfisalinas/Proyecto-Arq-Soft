package courses

// Course representa la estructura de un curso
type Course struct {
	ID           uint   `gorm:"primaryKey;AUTO_INCREMENT"` // Identificador único del curso
	Name         string `gorm:"type:longtext"`             // Nombre del curso
	Description  string `gorm:"type:longtext"`             // Descripción del curso
	Category     string `gorm:"type:longtext"`             // Categoría del curso
	Duration     string `gorm:"type:longtext"`             // Duración del curso
	InstructorID uint   `gorm:"foreignKey"`                // ID del instructor que imparte el curso
}

// CreateCourseRequest representa la solicitud para crear un curso
type CreateCourseRequest struct {
	Name         string `gorm:"type:longtext" binding:"required"`
	Description  string `gorm:"type:longtext" binding:"required"`
	Category     string `gorm:"type:longtext" binding:"required"`
	Duration     string `gorm:"type:longtext" binding:"required"`
	InstructorID uint   `gorm:"foreignKey" binding:"required"`
}

// UpdateCourseRequest representa la solicitud para actualizar un curso
type UpdateCourseRequest struct {
	ID           uint   `gorm:"primaryKey;AUTO_INCREMENT" binding:"required"`
	Name         string `gorm:"type:longtext"`
	Description  string `gorm:"type:longtext"`
	Category     string `gorm:"type:longtext"`
	Duration     string `gorm:"type:longtext"`
	InstructorID uint   `gorm:"foreignKey"`
}
