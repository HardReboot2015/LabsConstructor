package main

type User struct  {
	ID 				int
	Name 			string
	Email 			string
	Password 		string
	Is_student 		int
}

type Labs struct{
	ID 				int
	Number 			int
	Theme 			string
	Task 			string
	Time_to_complete int
	Access 			int
}

type Elements struct {
	ID 				int
	Name 			string
	Src 			string
	Svg 			string
	Input			[]Dots
	Output 			[]Dots
}
type Dots struct {
	X 				float64
	Y				float64
}