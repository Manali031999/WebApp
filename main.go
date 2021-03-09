package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Data struct {
	Op1      string `json:"op1"`
	Operator string `json:"operator"`
	Op2      string `json:"op2"`
	Result   string `json:"result"`
}

func Calculations(w http.ResponseWriter, r *http.Request) {
	var data Data
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	num1, err := strconv.ParseFloat(data.Op1, 64)
	if err != nil {
		log.Fatal("error in string to int")
	}

	num2, err := strconv.ParseFloat(data.Op2, 64)
	if err != nil {
		log.Fatal("error in string to int")
	}

	var ans float64
	switch data.Operator {
	case "+":
		ans = float64(num1) + float64(num2)
	case "-":
		ans = float64(num1) - float64(num2)
	case "*":
		ans = float64(num1) * float64(num2)
	case "/":
		ans = float64(num1) / float64(num2)
	case "!":
		ans = factorial(num1)
	}
	fmt.Println(ans)
	data.Result = fmt.Sprint(ans)
	w.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(data)
	resp, _ := json.Marshal(data)
	f, err := os.OpenFile("history.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(resp)
	if err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
	w.Write(resp)
	return
}

var factVal float64 = 1.0 // uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.
var i float64 = 1.0
var n float64

/*     function declaration        */
func factorial(n float64) float64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1.0; i <= n; i++ {
			factVal *= i // mismatched types int64 and int
		}

	}
	fmt.Println(factVal)
	return factVal /* return from function*/
}

func main() {
	router := mux.NewRouter()

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	router.PathPrefix("/static/").Handler(fs)
	http.Handle("/static/", router)
	router.HandleFunc("/Get", Calculations).Methods("POST")
	// fs := http.FileServer(http.Dir("static/"))
	// router.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", router)
}
