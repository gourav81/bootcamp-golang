package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Student struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	Subject  string `json:"subject"`
	Marks    string `json:"marks"`
}

func (b *User) TableName() string {
	return "user"
}

func (b *Student) TableName() string {
	return "student"
}
