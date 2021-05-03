package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//结构体
	//匿名结构体
	msg := &struct {  //定义部分
		id int
		data string
	}{  //初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)

	//多种方式创建和初始化结构体

	//Go语言结构体方法
	b := new(Bag)  //返回指针类型
	b.Insert(100)
	fmt.Println(b.items)

	//指针和非指针接收器的使用:
	//在计算机中，小对象由于值复制时的速度比较快，所以适合使用非指针接收器。
	//大对象复制性能低，适合指针接收器。在接收器和参数间传递时不进行复制，只是传递指针
	//指针接收器：可写数据。非指针接收器：只读

	//为基本类型添加方法
	var m MyInt
	fmt.Println(m.IsZero())
	m = 1
	fmt.Println(m.Add(2))

	//http请求
	/*client := &http.Client{}
	//创建一个http请求
	req, err := http.NewRequest("POST", "https://www.baidu.com/",
		strings.NewReader("key=value"))
	//发现错误就打印退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//为标头添加信息
	req.Header.Add("User-Agent", "myClient")
	//开始请求
	resp, err := client.Do(req)
	//处理请求的错误
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//读取服务器返回内容
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()*/

	//方法和函数的统一调用
	var delegate func(int)  //声明一个函数回调
	c := new(class)  //创建结构体实例
	delegate = c.Do  //将回调设为Do方法
	delegate(100) //调用
	delegate = funcDo  //将回调设为普通函数
	delegate(100)

	//实例：事件系统基本原理 (不关注事件注册的顺序调用)
	act := new(Actor)
	RegisterEvent("OnSkill", act.OnEvent) //粉丝关注偶像并注册事件
	RegisterEvent("OnSkill", GlobalEvent) //粉丝关注偶像并注册事件
	CallEvent("OnSkill", 100) //明星出绯闻回调函数粉丝就能知道事件

	//结构体内嵌
	var col Color
	col.R = 1  //内嵌结构体简化写法
	col.G = 1
	col.B = 0
	col.alpha = 1
	fmt.Printf("%+v\n", col)

	//结构体内嵌实现对象组合特性
	bird := new(Bird)
	bird.Fly()
	bird.Walk()
	human := new(Human)
	human.Walk()

	//使用匿名结构体分离JSON数据
	jsonData := getJsonData()
	fmt.Println(string(jsonData))
	//只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool  //名称必须和json中的名称一致
	}{}
	fmt.Println(screenAndTouch.HasTouchID)
	//反序列化到screenAndTouch中
	json.Unmarshal(jsonData, &screenAndTouch)  //数据和格式
	//输出screenAndTouch详细结构
	fmt.Printf("%+v\n", screenAndTouch)




}

func printMsgType(msg *struct{
	id int
	data string
})  {
	fmt.Printf("%T\n", msg)
}

type Cat struct {
	Color string
	Name string
}
//模拟构造函数重载
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}
//模拟父级构造调用
type BlackCat struct {
	Cat  //嵌入Cat，类似于派生(BlackCat 拥有 Cat 的所有成员实例化后可以自由访问 Cat 的所有成员)
}
//构造基类
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
//构造子类
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

//Go语言结构体方法
type Bag struct {
	items []int
}
func (b *Bag) Insert(item int)  {
	b.items = append(b.items, item)
}

//为基本类型添加方法
type MyInt int
func (m MyInt) IsZero() bool {
	return m == 0
}
func (m MyInt) Add(other int) int {
	return other + int(m)
}

//方法和函数的统一调用
type class struct {

}
func (c *class) Do(v int)  {  //结构体方法
	fmt.Println("call method do:", v)
}
func funcDo(v int)  {  //普通函数方法
	fmt.Println("call function do:", v)
}

//实例：事件系统基本原理
var eventByName = make(map[string][]func(interface{}))
//注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	//通过名字查找事件列表
	list := eventByName[name]
	//在列表切片中添加函数
	list = append(list, callback)
	//保存修改的事件列表切片
	eventByName[name] = list
}
//调用事件
func CallEvent(name string, param interface{}) {
	list := eventByName[name]
	for _, callback := range list {
		callback(param)
	}
}
//使用事件
type Actor struct {
}
//为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}
//全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

//结构体内嵌
type BasicColor struct {
	R, G, B float32
}
type Color struct {
	BasicColor
	alpha float32
}

//内嵌实现组合特性
//飞行
type Flying struct {

}
func (f Flying) Fly()  {
	fmt.Println("can fly")
}
//行走
type Walkable struct {

}
func (w *Walkable) Walk()  {
	fmt.Println("can walk")
}
type Human struct {
	Walkable  //人类能行走，把行走当成自己的属性内嵌
}
type Bird struct { //鸟类能行走并且能飞行
	Walkable
	Flying
}

//使用匿名结构体分离JSON数据
type Screen struct {  //屏幕
	Size float32   //屏幕尺寸
	ResX, ResY int  //屏幕水平和垂直分辨率
}
type Battery struct {  //电池
	Capacity int   //容量
}
//准备JSON数据
func getJsonData() []byte {
	//完整数据结构
	raw := &struct {  //匿名结构体
		Screen
		Battery
		HasTouchID bool
	}{  //初始化
		Screen{
			5.5,
			1920,
			1080,
		},
		Battery{
			2910,
		},
		true,
	}
	//Jiang数据序列化为JSON
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

