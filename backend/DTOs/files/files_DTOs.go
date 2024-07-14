package files

// CreateFileRequestDTO representa la solicitud para cargar un archivo
type CreateFileRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Content  string `json:"content" binding:"required"`
	UserID   uint   `json:"userId" binding:"required"`
	CourseID uint   `json:"courseId" binding:"required"`
}

// FileResponseDTO representa la respuesta de un archivo
type FileResponseDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Content  []byte `json:"content"`
	UserID   uint   `json:"userId"`
	CourseID uint   `json:"courseId"`
}
