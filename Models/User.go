//Models/User.go
package Models

import (
	"E3/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

func GetAllStudent(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}

func UpdateStudentMarks(student *Student, id string, subject string, marks string) (err error) {

	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	student.Marks = marks
	Config.DB.Save(student)
	return nil
}
