package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	name  string
	score int
}

func main() {
	s := Student{Person{"msz", 34}, "lxo", 99}

	// 継承構造で同じ名前のメソッドがある場合、最も近い親のメソッドが使用されます。
	fmt.Println(s.Person.name) // msz
	fmt.Println(s.name)        // lxo
	fmt.Println(s.Person.age)  // 34
	fmt.Println(s.age)         // 34
	fmt.Println(s.score)       // 99

	var err error = fmt.Errorf("これはエラーメッセージ")
	fmt.Println(err)
	var err2 error = errors.New("これはエラーメッセージ2")
	fmt.Println(err2)
}
