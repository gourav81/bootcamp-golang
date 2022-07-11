package Routes

import (
	"E3/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("student", Controllers.GetStudent)
	r.POST("student", Controllers.CreateStudent)
	r.GET("student/:id", Controllers.GetStudentByID)
	r.PUT("student/:id", Controllers.UpdateStudent)
	r.DELETE("student/:id", Controllers.DeleteStudent)
	r.PUT("student/update_marks/:id", Controllers.UpdateStudentMarks)

	return r
}
