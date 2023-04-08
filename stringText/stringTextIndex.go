package main

import (
	"fmt"
	"strings"
)

func main() {
	// Go语言编码方式是UTF-8,在UTF-8中一个汉字占3个字节
	str1 := "lnj"
	fmt.Println(len(str1)) // 3
	str2 := "公号：代码情缘"
	fmt.Println(len(str2)) // 21

	arr1 := []byte(str2)
	fmt.Println(len(arr1)) // 21
	for _, v := range arr1 {
		fmt.Printf("%c", v) // å¬å·ï¼ä»£çæç¼
	}

	arr2 := []rune(str2)
	fmt.Println(len(arr2)) // 7
	for _, v := range arr2 {
		fmt.Printf("%c", v) // 公号：代码情缘
	}

	// 查找`字符`在字符串中第一次出现的位置, 找不到返回-1
	res := strings.IndexByte("hello 李南江", 'l')
	fmt.Println(res) // 2
	// 查找`汉字`OR`字符`在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexRune("hello 李南江", '李')
	fmt.Println(res) // 6
	res = strings.IndexRune("hello 李南江", 'l')
	fmt.Println(res) // 2
	// 查找`汉字`OR`字符`中任意一个在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexAny("hello 李南江", "wml")
	fmt.Println(res) // 2
	// 会把wmhl拆开逐个查找, w、m、h、l只要任意一个被找到, 立刻停止查找
	res = strings.IndexAny("hello 李南江", "wmhl")
	fmt.Println(res) // 0
	// 查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.Index("hello 李南江", "llo")
	fmt.Println(res) // 2
	// 会把lle当做一个整体去查找, 而不是拆开
	res = strings.Index("hello 李南江", "lle")
	fmt.Println(res) // -1
	// 可以查找字符也可以查找汉字
	res = strings.Index("hello 李南江", "李")
	fmt.Println(res) // 6
	// 会将字符串先转换为[]rune, 然后遍历rune切片逐个取出传给自定义函数
	// 只要函数返回true,代表符合我们的需求, 既立即停止查找　⇒　ここに'o'がtrue
	res = strings.IndexFunc("hello 李南江", custom)
	fmt.Println(res) // 4
	// 倒序查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.LastIndex("hello 李南江", "l")
	fmt.Println(res) // 3

	// 查找`子串`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.Index函数
	resB := strings.Contains("hello 李南江", "llo")
	fmt.Println(resB) // true
	resB = strings.Contains("hello 李南江", "李")
	fmt.Println(resB) // true

	// 查找`汉字`OR`字符`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexRune函数
	resB = strings.ContainsRune("hello 李南江", 'l')
	fmt.Println(resB) // true
	resB = strings.ContainsRune("hello 李南江", '李')
	fmt.Println(resB) // true

	// 查找`汉字`OR`字符`中任意一个在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexAny函数
	resB = strings.ContainsAny("hello 代码情缘", "wmhl")
	fmt.Println(resB) // true

	// 判断字符串是否已某个字符串开头
	resB = strings.HasPrefix("lnj-book.avi", "lnj")
	fmt.Println(resB) // true

	// 判断字符串是否已某个字符串结尾
	resB = strings.HasSuffix("lnj-book.avi", ".avi")
	fmt.Println(resB) // true

	// 比较两个字符串大小, 会逐个字符地进行比较ASCII值
	// 第一个参数 > 第二个参数 返回 1
	// 第一个参数 < 第二个参数 返回 -1
	// 第一个参数 == 第二个参数 返回 0
	resC := strings.Compare("bcd", "abc")
	fmt.Println(resC) // 1
	resC = strings.Compare("bcd", "bdc")
	fmt.Println(resC) // -1
	resC = strings.Compare("bcd", "bcd")
	fmt.Println(resC) // 0

	// 判断两个字符串是否相等, 可以判断字符和中文
	// 判断时会忽略大小写进行判断
	res2 := strings.EqualFold("abc", "def")
	fmt.Println(res2) // false
	res2 = strings.EqualFold("abc", "abc")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("abc", "ABC")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("代码情缘", "代码情缘")
	fmt.Println(res2) // true
}

func custom(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == 'o' {
		return true
	}
	return false
}
