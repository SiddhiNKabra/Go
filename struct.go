package main

import "fmt"

type College struct {
	CollegeName string
	CollegeCity string
}

type Student struct {
	College       // college struct embedded into student struct(composition)
	StudentName   string
	StudentYear   int
	StudentRollno int
}
type Department struct {
	Student
	DepartmentName string
}

func main() {
	fmt.Println("Struct")

	d := Department{}
	d.CollegeName = "Pune Institute of Computer Technology"
	d.CollegeCity = "Pune"
	d.StudentName = "Jane Doe"
	d.StudentRollno = 21125
	d.StudentYear = 2
	d.DepartmentName = "Computer"

	fmt.Println(d)

}
