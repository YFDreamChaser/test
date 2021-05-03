package main

import (
	"base"
	_ "cls1"
	_ "cls2"
)

func main() {
	//根据字符串动态创建一个Class1实例
	c1 := base.Create("Class1")
	c1.Do()
	c2 := base.Create("Class2")
	c2.Do()
}
