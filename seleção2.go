package main

import "fmt"

func main() {

	var x int64
	var y int64
	var z int64

	fmt.Print("Quais são os números?")
	fmt.Scan(&x, &y, &z)

	menor := x
	if y < menor {
		menor = y
	}
	if z < menor {
		menor = z
	}
	fmt.Println("O menor número é:", menor)

}
