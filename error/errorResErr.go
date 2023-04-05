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

func main() {
	//res, err := div(10, 5)
	res, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
