package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1+3"
	for i := range str {
		operator := string(str[i])
		if operator == "+" || operator == "-" || operator == "*" {
			fmt.Println(operator)
			res := strings.Split(str, operator)
			fmt.Println(res)
			for i = 0; i < len(res)-1; i++ {
				if operator == "+" {
					a, _ := strconv.Atoi(res[i])
					b, _ := strconv.Atoi(res[i+1])
					fmt.Println(a + b)
				} else if operator == "-" {
					a, _ := strconv.Atoi(res[i])
					b, _ := strconv.Atoi(res[i+1])
					fmt.Println(a - b)
				} else if operator == "*" {
					a, _ := strconv.Atoi(res[i])
					b, _ := strconv.Atoi(res[i+1])
					fmt.Println(a * b)
				}
			}
		}

	}
}
