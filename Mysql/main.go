package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"encoding/json"
	"fmt"
)

var db *gorm.DB
var err error
const DNS="root:@/form?parseTime=true"

type User struct{
	gorm.Model
	FirstName string   `json:"firstname`
	LastName string    `json:"lastname"`
	Email string       `json:"email"`
}

func InitialMigration(){
	db,err=gorm.Open(mysql.Open(DNS),&gorm.Config{})
	if err!=nil{
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	var user User
	db.First(&user,params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}

func initializeRouter(){
	r:=mux.NewRouter()
	r.HandleFunc("/users",GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8907",r))
}

func main(){
	InitialMigration()
	initializeRouter()
}
