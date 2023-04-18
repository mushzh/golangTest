package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 1.打开一个文件
	// 注意: 文件不存在不会创建, 会报错
	// 注意: 通过Open打开只能读取, 不能写入
	fp, err := os.Open("F:/OVER/test/msz.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fp)
	}

	// 2.关闭一个文件
	defer func() {
		err = fp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 3.读取指定指定字节个数据
	// 注意点: \n也会被读取进来
	//buf := make([]byte, 50)
	//count, err := fp.Read(buf)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(count)
	//	fmt.Println(string(buf))
	//}

	// 4.读取文件中所有内容, 直到文件末尾为止
	//buf := make([]byte, 10)
	//for {
	//	count, err := fp.Read(buf)
	//	// 注意: 这行代码要放到判断EOF之前, 否则会出现少读一行情况
	//	fmt.Print(string(buf[:count]))
	//	if err == io.EOF {
	//		break
	//	}
	//}

	// 4.读取文件中所有内容, 直到文件末尾为止
	r := bufio.NewReader(fp)
	for {
		buf, err := r.ReadBytes('\n')
		//buf, err := r.ReadString('\n')
		fmt.Print(string(buf))
		if err == io.EOF {
			break
		}
	}
}
