package main

import (
	"fmt"
	f"fmt" // import package の命名が可能
	."fmt" // 省略できる
	"golang_udemy/lesson13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " "  + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
		fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	f.Println(foo.ReturnMin())
	Println("hoge")
	fmt.Println(appName())
	fmt.Println(Do("AAAAAA"))
}