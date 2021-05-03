package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	//1
	//标准声明并初始化
	var t bool = true
	fmt.Println(t)

	//短变量声明并初始化(只能用于未被声明的变量)
	hp := 100
	/*//net.Dial 提供按指定协议和地址发起网络连接，这个函数有两个返回值，一个是连接对象， 一个是err对象
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//等价于标准格式
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:8080")*/
	/*//注：在多个短变量声明中，至少有一个新声明变量出现，就不会报错。如：
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	conn2, err := net.Dial("tcp", "127.0.0.1:8080")*/
	fmt.Println(hp)

	//多个变量同时赋值
	//交换
	var a int = 100
	var b int = 200
	a, b = b, a
	fmt.Println(a, b)
	//交换排序  --> 定义一个切片类型
	var intSlice IntSlice = make([]int, 4)
	intSlice[0] = 1
	intSlice[1] = 2
	intSlice[2] = 3
	intSlice[3] = 4
	fmt.Println(intSlice)
	for i := 0; i < intSlice.Len() - 1; i++ {
		for j := i + 1; j < intSlice.Len(); j++ {
			if intSlice.Less(i, j) {
				intSlice.Swap(i, j)
			}
		}
	}
	fmt.Printf("%v\n", intSlice)

	//匿名变量（下划线） -> 不占用命名空间并且不分配内存
	d1, _ := GetData()
	_, d2 := GetData()
	fmt.Println(d1, d2)



	//2数据类型
	//哪些情况下使用int和uint？
	/*逻辑对整型范围没有特殊需求。
	例如，对象的长度使用内建 len（）函数返回，这个长度 可以根据不同平台的字节长度进行变化。实际使用中 ， 切片或 map 的元素数量等都可以用 int 来表示。
	反之，在二进制传输、 读写文件的结构描述时， 为了保持文件的结构不会受到不同编 译目标平台字节长度的影响，不要使用 int 和 uint。
	 */

	/*//浮点型：输出正弦函数图像
	//图片大小
	const size = 300
	//根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充为白色
			pic.SetGray(x, y, color.Gray{Y:255})
		}
	}
	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		//让sin的值的范围在0-2Pi之间
		s := float64(x) * 2 * math.Pi / size
		//sin的幅度为一半的像素，向下偏移一半像素并翻转
		y := size / 2 - math.Sin(s) * size / 2
		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{Y:0})
	}
	//创建文件
	file, err := os.Create("C:/Users/Administrator/Desktop/sin.png")
	if err != nil {
		//如果创建文件失败，返回错误，打印错误并终止。
		log.Fatal(err)
	}
	//使用png格式将数据写入文件
	png.Encode(file, pic)
	//关闭文件
	file.Close()*/

	//字符串
	fmt.Println("反斜杠为转义字符 -> str := \"c:\\Go\\bin\"")
	const str = `第一行
					第二行
					第三行`
	fmt.Println(str)

	//字符
	//一种是 uint8 类型， 或者叫 byte 型， 代表了 ASCII 码的一个字符。
	//另一种是 rune 类型，代表一个 UTF-8 字符(如中文，实际是int32)
	//使用 fmt.Printf 中的“%T”动词可以输出变量的实际类型
	var c1 byte = 'a'
	var c2 rune = '你'
	fmt.Printf("%d %T\n", c1, c1)
	fmt.Printf("%d %T\n", c2, c2) //unicode码20320(unicode是字符集。utf-8是编码规则)

	//切片--能动态分配的空间(arraylist?)
	q := make([]int, 3)
	q[0] = 1
	q[1] = 2
	q[2] = 3
	for index, value := range q {
		fmt.Printf("q[%d] = %d ", index, value)
	}
	//切片还可以在其元素集合内连续地选取一段区域作为新的切片
	q2 := "hello word"
	fmt.Println(q2[6:8])


	//3转换不同数据类型
	//指针
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
	/*//使用指针变量获取命令行输入信息(go run flagparse.go mode=fast )
	//定义命令行参数(参数名称，参数值的默认值，参数说明)
	var mode *string = flag.String("mode", "", "process mode")
	//解析命令行参数
	flag.Parse()
	//输出
	fmt.Println(*mode)*/
	//创建指针另一种方法：new()函数 [*在左边，表示指针指向的变量，右边，表示指针变量取值]
	str2 := new(string)
	*str2 = "ninja"
	fmt.Println(*str2)



	//4 字符串应用
	//len 通用，获取长度
	tip1 := "YangFan love WangLuanNi"
	tip2 := "忍者"  //utf-8格式保存，每个中文3个字节
	tip3 := utf8.RuneCountInString("忍者,fight")  //该函数统计Uncode个数
	fmt.Println(len(tip1), len(tip2), tip3)
	//遍历每一个ASCII字符 --> 汉字惨不忍睹
	theme := "狙击 start"
	/*for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}*/
	//按unicode字符遍历字符串
	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
	//获取字符串的某一段字符 --> 搜索的起始位置可以用切片方式进行
	tracer := "死神来了，死神bye bye"
	comma := strings.Index(tracer, "，") //ASCII码位置，12
	pos := strings.Index(tracer[comma:], "死神") // 3 --> 二次检索
	fmt.Println(comma, pos, tracer[comma+pos:]) //死神bye bye
	//修改字符串(string不可变,实际修改[]byte,然后通过切片重构string)
	angel := "Heros never die"
	angelBytes := []byte(angel) //转换成字符串数组
	for i := 5; i <= 10; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(angel)
	fmt.Println(angelBytes)
	fmt.Println(string(angelBytes))
	//字符串连接
	//可以使用+
	hammer := "yangfan" + "wln"
	//可以使用类似java sb的方式
	//声明字节缓冲
	var sb bytes.Buffer
	//把字符串写入缓冲
	sb.WriteString("yangfan")
	sb.WriteString("fan")
	fmt.Println(hammer, sb.String())
	//枚举 定义一个int类型的枚举类型
	type Weapon int
	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	//输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Println(FlagNone, FlagRed, FlagGreen, FlagBlue)
	//将枚举值转化为字符串
	fmt.Printf("%s %d\n", CPU, CPU)
	//在结构体成员嵌入时使用别名
	var v Vehicle   //声明车辆类型
	v.FakeBrand.Show()   //指定条用别名的show
	tv := reflect.TypeOf(v)   //取v类型的反射对象
	for i := 0; i < tv.NumField(); i++ {   //遍历v所有成员
		f := tv.Field(i)
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type)
	}
}

type IntSlice []int
type ChipType int
func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetData() (int, int) {
	return 100, 200
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
const (
	None ChipType = iota
	CPU
	GPU
)
//相当于重写String函数？
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
//定义商标结构
type Brand struct {

}
//为商标结构添加Show方法
func (t Brand) Show() {

}
//为Brand起别名
type FakeBrand = Brand
//结构体中嵌入
type Vehicle struct {
	//嵌入两个结构
	FakeBrand
	Brand
}





