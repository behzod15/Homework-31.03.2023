package main

import (
	"fmt"
	"pr1/funcs"
)
func main() {
	var year, num, num1 int
	var str,str1,str2 string
	fmt.Println("Ввидите год:")
	fmt.Scan(&year)
	str = funcs.Output (year)
	fmt.Println(str)
	fmt.Println("Ввидите число для второго вопроса:")
	fmt.Scan(&num)
	str1 = funcs.IdNum (num)
	fmt.Println(str1)
	fmt.Println("Ввидите число для третьего вопроса:")
	fmt.Scan(&num1)
	str2 = funcs.IdNum1(num1)
	fmt.Println(str2)
}