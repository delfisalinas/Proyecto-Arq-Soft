package comments

// CreateCommentRequestDTO representa la solicitud para crear un comentario
type CreateCommentRequestDTO struct {
	UserID   uint   `json:"user_id" binding:"required"`
	CourseID uint   `json:"course_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// CommentResponseDTO representa la respuesta con los detalles de un comentario
type CommentResponseDTO struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	CourseID uint   `json:"course_id"`
	Content  string `json:"content"`
}
