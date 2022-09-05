/*Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.*/

package main

import "fmt"

const pi = 3.14

func main() {
	var circleRadius float32 // радиус круга в сантиметрах

	fmt.Print("Введите диаметр круга = ")
	fmt.Scanln(&circleRadius)
	// площадь круга
	circleArea := (circleRadius) * (circleRadius) * pi

	fmt.Printf("Площадь круга:  %f32 см. кв.\n", circleArea)
}
