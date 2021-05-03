package rpc

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	//像服务器发送请求
	ch <- req

	fmt.Println("client send")
	//等待服务器返回  select中得case 通道操作会同时开启，哪个先返回就执行哪个
	select {
	case ack := <-ch:	//接收到服务器返回数据
		fmt.Println("client end")
		return ack, nil
	case <-time.After(time.Second):	 //超时
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {
		//接收客户端请求
		data := <-ch

		//打印接收到得数据
		fmt.Println("server  received:", data)

		//向客户端反馈已收到
		ch <- "test merge"
		ch <- "rogerxxxxxxxxxxx"
	}
}

