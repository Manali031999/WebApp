package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"time"
	"encoding/json"
)


type AutoGenerated struct {
	Code int `json:"code"`
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Pages int `json:"pages"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
		} `json:"pagination"`
	} `json:"meta"`
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	} `json:"data"`
}


func main(){
	fmt.Println("API Call ... ")
	resp,err:= http.Get("https://gorest.co.in/public-api/categories")
	if err!=nil{
		fmt.Println("ERR----------------->",err)
	}else{
		body,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("ERR---------------->",err)
		}else{
			fmt.Println("BODY--------------->",string(body))
			var auto AutoGenerated
			err=json.Unmarshal(body,&auto)
			if err!=nil{
				fmt.Println("Error in unMarshaling",err)
			}else{
				fmt.Println("-----------------UNMARSHALLED DATA--------------------")
				fmt.Println(auto)
			}

        	file, _ := json.MarshalIndent(auto, "", " ")
        	_ = ioutil.WriteFile("categories.json", file, 0644)
		}
	}
}
