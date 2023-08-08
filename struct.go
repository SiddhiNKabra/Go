package main

import "fmt"

type College struct {
	CollegeName string
	CollegeCity string
}

type Student struct {
	College // college struct embedded into student struct(composition)
	StudentName   string
	StudentYear   int
	StudentRollno int
}

func main() {
	fmt.Println("Struct")

	s := Student{}
	s.CollegeName = "Pune Institute of Computer Technology"
	s.CollegeCity = "Pune"
	s.StudentName = "Jane Doe"
	s.StudentRollno = 21125
	s.StudentYear = 2

	fmt.Println(s)

}
