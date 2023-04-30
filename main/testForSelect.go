package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//// create channel
	//var myCh = make(chan int)
	//var exitCh = make(chan bool)
	//
	//// generate data
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		myCh <- i
	//		time.Sleep(time.Second)
	//	}
	//	exitCh <- true
	//}()
	//
	//// read data
	//for {
	//	fmt.Println("read data source go")
	//	select {
	//	case num := <-myCh:
	//		fmt.Println("readed:", num)
	//	case <-exitCh:
	//		// break // 無駄、スキップされたものがselects
	//		return
	//	}
	//	fmt.Println("---------------")
	//}

	// create channel
	myCh := make(chan int, 5)
	exitCh := make(chan bool)

	// generate data
	go func() {
		for i := 0; i < 10; i++ {
			myCh <- i
			time.Sleep(time.Second * 3)
		}
	}()

	// read data
	go func() {
		for {
			select {
			case num := <-myCh:
				fmt.Println(num)
			case <-time.After(time.Second * 2):
				exitCh <- true
				runtime.Goexit()
			}
		}
	}()

	<-exitCh
	fmt.Println("Program exit")

}
