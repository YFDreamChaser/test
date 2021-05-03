package main

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

func main() {
	//测试函数中参数传递--> 所有结构是一块新内存，只有指针不会复制
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)
	out := passByValue(in)
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)


	//字符串的链式处理器的设计思想
	list := []string {
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	chain := []func(string) string {
		removePrefix,  //传入函数名
		strings.TrimSpace,
		strings.ToUpper,
	}
	//处理
	StringProccess(list, chain)
	//输出
	for _, str := range list {
		fmt.Println(str)
	}


	//匿名函数,后面的括号就是调用
	fmt.Println(func(data int) int {
		return data
	}(100))
	//匿名函数作为回调函数
	visit("helloword", func(str string) {
		fmt.Println(str)
	})


	//函数类型实现接口
	//结构体实现
	var invoker1 Invoker = new(Struct)
	invoker1.Call("Hello")
	//函数体实现  -->  将匿名函数转换为FuncCaller类型
	var invoker2 Invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker2.Call("hello")


	//闭包--> 引用了外部变量的匿名函数
	str := "hello word"
	foo := func() {
		//匿名函数中访问str
		str = "hello wln"
	}
	//调用
	foo()
	fmt.Println(str)

	//闭包累加器实现
	acc := Accumulate(1)
	fmt.Println(acc())
	fmt.Println(acc())
	//打印累加器函数地址
	fmt.Printf("%p\n", acc)
	acc2 := Accumulate(10)
	fmt.Println(acc2())
	fmt.Printf("%p\n", acc2)
	fmt.Println(acc())

	//实例：闭包实现生成器
	generator := playerGen("high noon")
	name, hp := generator()
	fmt.Println(name, hp)

	//使用错误
	fmt.Println(div(1, 0))
	//使用自定义错误
	var err = New("我是错误")
	fmt.Println(err.Error())

	//让程序在崩溃时继续执行
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		//使用panic传递上下文 --> recover捕获
		panic(panicContext{"手动触发panic"})
		fmt.Println("手动宕机后")
	})
	//故意造成空指针
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int = new(int)  //初始化指针
		b := 1
		*a = b
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")



}

func test1() (int, int, int) {
	return 1, 2, 3
}
func test2() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return
}

//测试函数中参数传递
type Data struct {
	complax []int
	instance InnerData
	ptr *InnerData
}
type InnerData struct {
	a int
}
func passByValue(inFunc Data) Data {
	//输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)
	//打印inFunc的指针（地址相同且类型相同表示同一块内存区域）
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}

//字符串的链式处理设计
func StringProccess(list []string, chain []func(string) string) {
	//遍历每个字符串
	for index, str := range list {
		result := str
		for _, proc := range chain {
			//处理
			result = proc(result)
		}
		//将结果返回切片
		list[index] = result
	}
}
//自定义移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

//匿名函数作为回调函数
func visit(str string, fun func(string)) {
	fun(str)
}


//函数类型实现接口
//调用器接口
type Invoker interface {
	//传入一个interface{}类型变量，表示任意类型的值
	Call(interface{})
}
//结构体实现接口
type Struct struct {
}
//实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}
//函数体实现接口
type FuncCaller func(interface{})
//实现
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

//累加器
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++
		return value
	}
}

//闭包实现生成器
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

//使用错误
var err = errors.New("division by zero")
func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, err
	}
	return dividend / divisor, nil
}

//自定义错误
type errorString struct {
	s string
}
func New(text string) error {
	return &errorString{text}
}
func (e *errorString) Error() string {
	return e.s
}

//让程序崩溃时继续运行
type panicContext struct {  //崩溃时需传递的上下文信息
	function string
}
func ProtectRun(entry func()) {  //保护方式允许一个函数
	defer func() {
		//发生宕机时，获取panic传递的上下文信息
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}




