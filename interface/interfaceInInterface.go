package main

import "fmt"

// サンプル1
//	type A interface {
//		fna()
//	}
//
//	type B interface {
//		fnb()
//	}
//
//	type C interface {
//		A // 嵌入A接口
//		B // 嵌入B接口
//		fnc()
//	}
//
// type Person struct{}
//
//	func (p Person) fna() {
//		fmt.Println("实现A接口中的方法")
//	}
//
//	func (p Person) fnb() {
//		fmt.Println("实现B接口中的方法")
//	}
//
//	func (p Person) fnc() {
//		fmt.Println("实现C接口中的方法")
//	}
//
//	func main() {
//		p := Person{}
//		p.fnb() // 实现B接口中的方法
//		p.fna() // 实现A接口中的方法
//		p.fnc() // 实现C接口中的方法
//	}

// サンプル2
//type aer interface {
//	fna()
//}
//type ber interface {
//	aer
//	fnb()
//}
//
//// Person实现了超集接口所有方法
//type Person struct{}
//
//func (p Person) fna() {
//	fmt.Println("实现A接口中的方法")
//}
//func (p Person) fnb() {
//	fmt.Println("实现B接口中的方法")
//}
//
//// Student实现了子集接口所有方法
//type Student struct{}
//
//func (p Student) fna() {
//	fmt.Println("实现A接口中的方法")
//}

// サンプル3
type studier interface {
	read()
}
type Person struct {
	name string
	age  int
}

func (p Person) read() {
	fmt.Println(p.name, "勉強中")
}

func main() {
	// サンプル2
	//var i ber
	//// 子集接口变量不能转换为超集接口变量
	////i = Student{}
	//fmt.Println(i)
	//var j aer
	//// 超集接口变量可以自动转换成子集接口变量,
	//j = Person{}
	//fmt.Println(j)

	// サンプル3
	// 1.定义一个接口类型变量
	var s studier
	// 2.用接口类型变量接收实现了接口的结构体
	s = Person{"msz", 34}
	//s.name = "zs" // 报错, 由于s是接口类型, 所以不能访问属性
	fmt.Println(s) // {msz 34}
	//var p Person
	//// インタフェース型を元の型に強制的に変換することはできません
	//p = Person(s)

	// ok-idiom
	if p, ok := s.(Person); ok {
		p.name = "lxo"
		fmt.Println(p) // {lxo 34}
	}

	// type switch
	switch p := s.(type) {
	case Person:
		p.name = "zs"
		fmt.Println(p) // {zs 34}
	default:
		fmt.Println("Person型ではない")
	}
}
