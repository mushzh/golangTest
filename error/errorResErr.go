package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数が0であってはいけません")
	} else {
		res = a / b
	}
	return
}

func div2(a, b int) (res int) {
	if b == 0 {
		panic("除数はゼロにすることはできません。")
	} else {
		res = a / b
	}
	return
}

func div3(a, b int) (res int) {
	defer func() {
		if err := recover(); err != nil {
			res = -1
			fmt.Println(err) // 除数はゼロにすることはできません
		}
	}()
	if b == 0 {
		panic("除数はゼロにすることはできません")
	} else {
		res = a / b
	}
	return
}

func main() {
	//res, err := div(10, 5)
	res, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	res = div2(10, 2) // 10,0の場合、ここまで止まれた　⇒　panic: 除数はゼロにすることはできません。
	fmt.Println(res)

	var arr = [3]int{1, 3, 4}
	//arr[5] = 666 // エラー発生:invalid argument: index 5 out of bounds [0:3]
	fmt.Println(arr)

	//var res2 = 10 / 0 // エラー発生:invalid operation: division by zero
	//fmt.Println(res2)

	res3 := div3(10, 0)
	fmt.Println(res3) // -1

	panic("异常A") // 相当于return, 后面代码不会继续执行
	panic("异常B") // 多个异常,只有第一个会被捕获
}
