package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 1.コピー元
	scrPath := "F:\\OVER\\test\\benchi_1.jpg"
	destPath := "F:\\OVER\\test\\benchi_2.png"

	// 2.コピー元のファイルを開けます
	fr, err := os.Open(scrPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3.開けたファイルを閉めます
	defer func() {
		err := fr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 4.buffを取得
	r := bufio.NewReader(fr)

	// 1.コピー先のファイルを作成
	fw, err := os.Create(destPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2.コピー先のファイルを閉めます
	defer func() {
		err := fw.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 3.buffに記入
	w := bufio.NewWriter(fw)
	// 4.システムのcopyメソッドを使ってコピーする
	count, err := io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(count)
	fmt.Println("コピー済み")
}
