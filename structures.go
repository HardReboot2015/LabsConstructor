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
type Results struct {
	ID 				int
	User 			User
	Labs			[]Labs
	Result			int
	Src 			string
	Date 			string
}

type UserLabs struct {
	ID_User			int
	Username		string
	Number 			int
	Theme 			string
	Result			int
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
	X	float64
	Y	float64

}
type Node struct {
	ID_elem		int
	Input		[]Dots
	Output 		[]Dots
}
type Edge struct {
	Node_out	Node
	Node_in		Node
	Dot			string
}
type Graph struct {
	Edges		[]Edge
}
