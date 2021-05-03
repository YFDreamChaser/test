package main

import (
	"container/list"
	"fmt"
	"sort"
	"sync"
)

func main() {
	//1.数组定义
	var team = [...]string{"a", "b", "c"}
	for k, v := range team {
		fmt.Println(k, v)
	}


	//2.切片  --> 对切片追加若没突破其底层数组的长度，则返回的切片仍然指向原来的数组。如果突破了，则重新申请数组内存，指向该数组
	//从数组或切片生成新的切片
	fmt.Println(team, team[1:2]) //含起始位置，不包含结束位置
	fmt.Println(team[:])
	team2 := team[:]
	team3 := team[1:3]
	team2[0] = "b"
	team3[0] = "e"
	fmt.Println(team[0:0], team, team2, team3)
	//切片声明
	//var strList []string  --> 未发生内存分配操作
	//make函数构造(类型, size, cap)  cap是提前分配多少空间 -> 一定发送了内存分配操作
	a1 := make([]int, 2, 10)
	a1[0] = 1
	a1[1] = 2
	a1 = append(a1, 3, 4, 5)
	fmt.Println(a1, len(a1), cap(a1))
	//切片扩容
	var numbers []int
	for i := 0; i < 3; i++ {
		numbers = append(numbers, i)  //添加一个元素  成倍扩充
		fmt.Printf("len: %d, cap: %d, pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	//一个切片添加另外一个切片的所有元素 --> 扩容，会重新分配内存
	numbers = append(numbers, a1...)
	numbers[5] = 100
	fmt.Println(numbers, a1)
	//复制切片元素到另一个切片
	const elemCount = 1000
	srcData := make([]int, elemCount)
	for i := 0; i < elemCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copyData := make([]int, elemCount)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[elemCount - 1])
	copyData[999] = 1
	copy(copyData, srcData[:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}
	fmt.Println(copyData[999])
	//从切片中删除  --> 内存是一样的
	seq := []string{"a", "b", "c", "d", "e"}
	index := 2
	a := seq[:index]
	b := seq[index+1:]
	seq = append(a, b...)
	a[0] = "f"
	b[0] = "h"
	fmt.Println(seq)


	//3 映射（map）
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	//判断key是否在map中再加个变量
	v, ok := scene["route2"]
	fmt.Println(v, ok)
	//声明时填充
	m := map[string]string {
		"w": "forward",
		"a": "left",
		"d": "right",
	}
	for k, v := range m {  //注意是无序的
		fmt.Println(k, v)
	}
	//如果需要特定顺序的遍历结果
	var sortm []string
	for k := range m {
		sortm = append(sortm, k)
	}
	//对切片进行排序
	sort.Strings(sortm)
	fmt.Println(sortm)
	//删除键值对
	delete(m, "a")
	fmt.Println(m)
	//并发的map: sync.map
	/*//使用map
	m2 := make(map[int]int)
	go func() {
		for {
			m2[1] = 1  //不断写入
		}
	}()
	go func() {
		for  {
			_ = m2[1]  //不断读取
		}
	}()
	for {  //死循环，让并发程序在后台不断执行

	}*/
	//使用sync.map
	var m3 sync.Map
	m3.Store("greece", "97") //存储
	m3.Store("london", "100")
	m3.Store("egypt", "200")
	fmt.Println(m3.Load("london")) //取值 100 ture
	m3.Delete("london") //删除
	m3.Range(func(k, v interface{}) bool {  //range调用匿名函数将结果返回
		fmt.Println("iterate", k, v)
		return true
	})



	//4 列表list,可快速增删的非连续空间容器
	//插入相关api
	l1 := list.New()
	l2 := list.New()
	s1 := l1.PushBack("second")
	l1.PushBack("third")
	l1.PushFront("first")
	l1.InsertAfter("second2", s1)
	l2.PushBack("lll")
	l2.PushFrontList(l1)
	/*for i := l2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}*/
	//删除
	rem := l1.Remove(s1)  //这个只能删除l1中的s1
	fmt.Println(rem)
	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

