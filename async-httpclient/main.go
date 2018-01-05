package main

import (
	"fmt"
	"os"
	"time"

	"github.com/painterdrown/service-computing/async-httpclient/ah"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("错误: 缺乏参数 1/2\n")
		println("")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		begin := time.Now()
		ah.Start1()
		elapsed := time.Since(begin) / 1000000
		fmt.Printf("运行时间: %dms\n", elapsed)
		break
	case "2":
		begin := time.Now()
		ah.Start2()
		elapsed := time.Since(begin) / 1000000
		fmt.Printf("运行时间: %dms\n", elapsed)
		break
	default:
		fmt.Println("错误: 参数值只能是 1/2")
	}

}
