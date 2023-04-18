package main

import (
	"fmt"
	"os"
)

func main() {
	// 1.创建一个文件
	fp, err := os.Create("F:/OVER/test/mszWrite.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 2.关闭打开的文件
	defer func() {
		err := fp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 2.往文件中写入数据
	// 注意: Windows换行是\r\n
	bytes := []byte{'m', 's', 'z', '\r', '\n'}
	fp.Write(bytes)
	fp.WriteString("www.it666.com\r\n")
	fp.WriteString("www.itkk.com\r\n")
	// 注意: Go语言采用UTF-8编码, 一个中文占用3个字节
	fp.WriteString("中文占3个字节")
}
