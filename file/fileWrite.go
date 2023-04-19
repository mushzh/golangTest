package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 1.创建一个文件
	//fp, err := os.Create("F:/OVER/test/mszWrite.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 2.关闭打开的文件
	//defer func() {
	//	err := fp.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//// 2.往文件中写入数据
	//// 注意: Windows换行是\r\n
	//bytes := []byte{'m', 's', 'z', '\r', '\n'}
	//fp.Write(bytes)
	//fp.WriteString("www.it666.com\r\n")
	//fp.WriteString("www.itkk.com\r\n")
	//// 注意: Go语言采用UTF-8编码, 一个中文占用3个字节
	//fp.WriteString("中文占3个字节")
	//
	//// 注意点: 第三个参数在Windows没有效果
	//// -rw-rw-rw- (666)  所有用户都有文件读、写权限。
	////-rwxrwxrwx (777) 所有用户都有读、写、执行权限。
	//// 1.打开文件
	////fp, err := os.OpenFile("F:/OVER/test/mszWrite.txt", os.O_CREATE|os.O_RDWR, 0666)
	//fp, err = os.OpenFile("F:/OVER/test/mszWrite.txt", os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 2.关闭打开的文件
	////defer func() {
	////	err := fp.Close()
	////	if err != nil {
	////		fmt.Println(err)
	////	}
	////}()
	//// 注意点:
	//// 如果O_RDWR模式打开, 被打开文件已经有内容, 会从最前面开始覆盖
	//// 如果O_APPEND模式打开, 被打开文件已经有内容, 会从在最后追加
	//// 3.往文件中写入数据
	//bytes = []byte{'l', 'n', 'j', '\r', '\n'}
	//fp.Write(bytes)
	//fp.WriteString("www.it666.com\r\n")

	// 带缓冲区写入
	// 1.打开文件
	//fp, err := os.OpenFile("F:/OVER/test/mszWrite.txt", os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 2.关闭打开的文件
	//defer func() {
	//	err := fp.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//// 3.创建缓冲区
	//w := bufio.NewWriter(fp)
	//// 4.写入数据到缓冲区
	//bytes := []byte{'l', 'n', 'j', '\r', '\n'}
	//w.Write(bytes)
	//w.WriteString("www.it666.com\r\n")
	//// 5.将缓冲区中的数据刷新到文件
	//w.Flush()

	// 1.写入数据到指定文件
	data := []byte{'l', 'n', 'j', '\r', '\n'}
	err := ioutil.WriteFile("F:/OVER/test/mszWrite.txt", data, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("写入成功")
	}

	info, err := os.Stat("F:/OVER/test/mszWrite.txt")
	if err == nil {
		fmt.Println("文件存在")
		fmt.Println(info.Name())
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("不确定")
	}

	// 1.读取一个文件
	buf, err := ioutil.ReadFile("F:/OVER/test/mszWrite.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.写入读取的数据到另一个文件
	err = ioutil.WriteFile("F:/OVER/test/mszWrite2.txt", buf, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝完成")
}
