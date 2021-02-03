package controller

import (
    "fmt"
	"text/template"
	"net/http"
    "bacancy/Employee/packages/db"
    "bacancy/Employee/packages/structure"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func Edit(w http.ResponseWriter, r *http.Request) {
    db := db.DbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
        fmt.Fprintf(w,"%v\v",err)
    }
    emp := structure.Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
            fmt.Fprintf(w,"%v\v",err)
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
    }
    tmpl.ExecuteTemplate(w, "edit.html", emp)
    defer db.Close()
}