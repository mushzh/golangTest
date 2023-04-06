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

func main() {
	//res, err := div(10, 5)
	res, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	res = div2(10, 0)
	fmt.Println(res)

	var arr = [3]int{1, 3, 4}
	arr[5] = 666 // エラー発生:invalid argument: index 5 out of bounds [0:3]
	fmt.Println(arr)

	var res2 = 10 / 0 // エラー発生:invalid operation: division by zero
	fmt.Println(res2)
}
