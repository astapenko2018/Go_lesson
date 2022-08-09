package main

import "fmt"

func main() {
	var a, b float32
	fmt.Println("Длины сторон прямоугольника:")

	fmt.Print("a = ")
	fmt.Scanln(&a)

	fmt.Print("b = ")
	fmt.Scanln(&b)

	fmt.Println("Площадь:", a*b)
	
}
