package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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
	_, err = stmt.Exec(elem.Name, elem.Src, elem.Svg, elem.Input, elem.Output)
	if err != nil{return err}
	return nil
}