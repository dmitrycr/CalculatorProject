package main

import "fmt"

func main() {
	expression := "(1+2)*3"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("%.2f\n", result)
	}
}

/*
	s := ")("
	s1 := ")"
	if strings.Contains(s, s1) {
		fmt.Println("True")
	} else {
		fmt.Println("error")
	}
*/
