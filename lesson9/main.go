package main

import (
	"fmt"
	"time"
)

func print_channel_basic() {
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// // 送信専用
	// car ch3 chan<- int
	// channelはFIFO

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 3
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 5
	fmt.Println("len", len(ch3))

	i := <- ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <- ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<- ch3)
	fmt.Println("len", len(ch3))

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
}

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func print_channel_and_goroutin() {
	// fmt.Println(<-ch1)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i <100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}