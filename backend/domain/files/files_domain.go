package files

// File representa la estructura de un archivo cargado por un usuario
type File struct {
	ID       uint   `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"type:longtext" json:"name"`
	Content  []byte `gorm:"type:blob" json:"content"`
	UserID   uint   `gorm:"foreignKey" json:"userId"`
	CourseID uint   `gorm:"foreignKey" json:"courseId"`
}
