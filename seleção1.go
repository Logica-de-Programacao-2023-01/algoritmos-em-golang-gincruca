package main

import "fmt"

func main() {

	var x float64
	var y float64

	fmt.Print("Qual é o primeiro valor?")
	fmt.Scan(&x)
	fmt.Print("Qual é o segundo valor?")
	fmt.Scan(&y)

	if x > y {
		fmt.Println(x, "é maior que", y)
	} else if x < y {
		fmt.Println(y, "é maior que", x)
	} else {
		fmt.Println("São iguais")
	}
}
