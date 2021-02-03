package controller

import (
    "fmt"
	"bacancy/Employee/packages/db"
	"net/http"
)



func Insert(w http.ResponseWriter, r *http.Request) {
    db := db.DbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id:=r.FormValue("id")
        insForm, err := db.Prepare("INSERT INTO employee (id,name, city) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
            fmt.Fprintf(w,"%v\v",err)
        }
        insForm.Exec(id,name, city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}