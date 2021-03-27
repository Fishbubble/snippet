// @Desc: 将标准输出重定向,输入到指定文件
// @Author: QianQingnian 2021/3/27 20:55

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, _ := os.OpenFile("./golog.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	golog := &log.Logger{}
	golog = log.New(f, "golog ", log.LstdFlags|log.Lmsgprefix)

	os.Stdout = f

	golog.Println("test golang logger 1")
	fmt.Println("test std logger 1")
	golog.Println("test golang logger 2")
	fmt.Println("test std logger 2")
	golog.Println("test golang logger 3")
	fmt.Println("test std logger 3")
	golog.Println("test golang logger 4")
	fmt.Println("test std logger 4")
}
