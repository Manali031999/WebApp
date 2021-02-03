package controller

import (
    "fmt"
	"net/http"
    "bacancy/Employee/packages/db"
    "bacancy/Employee/packages/structure"
)

func Index(w http.ResponseWriter, r *http.Request) {
    db := db.DbConn()
    selDB, err := db.Query("SELECT * FROM employee")
    if err != nil {
        panic(err.Error())
        fmt.Fprintf(w,"%v\v",err)
    }
    emp := structure.Employee{}
    res := []structure.Employee{}
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
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "index.html", res)
    defer db.Close()
}