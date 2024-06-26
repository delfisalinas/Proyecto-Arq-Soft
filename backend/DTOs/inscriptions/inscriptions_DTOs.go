package inscriptions

// CreateInscriptionRequestDTO representa la solicitud para crear una inscripción
type CreateInscriptionRequestDTO struct {
	UserID   uint `json:"user_id" binding:"required"`
	CourseID uint `json:"course_id" binding:"required"`
}

// InscriptionResponseDTO representa la respuesta con los detalles de una inscripción
type InscriptionResponseDTO struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}
