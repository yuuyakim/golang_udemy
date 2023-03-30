package main

import "fmt"

func print_calc()  {
	fmt.Println(1+2)
	fmt.Println("ABC" + "DE")

	fmt.Println(1-2)
	fmt.Println(1*2)
	fmt.Println(4/2)
	fmt.Println(9%2)

	n := 0
	n += 4
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)
}

func print_compare()  {
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(4 <= 8)
	fmt.Println(4 >= 8)
}

func print_bool_calc() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)

}


func main() {
	print_calc()
	print_compare()
	print_bool_calc()
}