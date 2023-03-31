package main

import (
	"fmt"
	// "time"
	// "time"
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

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
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

func channel_close() {
	ch1 := make(chan int, 2)

	ch1 <- 1
	close(ch1)
	// ch1 <- 1

	// fmt.Println(<-ch1)
	// チャネルの受信　戻りの2つ目はchannelの疎通状態 or channelにデータがあるかどうか
	i, ok := <-ch1
	fmt.Println(i, ok)

	i2, ok := <-ch1
	fmt.Println(i2, ok)
}

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go reciever(ch1)
	// go reciever(ch2)

	// i := 0
	// for i <100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Millisecond)
	// 	i++
	// }

	// channel_close()

	// ch1 := make(chan int, 3)

	// go reciever2("1.goroutin", ch1)
	// go reciever2("2.goroutin", ch1)
	// go reciever2("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }
	// close(ch1)
	// time.Sleep(3 * time.Second)

	// ch1 := make(chan int, 3)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3

	// // channelはcloseしていない状態で値を取得しようとすると、データの受信待ち状態から抜けられずDeadlockとなるのでcloseしてからLoopを回す
	// close(ch1)
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どっちでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}
}
