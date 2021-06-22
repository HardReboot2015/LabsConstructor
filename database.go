package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
const photoDirPath = "./photo/"
const dbDirPath = "./database/"
const passPath = "./database/passwords.txt"
func checkFolders()error{
	err := os.MkdirAll(photoDirPath, os.ModePerm)
	if err != nil {return err}
	/*err = os.MkdirAll(contentListDirPath, os.ModePerm)
	if err != nil {return err}*/
	err = os.MkdirAll(dbDirPath, os.ModePerm)
	if err != nil {return err}
	return nil
}
func openDB()(*sql.DB, error){
	connstr := "root:1107@tcp(127.0.0.1:3306)/labsdb"
	db, err := sql.Open("mysql", connstr)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return db, nil
}


func getUsers()([]User, error){
	var result []User
	db, err := openDB()
	if err != nil{return nil, err}
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {return nil, err}
	for rows.Next(){
		u:= User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email,&u.Password, &u.Is_student)
		if err != nil{return nil, err}
		result = append(result, u)
	}
	return result, nil
}

func getStudents()([]User, error){
	var result []User
	db, err := openDB()
	if err != nil{return nil, err}
	rows, err := db.Query("SELECT id, name, email, password FROM user WHERE is_student = 1")
	if err != nil { return nil, err}
	for rows.Next (){
		u:= User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password )
		if err != nil{return nil, err}
		result = append(result, u)
	}
	return result, nil
}
func getTeachers() ([]User, error) {
	var result []User
	db, err := openDB()
	if err != nil{return nil, err}
	rows, err := db.Query("SELECT id, name, email, password FROM user WHERE is_student = 0")
	if err != nil { return nil, err}
	for rows.Next (){
		u:= User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password )
		if err != nil{return nil, err}
		result = append(result, u)
	}
	return result, nil

}

func getLabs()([]Labs, error){
	var result []Labs
	db, err := openDB()
	if err != nil{return nil, err}
	rows, err := db.Query("SELECT id, number, theme, access FROM labs")
	if err != nil {return nil, err}
	for rows.Next(){
		l := Labs{}
		err = rows.Scan(&l.ID, &l.Number, &l.Theme, &l.Access)
		if err != nil{return nil, err}
		result = append(result, l)
	}
	return result, nil
}


func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func addUser(user User) error {
	db, err := openDB()
	if err != nil{return err}
	rand.Seed(time.Now().UnixNano())
	password := randSeq(10)
	stmt, err := db.Prepare("INSERT INTO user(name, email, password, is_student) VALUES (?,?,?,?)")
	if err != nil {return err}
	_, err = stmt.Exec(user.Name, user.Email, password, user.Is_student)
	if err != nil{return err}
	return nil
}

func addElement(elem Elements) error{
	db, err := openDB()
	if err != nil {return err}
	stmt , err := db.Prepare("INSERT INTO element(name, src, svg, inputs, outputs) VALUES (?, ?, ?, ?, ?)")
	if err != nil {return err}

	input, err := json.Marshal(elem.Input)
	output, err := json.Marshal(elem.Output)

	_, err = stmt.Exec(elem.Name, elem.Src, elem.Svg, string(input), string(output))
	if err != nil{return err}
	return nil
}

func userLabs(student int)([]UserLabs, string, error)  {
	var result []UserLabs
	var name string
	db, err := openDB()
	if err != nil { return nil, "", err }
	rows, err := db.Query("SELECT user.id, user.name,labs.number, labs.theme, result.result  FROM result, labs, user WHERE result.id_user = user.id AND id_lab = labs.id AND user.id = ?", student)
	if err != nil {return nil, "", err}

	for rows.Next(){
		l := UserLabs{}
		err = rows.Scan(&l.ID_User, &l.Username, &l.Number, &l.Theme, &l.Result)
		if err != nil{return nil,"",  err}
		result = append(result, l)
		name = l.Username
	}

	return result, name, nil
}

//Работа с элементами

func getElements() ([]Elements, error) {
	db, err := openDB()
	var result []Elements
	if err != nil{return nil, err}
	rows, err := db.Query("SELECT id, name, src, svg, inputs, outputs FROM element")

	if err != nil{return nil, err}
	for rows.Next(){
		elem := Elements{}
		var inputdots []Dots
		var outputdots []Dots
		var inputjson []byte
		var outputjson []byte
		err = rows.Scan(&elem.ID, &elem.Name, &elem.Src, &elem.Svg, &inputjson, &outputjson)
		if err != nil {return nil, err}
		err = json.Unmarshal(inputjson, &inputdots)
		if err != nil {return nil, err}
		err = json.Unmarshal(outputjson, &outputdots)
		if err != nil {return nil, err}
		elem.Input = inputdots
		elem.Output = outputdots
		result = append(result, elem)

	}
	return result, nil
}

//func addGraph(ems []Elements)  {
//
//
//}

