package main

import (
	"fmt"
	"runtime"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在唱歌")
		//time.Sleep(time.Millisecond)

		// Gosched使当前go程放弃处理器，以让其它go程运行。
		// 它不会挂起当前go程，因此当前go程未来会恢复执行
		runtime.Gosched()
	}
}
func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在跳舞---")
		//time.Sleep(time.Millisecond)

		runtime.Gosched()
	}
}

func main() {

	// 串行: 必须先唱完歌才能跳舞
	//sing()
	//dance()
	// 并行: 可以边唱歌, 边跳舞
	// 注意点: 主线程不能死, 否则程序就退出了
	go sing()  // 开启一个协程
	go dance() // 开启一个协程
	for {

	}
}
