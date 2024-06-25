package comments

// Comment representa la estructura de un comentario
type Comment struct {
	ID       uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"foreignKey"`
	CourseID uint   `gorm:"foreignKey"`
}
