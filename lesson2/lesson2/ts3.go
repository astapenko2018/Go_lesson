package main


import (
	"fmt"
	"strconv"
)

func hundredsTensUnits(num int64) (string, error) {
	fmt.Printf("Ваш номер: %d \n", num)

	var result string

	if num >= 100 && num <= 999 {
		mystring := strconv.FormatInt(int64(num), 10)

		for i := 0; i < len(mystring); i++ {
			switch i {
			case 0:
				result += "сотни = " + string(mystring[i]) + "\n"
			case 1:
				result += "десятки = " + string(mystring[i]) + "\n"
			case 2:
				result += "еденицы = " + string(mystring[i])

			}
		}
	}
	return result, nil
}

func getNumber() {
	var mynum int64
	var properNumber bool = false
	for !properNumber {
		fmt.Print("Введите число между 100 и 999 : ")
		fmt.Scanln(&mynum)
		properNumber, _ = checkNumber(mynum)
	}
	htu, _ := hundredsTensUnits(mynum)
	fmt.Println(htu)
	fmt.Println("")

}

func checkNumber(x int64) (bool, error) {
	var result bool = false
	if x >= 100 && x <= 999 {
		result = true
	}
	return result, nil
}

func main() {

	fmt.Println("Ввод 3-х значного числа")
	getNumber()

}
