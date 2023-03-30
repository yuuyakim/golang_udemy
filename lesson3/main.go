package main

import (
	"fmt"
	// "strconv"
)

func print_int()  {

	// integer
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)
	// fmt.Println(i2 + i)

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))
}

func print_float()  {
	// float
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println((pinf))

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}

func print_bool() {
	// boolean

	var t, f bool = true, false
	fmt.Println(t, f)
}

func print_string() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
		test`)
	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))
}

func print_byte()  {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}

func print_array()  {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr2[2] = "c"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// println(arr5)

	fmt.Println(len(arr1))
}

func print_interface_nil() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "H"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}

func print_type() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	// var s string = "100"
	// fmt.Printf("s = %T\n", s)
	// i, _ := strconv.Atoi(s)
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	var h string = "HelloWorld"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}

func main() {
	print_int()
	print_float()
	print_bool()
	print_string()
	print_byte()
	print_array()
	print_interface_nil()
	print_type()
}