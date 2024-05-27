<<<<<<< HEAD
package controllers
=======
package course

import (
	"backendCreateCourse/services/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var req course.CreateCourseRequest
	context.BindJSON(&req)
	course := course.CreateCourse(req)
	context.JSON(http.StatusOK, course)
}

func UpdateCourse(context *gin.Context) {
	var req course.UpdateCourseRequest
	context.BindJSON(&req)
	course := course.UpdateCourse(req)
	context.JSON(http.StatusOK, course)
}

func DeleteCourse(context *gin.Context) {
	courseID := context.Param("id")
	success := course.DeleteCourse(courseID)
	if success {
		context.Status(http.StatusOK)
	} else {
		context.Status(http.StatusInternalServerError)
	}
}
>>>>>>> 36d29088c862b09a3e807827a501ee053dba7df2
