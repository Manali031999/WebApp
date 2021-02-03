package controller

import (
    "fmt"
	"net/http"
	_ "database/sql"
    "bacancy/Employee/packages/db"
)



func Delete(w http.ResponseWriter, r *http.Request) {
    db := db.DbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
    if err != nil {
        panic(err.Error())
        fmt.Fprintf(w,"%v\v",err)
    }
    delForm.Exec(emp)
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}
