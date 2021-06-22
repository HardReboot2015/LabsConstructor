package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"strconv"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request)(userName string){
	if cookie, err := request.Cookie("session"); err == nil{
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil{
			userName = cookieValue["name"]

		}
	}


	return userName
}
func isStudent(request *http.Request)(is_student int){
	if cookie, err := request.Cookie("session"); err == nil{
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil{
			is_student, _ = strconv.Atoi(cookieValue["is_student"])

		}
	}
	return is_student
}
func getId(request *http.Request)(id int){
	if cookie, err := request.Cookie("session"); err == nil{
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil{
			id, _ = strconv.Atoi(cookieValue["id"])

		}
	}

	return id

}
func setSession(id int,userName string,is_student int, response http.ResponseWriter)  {
	value := map[string]string{
		"id" : strconv.Itoa(id),
		"name" : userName,
		"is_student": strconv.Itoa(is_student),

	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil{
		cookie := &http.Cookie{
			Name : "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter)  {
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,

	}
	http.SetCookie(response, cookie)
}

func checkLoginData(mail string, password string, users []User)(bool, int, int){
	for _, u := range users{
		if mail == u.Email{
			if password == u.Password {return true, u.ID, u.Is_student}
		}
	}
	return false, 0, 0
}


func checkLogin(w http.ResponseWriter, r *http.Request) {
	if getUserName(r) == ""{
		http.Redirect(w,r, "/login", 302)
	}

}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/login", 302)
}
