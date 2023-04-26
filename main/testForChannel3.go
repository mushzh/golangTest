package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义生产者
func producer3(myCh chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了: ", num)
		// 往管道中写入数据
		myCh <- num
		//time.Sleep(time.Millisecond * 500)
	}
	// 生产完毕之后关闭管道
	close(myCh)
	fmt.Println("生产者停止生产")
}

// 定义消费者
func consumer3(myCh <-chan int) {
	// 不断从管道中获取数据, 直到管道关闭位置
	for {
		if num, ok := <-myCh; !ok {
			break
		} else {
			fmt.Println("---消费者消费了", num)
		}
	}
	fmt.Println("消费者停止消费")
}
func main() {
	// 定义缓冲区
	var myCh = make(chan int, 5)
	go producer3(myCh)
	consumer3(myCh)
}

//func main() {
//	// 1.定义一个双向管道
//	var myCh chan int = make(chan int, 5)
//	// 2.将双向管道转换单向管道
//	var myCh2 chan<- int
//	myCh2 = myCh
//	fmt.Println(myCh2)
//	var myCh3 <-chan int
//	myCh3 = myCh
//	fmt.Println(myCh3)
//	// 3.双向管道,可读可写
//	myCh <- 1
//	myCh <- 2
//	myCh <- 3
//	fmt.Println(<-myCh)
//	// 3.只写管道,只能写, 不能读
//	//myCh2 <- 666
//	//fmt.Println(<-myCh2)
//	// 4.指读管道, 只能读,不能写
//	fmt.Println(<-myCh3)
//	//myCh3 <- 666
//	// 注意点: 管道之间赋值是地址传递, 以上三个管道底层指向相同容器
//}
