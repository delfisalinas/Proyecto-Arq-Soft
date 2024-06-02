package inscriptions

// Inscription representa la estructura de una inscripción
type Inscription struct {
	ID       uint `gorm:"primaryKey;AUTO_INCREMENT"` // Identificador único de la inscripción
	UserID   uint `gorm:"foreignKey"`                // ID del usuario que se inscribe
	CourseID uint `gorm:"foreignKey"`                // ID del curso al que se inscribe
}
