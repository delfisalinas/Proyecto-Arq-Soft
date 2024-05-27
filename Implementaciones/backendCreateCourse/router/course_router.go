package router

import (
	"backend/controllers/courses"

	"github.com/gin-gonic/gin"
)

func MapCourseUrls(engine *gin.Engine) {
	engine.POST("/courses", courses.CreateCourse)
	engine.PUT("/courses/:id", courses.UpdateCourse)
	engine.DELETE("/courses/:id", courses.DeleteCourse)
}
