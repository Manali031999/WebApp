package controller

import (
    "fmt"
    "net/http"
    "bacancy/Employee/packages/db"
)


func Update(w http.ResponseWriter, r *http.Request) {
    db := db.DbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
        if err != nil {
            panic(err.Error())
            fmt.Fprintf(w,"%v\v",err)
        }
        insForm.Exec(name, city, id)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}