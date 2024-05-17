package courses

type Course struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID string `json:"instructor_id"`
}

type CreateCourseRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID string `json:"instructor_id"`
}

type UpdateCourseRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID string `json:"instructor_id"`
}
