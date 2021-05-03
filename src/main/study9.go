package main

import (
	"fmt"
	"rpc"
	"time"
)

func main() {
	/*//执行并发程序
	go running()
	//接受命令行输入，不做任何事情
	var input string
	fmt.Scanln(&input)*/


	//实例RPC
	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	// 并发执行服务器逻辑
	go rpc.RPCServer(ch)

	// 客户端请求数据和接收数据
	recv, err := rpc.RPCClient(ch, "hi")
	if err != nil {
		// 发生错误打印
		fmt.Println(err)
	} else {
		// 正常接收到数据
		fmt.Println("client received", recv)
	}
	close(ch)
}

func running() {
	var times int
	//构建一个无限循环
	for true {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}
