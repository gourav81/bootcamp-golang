package Routes

import (
	"E3/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp2 := r.Group("/student-api")
	{
		grp2.GET("student", Controllers.GetStudent)
		grp2.POST("student", Controllers.CreateStudent)
		grp2.GET("student/:id", Controllers.GetStudentByID)
		grp2.PUT("student/:id", Controllers.UpdateStudent)
		grp2.DELETE("student/:id", Controllers.DeleteStudent)
		grp2.PUT("student/update_marks/:id", Controllers.UpdateStudentMarks)
	}
	return r
}
