package main

import "fmt"

func main() {
	// 1.声明一个管道
	var mych chan int
	// 2.初始化一个管道
	mych = make(chan int, 3)
	// 3.查看管道的长度和容量
	fmt.Println("长度是", len(mych), "容量是", cap(mych)) // 长度是 0 容量是 3
	// 4.像管道中写入数据
	mych <- 666
	fmt.Println("长度是", len(mych), "容量是", cap(mych)) // 长度是 1 容量是 3
	// 5.取出管道中写入的数据
	num := <-mych
	fmt.Println("num = ", num)                            // num =  666
	fmt.Println("长度是", len(mych), "容量是", cap(mych)) // 长度是 0 容量是 3

	// 注意点: 管道中只能存放声明的数据类型, 不能存放其它数据类型
	//mych<-3.14
	// 注意点: 管道中如果已经没有数据,
	// 并且检测不到有其它协程再往管道中写入数据, 那么再取就会报错
	//num = <-mych
	//fmt.Println("num = ", num)
	// 注意点: 如果管道中数据已满, 再写入就会报错
	mych <- 666
	mych <- 777
	mych <- 888
	mych <- 999
}
