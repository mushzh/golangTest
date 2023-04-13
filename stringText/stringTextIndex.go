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

	str := "café"

	// 将 ASCII 字符转换为大写字母
	fmt.Println(strings.ToUpper(str)) // 输出: CAFÉ

	// 将字符串中的所有字符转换为标题格式
	fmt.Println(strings.ToTitle(str)) // 输出: CAFé　⇒　実際：CAFÉ

	// 在 Go 1.2 中，ToUpper 和 ToTitle 函数输出的结果相同，都是 "CAFÉ"
	// 但是，从 Go 1.3 开始，ToTitle 函数将根据语言规则处理字符，因此输出可能会有所不同。

	// 将字符串转换为小写
	res3 := strings.ToLower("ABC")
	fmt.Println(res3) // abc

	// 将单词首字母变为大写, 其它字符不变
	// 单词之间用空格OR特殊字符隔开
	res4 := strings.Title("hello world")
	fmt.Println(res4) // Hello World

	// 按照指定字符串切割原字符串
	// 用,切割字符串
	arr11 := strings.Split("a,b,c", ",")
	fmt.Println(arr11) // [a b c]
	arr21 := strings.Split("ambmc", "m")
	fmt.Println(arr21) // [a b c]
	// 按照指定字符串切割原字符串, 并且指定切割为几份
	// 如果最后一个参数为0, 那么会范围一个空数组
	arr3 := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(arr3) // [a b,c]
	arr4 := strings.SplitN("a,b,c", ",", 0)
	fmt.Println(arr4) // []
	// 按照指定字符串切割原字符串, 切割时包含指定字符串
	arr5 := strings.SplitAfter("a,b,c", ",")
	fmt.Println(arr5) // [a, b, c]

	// 按照指定字符串切割原字符串, 切割时包含指定字符串, 并且指定切割为几份
	arr6 := strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println(arr6) // [a, b,c]
	// 按照空格切割字符串, 多个空格会合并为一个空格处理
	arr7 := strings.Fields("a b c     d")
	fmt.Println(arr7) // [a b c d]
	// 将字符串转换成切片传递给函数之后由函数决定如何切割
	// 类似于IndexFunc
	arr8 := strings.FieldsFunc("aoboc", custom)
	fmt.Println(arr8) // [a b c]
	// 将字符串切片按照指定连接符号转换为字符串
	sce := []string{"aa", "bb", "cc"}
	str12 := strings.Join(sce, "-")
	fmt.Println(str12) // aa-bb-cc

	// 返回count个s串联的指定字符串
	str3 := strings.Repeat("abc", 2)
	fmt.Println(str3) // abcabc
	// 第一个参数: 需要替换的字符串
	// 第二个参数: 旧字符串
	// 第三个参数: 新字符串
	// 第四个参数: 用新字符串 替换 多少个旧字符串
	// 注意点: 传入-1代表只要有旧字符串就替换
	// 注意点: 替换之后会生成新字符串, 原字符串不会受到影响
	str5 := "abcdefabcdefabc"
	str4 := strings.Replace(str5, "abc", "mmm", -1)
	str6 := strings.Replace(str5, "abc", "mmm", 2)
	fmt.Println(str5) // abcdefabcdefabc
	fmt.Println(str4) // mmmdefmmmdefmmm
	fmt.Println(str6) // mmmdefmmmdefabc

	// 去除字符串两端指定字符
	str41 := strings.Trim("!!!abc!!!def!!!", "!")
	fmt.Println(str41) // abc!!!def
	// 去除字符串左端指定字符
	str42 := strings.TrimLeft("!!!abc!!!def!!!", "!")
	fmt.Println(str42) // abc!!!def!!!
	// 去除字符串右端指定字符
	str43 := strings.TrimRight("!!!abc!!!def!!!", "!")
	fmt.Println(str43) // !!!abc!!!def
	// // 去除字符串两端空格
	str44 := strings.TrimSpace("    abc!!!def ")
	fmt.Println(str44) // abc!!!def
	// 按照方法定义规则,去除字符串两端符合规则内容
	str45 := strings.TrimFunc("oooabcooodefooo", custom)

	fmt.Println(str45) // abcooodef
	// 按照方法定义规则,去除字符串左端符合规则内容
	str46 := strings.TrimLeftFunc("oooabcooodefooo", custom)
	fmt.Println(str46) // abcooodefooo
	// 按照方法定义规则,去除字符串右端符合规则内容
	str47 := strings.TrimRightFunc("oooabcooodefooo", custom)
	fmt.Println(str47) // oooabcooodef
	// 取出字符串开头的指定字符串
	str48 := strings.TrimPrefix("lnj-book.avi", "lnj-")
	fmt.Println(str48) // book.avi
	// 取出字符串结尾的指定字符串
	str49 := strings.TrimSuffix("lnj-book.avi", ".avi")
	fmt.Println(str49) // lnj-book
}

func custom(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == 'o' {
		return true
	}
	return false
}
