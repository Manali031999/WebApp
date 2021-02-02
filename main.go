package main

import (
	"fmt"
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
	"text/template"
)

const (
	host = "localhost"
	port = 5432 
	user = "postgres"
	password = "9825"
	dbname = "form"
)

type Data struct {
	Firstname string
    Lastname string
    DOB string
    Email string
    Mobile string
} 

var t = template.Must(template.ParseGlob("template/*"))

func open(w http.ResponseWriter, r *http.Request){
	t.ExecuteTemplate(w,"form.html",nil)
}

func dbConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error in connection")
		fmt.Println(err)
	}else{
		fmt.Println("Successfully connected!")
		return db
	}
	return nil
}

func insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := dbConn()
		sql := `INSERT INTO form (fname,lname,dob,emailid,number) VALUES ($1,$2,$3,$4,$5)`
		data:= Data{
			Firstname:    r.FormValue("fname"),
			Lastname:     r.FormValue("lname"),
			DOB:  		  r.FormValue("dob"),
			Email:        r.FormValue("email"),
			Mobile: 	  r.FormValue("no"),
		}
		_, err := db.Exec(sql, data.Firstname,data.Lastname,data.DOB,data.Email,data.Mobile)
		if err != nil {
			fmt.Println("error in inserting")
			fmt.Println(err)
		}

		defer db.Close()
	}
	t.ExecuteTemplate(w,"thankyou.html",nil)
}

func getallUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var datas []Data
	rows, err := db.Query("SELECT *FROM form")
	if err != nil {
		fmt.Println("error in selecting")
		fmt.Println(err)
	}
	for rows.Next() {
		var data Data
		err = rows.Scan(&data.Firstname, &data.Lastname, &data.DOB, &data.Email, &data.Mobile)
		if err != nil {
			fmt.Println("error in scanning")
			fmt.Println(err)
		}
		datas = append(datas, data)
	}
	defer db.Close()
	t.ExecuteTemplate(w,"display.html",datas)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("email")
	sql := "DELETE FROM form where emailid = $1"
	db := dbConn()
	_, err := db.Exec(sql,id)
	if err != nil {
		fmt.Println("error in deleting")
		fmt.Println(err)
	}
	defer db.Close()
	http.Redirect(w, r, "/display", 301)
}

func edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("email")
	sql := "SELECT *FROM form WHERE emailid=" + id
	db := dbConn()
	rows := db.QueryRow(sql)
	var data Data
	err := rows.Scan(&data.Firstname, &data.Lastname, &data.DOB, &data.Email, &data.Mobile)
	if err != nil {
		fmt.Println("error in scan")
		fmt.Println(err)
	}
	defer db.Close()
	t.ExecuteTemplate(w,"edit.html",data)
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := dbConn()
		sql := `UPDATE form SET fname = $1,lname = $2,dob = $3 ,number =$5 WHERE emailid = $4;`
		data:= Data{
			Firstname:    r.FormValue("fname"),
			Lastname:     r.FormValue("lname"),
			DOB:  		  r.FormValue("dob"),
			Email:        r.FormValue("email"),
			Mobile: 	  r.FormValue("no"),
		}
		_, err := db.Exec(sql, data.Firstname,data.Lastname,data.DOB,data.Email,data.Mobile)
		if err != nil {
			fmt.Println(err)
		}

		defer db.Close()
	}
	t.ExecuteTemplate(w,"thankyou.html",nil)
}

func main(){
	http.HandleFunc("/",open)
	http.HandleFunc("/form",insert)
	http.HandleFunc("/display.html",getallUser)
	http.HandleFunc("/delete",delete)
	http.HandleFunc("/edit",edit)
	http.HandleFunc("/update",update)
	http.ListenAndServe(":8010",nil)
}