package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}


// customError

// ↓Goで定義されているError interface
// type error interface {
// 	Error() string
// }



type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "Raise custom Error", ErrCode: 123}
}

// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string{
	return fmt.Sprintf("<<%v, %v >>", p.A, p.B)
}

func main() {
	// vs := []Stringfy{
	// 	&Person{Name: "tarou", Age: 21},
	// 	&Car{Number: "12345", Model: "AB-1234"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

	// err := RaiseError()
	// fmt.Println(err.Error())

	// // 型アサーションで取り出せる
	// e, ok := err.(*MyError)
	// if ok {
	// 	fmt.Println(e.ErrCode)
	// }

	p := &Point{100, "ABC"}
	fmt.Println(p)
}