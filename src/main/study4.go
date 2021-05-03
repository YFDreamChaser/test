package main

import "fmt"

func main() {
	//跳出指定循环以及开启指定循环
OuterLoop: //break结束循环标签(任意取名)
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
