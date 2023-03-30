package main

import (
	"fmt"
	"time"
	// "os"
	"strconv"
)

func print_if() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")

		x := 0
		if x := 2; true {
			fmt.Println(x)
		}

		fmt.Println(x)
	}
}

func print_error()  {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}

func print_for() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("loop")
	}
}

func print_for_if() {
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
}

func print_for_formal() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func print_for_array() {
	arr := [3]int{1, 2, 3}
	for i:=0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func print_for_range() {
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func print_for_slice() {
	sl := []string{"python", "php", "go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
}

func print_for_map() {
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func print_switch() {
	// n := 5
	// switch n {
	// case 1,2:
	// 	fmt.Println("1 or 2")

	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I dont know")
	// }

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}
}

func anything(a interface{})  {
	fmt.Println(a)
}

func type_assertion()  {
	var x interface{} = 3
	// i, isInt := x.(int)
	// fmt.Println(i, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is String")
	} else {
		fmt.Println("i dont know")
	}
}

func switch_type_assertion(x interface{}) {
	// var x interface{} = 3

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case float64:
		fmt.Println(v, "float64")
	default:
		fmt.Println(v, "i dont know")
	} 
}

func print_for_label() {
// Loop:
// 	for {
// 		for {
// 			for {
// 				fmt.Println("Start")
// 				break Loop
// 			}
// 			fmt.Println("dont execute")
// 		}
// 		fmt.Println("dont execute")
// 	}
// 	fmt.Println("END")
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("dont execute")
	}
}

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("Start")
}

func RunDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	// print_for()
	// print_for_if()
	// print_for_formal()
	// print_for_array()
	// print_for_range()
	// print_for_slice()
	// print_for_map()
	// print_switch()
	// anything("aaaa")
	// anything(1)

	// type_assertion()

	// switch_type_assertion("hogehoge")
	// switch_type_assertion(12)

	// print_for_label()

	// TestDefer()

	// defer func() {
	// 	fmt.Println(1)
	// 	fmt.Println(2)
	// 	fmt.Println(3)
	// }()

	// RunDefer()

	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// file.Write([]byte("hello"))

	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }()
	// panic("runtime error")
	// fmt.Println("Start")

	// go sub()
	// go sub()

	// for {
	// 	fmt.Println("main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	fmt.Println("main")
}