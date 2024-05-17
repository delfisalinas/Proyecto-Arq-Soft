package course

import "backend/domain/course"

func CreateCourse(request course.CreateRequest) course.CreateResponse {
	//validar contra la base de datos

	return course.CreateResponse{
		Message: "Course created",
	}
}
