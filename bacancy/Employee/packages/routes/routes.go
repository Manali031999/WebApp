package routes

import (
	"fmt"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "bacancy/Employee/packages/controller"
)


func Routes(){
	fmt.Println("Server started on: http://localhost:9080")
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/edit", controller.Edit)	
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)
	http.ListenAndServe(":9080", nil)
}
