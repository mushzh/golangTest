package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	// 创建一个正则表达式匹配规则对象
	// reg := regexp.MustCompile(规则字符串)
	// 利用正则表达式匹配规则对象匹配指定字符串
	// 会将所有匹配到的数据放到一个字符串切片中返回
	// 如果没有匹配到数据会返回nil
	// res := reg.FindAllString(需要匹配的字符串, 匹配多少个)
	str := "Hello 李南江 1232"
	reg := regexp.MustCompile("2")
	res := reg.FindAllString(str, -1)
	fmt.Println(res) // [2 2]
	res = reg.FindAllString(str, 1)
	fmt.Println(res) // [2]

	res2 := findPhoneNumber("13554488233")
	fmt.Println(res2) // true
	res2 = findPhoneNumber("03554499311")
	fmt.Println(res2) // false
	res2 = findPhoneNumber("1355449931")
	fmt.Println(res2) // false

	res2 = findEmail("123@qq.com")
	fmt.Println(res2) // true
	res2 = findEmail("ab?de@qq.com")
	fmt.Println(res2) // false
	res2 = findEmail("123@qqcom")
	fmt.Println(res2) // false

	var t time.Time = time.Now()
	fmt.Println(t) // 2023-04-13 20:19:02.0683815 +0900 JST m=+0.007580701

	fmt.Println(t.Year())   // 2023
	fmt.Println(t.Month())  // April
	fmt.Println(t.Day())    // 13
	fmt.Println(t.Hour())   // 20
	fmt.Println(t.Minute()) // 30
	fmt.Println(t.Second()) // 44
}

func findPhoneNumber(str string) bool {
	// 创建一个正则表达式匹配规则对象
	reg := regexp.MustCompile("^1[1-9]{10}")
	// 利用正则表达式匹配规则对象匹配指定字符串
	res := reg.FindAllString(str, -1)
	if res == nil {
		return false
	}
	return true
}

func findEmail(str string) bool {
	reg := regexp.MustCompile("^[a-zA-Z0-9_]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+")
	res := reg.FindAllString(str, -1)
	if res == nil {
		return false
	}
	return true
}
