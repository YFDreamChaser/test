package main

import (
	"fmt"
	"logger"
	"sort"
)

func main() {
	//数据写入器的抽像
	f := new(file)
	var writer DataWriter = f  //接口复制(类似父类引用指向字类对象)
	writer.WriteData("data")

	//接口中所有方法都被实现才能使用


	//一个类型可以实现多个接口
	socket := new(Socket)
	usingWriter(socket)
	usingCloser(socket)

	//多个类型可以实现一个接口
	var s Service = new(GameService)
	s.Start()
	s.Log("")


	//实例：便于扩展输出方式的日志系统
	l := createLogger()
	l.Log("hello")


	//实例：使用接口进行数据排序
	/*names := MyStringList{  //只是扩充原类，之前有的功能也能写
		"3",
		"5",
		"2",
		"4",
		"1",
	}
	sort.Sort(names)*/
	//其实有了这个，直接用
	/*names := sort.StringSlice{
		"3",
		"5",
		"2",
		"4",
		"1",
	}
	sort.Sort(names)*/
	//更加简化
	/*names := []string{
		"3",
		"5",
		"2",
		"4",
		"1",
	}
	sort.Strings(names)*/
	//排序int
	names := []int{
		2,
		1,
		3,
		5,
		4,
	}
	sort.Ints(names)
	for _, v := range names {
		fmt.Println(v)
	}

	//结构体排序实例1
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Ass},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Ass},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
	//更加简便的结构体排序实例2
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind > heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}



	//将接口转换为其他接口
	animals := map[string]interface{} {
		"bird": new(bird),
		"pig": new(pig),
	}
	for name, obj := range animals {
		//类型断言，f(转变后的变量),isFlyer(obj接口是否实现f类型),obj(接口变量)
		//obj是否可转化为Flyer
		f, isFlyer := obj.(Flyer)
		//判断对象是否是行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}


	//空接口断言
	var a int = 1
	var i interface{} = a
	//编译报错
	//var b int = i
	var b int = i.(int)
	fmt.Println(b)

	//使用类型分支判断基本类型
	printType(1024)
}

//数据写入器的抽像
type DataWriter interface {
	WriteData(data interface{}) error
	//新增一个方法 --> 则报错
	//CanWrite() bool
}
type file struct { //定义文件结构实现DataWriter
}
func (d *file) WriteData(data interface{}) error { //实现方法
	//模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

//一个类型可以实现多个接口
type Socket struct {
}
func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("use writer")
	return 0, nil
}
func (s *Socket) Close() error {
	fmt.Println("use closer")
	return nil
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
func usingWriter(writer Writer) { //把Socket赋值给Writer,与Closer无关
	writer.Write(nil)
}
func usingCloser(closer Closer) {
	closer.Close()
}

//多个类型可以实现一个接口
type Service interface {
	Start()  //开启服务
	Log(string)  //日志输出
}
type Logger struct {
}
func (l *Logger) Log(string)  {
	fmt.Println("输出日志")
}
type GameService struct {
	Logger  //嵌入日志器
}
func (g *GameService) Start() {
	fmt.Println("开启服务")
}

//创建日志器
func createLogger() *logger.Logger {
	//创建日志器
	l := logger.NewLogger()
	//创建命令行写入器
	cw := logger.NewConsoleWriter()
	//注册命令行写入器到日志器中
	l.RegisterWriter(cw)
	//创建文件写入器
	fw := logger.NewFileWriter()
	//设置文件名
	if err := fw.SetFile("log.txt"); err != nil {
		fmt.Println(err)
	}
	//注册文件写入器到日志器
	l.RegisterWriter(fw)
	return l
}

//实例：使用接口进行数据排序
type MyStringList []string
//实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}
//实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}
//实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

//结构体排序实例1
//声明英雄分类
type HeroKind int
//定义常量
const (
	No HeroKind = iota
	Tank
	Ass
	Mage
)
//定义英雄名单结构
type Hero struct {
	Name string   //英雄名字
	Kind HeroKind  //英雄种类
}
//将英雄指针的切片定义为Heros类型(指针类型，因为要修改结构体内容)
type Heros []*Hero
//长度接口
func (s Heros) Len() int {
	return len(s)
}
//比较接口
func (s Heros) Less(i, j int) bool {
	//如果英雄的分类不一致，优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	return s[i].Name < s[j].Name
}
//交换元素  --> 交换指针可能比交换数据快
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//将接口转换为其他接口
type Flyer interface {
	Fly()
}
type Walker interface {
	Walk()
}
type bird struct {
}
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}
type pig struct {
}
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

//使用类型分支判断基本类型
func printType(v interface{})  {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}


