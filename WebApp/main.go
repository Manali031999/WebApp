package main
 
import (
    "encoding/json"
    "html/template"
    "net/http"
    "fmt"
    "io/ioutil"
)
type Data struct{
    Firstname string
    Lastname string
    DOB string
    Email string
    Mobile string
}
func open(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("form.html")
    t.Execute(w, nil)
}
func adddata(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
        var data Data
        data.Firstname = r.FormValue("fname")
        data.Lastname=r.FormValue("lname")
        data.DOB = r.FormValue("dob")
        data.Email = r.FormValue("email")
        data.Mobile = r.FormValue("no")
	    filedata, err := ioutil.ReadFile("form.json")
	    if err != nil {
		    fmt.Println(err)
        }
        var all []Data
        err = json.Unmarshal([]byte(filedata), &all)
        if err != nil {
            fmt.Println("Error JSON Unmarshling for user file")
            fmt.Println(err)
        }
        all = append(all, data)
        file, _ := json.MarshalIndent(all, "", " ")
        _ = ioutil.WriteFile("form.json", file, 0644)
    }
    t,_:=template.ParseFiles("thankyou.html")
    t.Execute(w,nil)
}
func main() {
    http.HandleFunc("/", open)
    http.HandleFunc("/form", adddata)
    http.ListenAndServe(":8080", nil)
    
}
