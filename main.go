package main

import (
    "fmt"
    "database/sql"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Name  string
    City string
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/form")
    if err != nil {
        panic(err.Error())
	}
	fmt.Println("Connected Successfully")
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM employee")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "index.html", res)
    defer db.Close()
}
func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "new.html", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
    }
    tmpl.ExecuteTemplate(w, "edit.html", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id:=r.FormValue("id")
        insForm, err := db.Prepare("INSERT INTO employee (id,name, city) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id,name, city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city, id)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    fmt.Println("Server started on: http://localhost:9080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":9080", nil)
}