package comments

import (
	"backend/domain/comments"
	coursesDomain "backend/domain/courses"
	inscriptions "backend/domain/inscriptions"
	usersDomain "backend/domain/users"
	dtos "backend/dtos/comments"
	"errors"

	"gorm.io/gorm"
)

// isUserEnrolledInCourse verifica si un usuario está inscrito en un curso
func isUserEnrolledInCourse(db *gorm.DB, userID, courseID uint) (bool, error) {
	var inscription inscriptions.Inscription
	err := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&inscription).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // Usuario no está inscrito en el curso
		}
		return false, err
	}
	return true, nil // Usuario está inscrito en el curso
}

// CreateComment crea un nuevo comentario en la base de datos
func CreateComment(db *gorm.DB, req dtos.CreateCommentRequestDTO) (dtos.CommentResponseDTO, error) {
	// Verificar si el usuario existe
	var user usersDomain.User
	if err := db.First(&user, req.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.CommentResponseDTO{}, errors.New("user not found")
		}
		return dtos.CommentResponseDTO{}, err
	}

	// Verificar si el curso existe
	var course coursesDomain.Course
	if err := db.First(&course, req.CourseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.CommentResponseDTO{}, errors.New("course not found")
		}
		return dtos.CommentResponseDTO{}, err
	}

	// Verificar si el usuario está inscrito en el curso
	isEnrolled, err := isUserEnrolledInCourse(db, req.UserID, req.CourseID)
	if err != nil {
		return dtos.CommentResponseDTO{}, err
	}
	if !isEnrolled {
		return dtos.CommentResponseDTO{}, errors.New("user not enrolled in the course")
	}

	// Crear una instancia de comentario con los datos recibidos
	comment := comments.Comment{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Content:  req.Content,
	}

	// Intentar guardar el comentario en la base de datos
	if err := db.Create(&comment).Error; err != nil {
		return dtos.CommentResponseDTO{}, err
	}

	return dtos.CommentResponseDTO{
		ID:       comment.ID,
		UserID:   comment.UserID,
		CourseID: comment.CourseID,
		Content:  comment.Content,
	}, nil
}

// GetCommentsByCourse obtiene todos los comentarios de un curso por su ID
func GetCommentsByCourse(db *gorm.DB, courseID uint) ([]dtos.CommentResponseDTO, error) {
	var comments []comments.Comment
	// Buscar todos los comentarios para el curso dado
	if err := db.Where("course_id = ?", courseID).Find(&comments).Error; err != nil {
		return nil, err
	}

	// Mapear los comentarios a DTOs de respuesta
	var commentsDTO []dtos.CommentResponseDTO
	for _, comment := range comments {
		commentsDTO = append(commentsDTO, dtos.CommentResponseDTO{
			ID:       comment.ID,
			Content:  comment.Content,
			UserID:   comment.UserID,
			CourseID: comment.CourseID,
		})
	}

	return commentsDTO, nil
}
