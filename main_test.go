package main

import (
	"E3/Config"
	"E3/Controllers"
	"E3/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateStudent(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := SetUpRouter()

	r.POST("/student", Controllers.CreateStudent)
	student := Models.Student{
		Name:     "T",
		LastName: " Y",
		DOB:      "2/2/2000",
		Address:  "delhi",
		Subject:  "maths",
		Marks:    "50",
	}
	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetStudent(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Student{})
	r := SetUpRouter()
	r.GET("/student", Controllers.GetStudent)
	req, _ := http.NewRequest("GET", "/student", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var students []Models.Student
	json.Unmarshal(w.Body.Bytes(), &students)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, students)
}

//GetStudentByID
func TestGetStudentByID(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Student{})
	r := SetUpRouter()
	r.GET("/student/:id", Controllers.GetStudentByID)
	req, _ := http.NewRequest("GET", "/student/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateStudent(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})

	r := SetUpRouter()
	r.PUT("/student/:id", Controllers.UpdateStudent)
	student := Models.Student{
		Name:     "q",
		LastName: "w",
		DOB:      "1/10/2000",
		Address:  "delhi",
		Subject:  "maths",
		Marks:    "100",
	}
	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest("PUT", "/student/17", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteStudent(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Student{})
	r := SetUpRouter()
	r.DELETE("/student/:id", Controllers.DeleteStudent)
	req, _ := http.NewRequest("DELETE", "/student/3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateStudentMarks(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Student{})
	r := SetUpRouter()
	student := Models.Student{
		Subject: "maths",
		Marks:   "50",
	}
	jsonValue, _ := json.Marshal(student)
	r.PUT("/student/update_marks/:id", Controllers.UpdateStudentMarks)
	req, _ := http.NewRequest("PUT", "/student/update_marks/18", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
