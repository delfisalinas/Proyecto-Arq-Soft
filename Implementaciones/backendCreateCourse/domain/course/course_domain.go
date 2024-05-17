package course

type CreateRequest struct {
	CourseName string `json:"coursename"`
}

type CreateResponse struct {
	Message string `json:"message"`
}
