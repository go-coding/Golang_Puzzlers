package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("/Users/jinrong6/develop/sinawww/gitlab/history_bill/log/201912/dianXinYi_20191126_20191225.txt_2019-12.txt")

	if err != nil {
		fmt.Printf(" throw error %s \n", err)
	}

	buf := bufio.NewReader(file)

	for {
		// delim
		//line, err := buf.ReadString('\n')
		line, _, err := buf.ReadLine()
		if err != nil || err == io.EOF {
			break
		} else {
			//fmt.Print(line)
			fmt.Println(line)
		}
	}

}
