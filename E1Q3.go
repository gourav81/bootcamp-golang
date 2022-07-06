// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Employee interface {
	salary()
}

type FullTime struct {
	basic int
}

type Contract struct {
	basic int
}

type Freelance struct {
	basic int
	hour  int
}

func (employee FullTime) salary() {
	fmt.Println(employee.basic * 28)
}

func (employee Contract) salary() {
	fmt.Println(employee.basic * 28)
}

func (employee Freelance) salary() {
	fmt.Println(employee.basic * employee.hour)
}

func main() {
	freelanceEmployee := Freelance{10, 20}
	freelanceEmployee.salary()
}
