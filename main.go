package main

import (
	"fmt"
	"github.com/max75025/httprouter"
	"github.com/max75025/open-golang/open"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)
func init(){
	err := checkFolders()
	_, err = openDB()
	if err != nil{
		log.Fatal(err)
	}

}
func main(){
	open.Start("http://localhost:8000/")
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		redirectTarget := "/login"
		http.Redirect(w, r, redirectTarget, 302)

	})

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	})

	router.POST("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		email := r.FormValue("email")
		pass := r.FormValue("password")
		redirectTarget :="/errorLogin/неверная почта или пароль"
		users, err := getUsers()
		if err != nil{
			log.Println(err)
			http.Redirect(w,r,"/errorLogin/serverError", 302)
		}
		if ok, id, is_student := checkLoginData(email, pass, users); email != "" && pass != "" && ok{
			setSession(id, email, is_student, w)
			redirectTarget = "/cabinet/" + strconv.Itoa(id)
		}
		http.Redirect(w,r, redirectTarget, 302)


	})

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logoutHandler(w, r)
	})

	router.GET("/cabinet/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w, r)
		is_student := isStudent(r)
		t, _ := template.ParseFiles("templates/cabinet.html")

			labs, err := getLabs()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				Labs []Labs
				Is_Student int
			}{labs, is_student }
			t.Execute(w, data)



	})

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w, r)
		//if isStudent(r) == 1{http.Redirect(w, r, "/login", 302)}
		t, _ := template.ParseFiles("templates/students.html")
		id := getId(r)
		if isStudent(r) == 0 {
			student, err := getStudents()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				Id       int
				User []User
			}{id, student,
			}
			t.Execute(w, data)

		}

	})

	router.GET("/teachers", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w, r)
		//if isStudent(r) == 1{http.Redirect(w, r, "/login", 302)}
		t, _ := template.ParseFiles("templates/students.html")
		id := getId(r)
		if isStudent(r) == 0 {
			teachers, err := getTeachers()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				Id       int
				User []User
			}{id, teachers,
			}
			t.Execute(w, data)

		}
	})
	router.GET("/add_user", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		checkLogin(w, r)

		id := getId(r)
		data := struct {
			Id int
		}{ id }
		t, _ := template.ParseFiles("templates/add_user.html")
		t.Execute(w, data)
	})

	router.POST("/add_user", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()


		name := r.FormValue("name")
		email := r.FormValue("email")
		is_student, _ := strconv.Atoi(r.FormValue("student"))

		if name != "" && email != "" {

		if err := addUser(User{
				ID: 0,
				Name: name,
				Email: email,
				Is_student: is_student,
			});err != nil{
				log.Println(err)

			}
		}

	})

	router.GET("/addelements", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		t, _ := template.ParseFiles("templates/addelements.html")
		t.Execute(w, nil)
	})

	router.POST("/addelements", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		name := r.FormValue("name")
		src := r.FormValue("src")
		svg := r.FormValue("svg")
		inputx, _ := strconv.ParseFloat(r.FormValue("inputx"), 32)
		inputy, _ := strconv.ParseFloat(r.FormValue("inputy"),32)
		outputx,_ := strconv.ParseFloat(r.FormValue("outputx"), 32)
		outputy, _ := strconv.ParseFloat(r.FormValue("outputx"), 32)
		if err := addElement(Elements{
			ID: 0,
			Name: name,
			Src: src,
			Svg: svg,
			Input: []Dots{
				{inputx, inputy},
			},
			Output: []Dots{{outputx, outputy}},
		}); err != nil{log.Println(err)}

	})
	server := http.Server{
		Addr: ":8000",
		ReadTimeout: time.Duration(30) * time.Second,
		WriteTimeout:  time.Duration(30) * time.Second,
		Handler: router,
	}
	fmt.Println("server listen and serve on port 8000...")
	panic(server.ListenAndServe())
}
