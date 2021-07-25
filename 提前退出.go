package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		go func() {
			func() {
				fmt.Println("child inner function")
				return //这是返回当前函数
				//os.Exit(-1) //退出进程
				//runtime.Goexit()//退出子当前go程
			}()
			fmt.Println("child process end")
			fmt.Println("go 222")
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("go 111")
	}()
	fmt.Println("this father process")
	time.Sleep(3 * time.Second)
	fmt.Println("over")
}