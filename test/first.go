package main
import (
	"fmt"
	"time"
	// "./export"
	// "./arr"
	// "./point"
	// "./recursion"
	"./godemo"
)
func main() {
	// sum:= export.Add(1,2)
	// fmt.Println(sum,export.Total_sum)

	// 数组
	//数组是类型相同的元素的集合。例如，整数 5, 8, 9, 79, 76 的集合就构成了一个数组
	// result := arr.Arr
	// var arrs = [3]int{9,8,7}
	// fmt.Println(result,arrs)

	// 指针
	// var a int
	// var ip *int
	// a = 4
	// ip = &a
	// fmt.Println(a,*ip)

	// 指针数组
	// var arr1 = [...]int{10,20,30}
	// const MAX int = 3
	// var prt [MAX]*int
	// for i:=0;i < MAX;i++ {
	// 	prt[i] = &arr1[i]
	// 	fmt.Printf("arr1[%d]= %d\n",i,arr1[i])
	// }
	// for i:=0;i<MAX;i++ {
	// 	fmt.Println(i,prt[i])
	// }

	// 指向指针的指针
	// var prt0 *int
	// var prt **int
	// a:= 1
	// prt0 = &a
	// prt = &prt0
	// fmt.Println(a,*prt0,**prt)

	// 结构体 类似于js的对象
	// 属性的首字母大小写问题
	// 首字母大写相当于 public。
  // 首字母小写相当于 private。
	//注意: 这个 public 和 private 是相对于包（go 文件首行的 package 后面跟的包名）来说的。
  //敲黑板，划重点
  //当要将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON。
	// type Books struct{
	// 	name string
	// 	number int
	// }
	// // var typeArr = [2]string{"科技","教育"}
	// var book1 Books
	// book1.name = "书名"
	// book1.number = 2
	// // book1.bookType = typeArr
	// fmt.Println(book1)

	// 切片：是对数组的抽象、也被称为动态数组
	// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	// 切片（slice）是建立在数组之上的更方便，更灵活，更强大的数据结构。切片并不存储任何元素而只是对现有数组的引用。
	// 数组是不变的，数组是属于值传递,切片便是对于数组的处理。
	//new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址, ps:适用于值类型 数组与结构体
	//make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型:切片、map 和 channel。
	// var s = [...] int{1,2,3}
	// var sl []int
	// sl = s[:2]
	// fmt.Println("切片sl（整个数组）",sl)
	// //拷贝方法 copy（）
	// var sl1  = make([]int, len(sl), cap(sl)*2)
	// copy(sl1, sl)
	// fmt.Println("拷贝后的sl1",sl1,"拷贝后的sl",sl)
	// // 切片增加
	// // 当原本容量不足时，会自动扩增，以倍增的方式
	// // 通过 append() 函数向数组中添加元素，首先 cap 会以二倍的速度增长，
	// //如果发现增长 2 倍以上的容量可以满足扩容后的需求，那么 cap*2，否则就会看扩容后数组的 length 是多少 cap=length+1。
	// sl1 = append(sl1, sl[0],sl[1],3,4)
	// fmt.Println("增加sl后的sl1",sl1)

	// range 
	// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
	// a := []int{1,2,3,4}
	// for i,val := range a {
	// 	fmt.Println("i",i)
	// 	fmt.Println("val",val)
	// }

	// map
	// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	// Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
	// var mapD = map[string]string{"name":"郭1","age":"16","play":"game"}
	// mapD := make(map[string]string)
	// mapD["name"] = "郭1"
	// mapD["age"] = "14"
	// for key,val := range mapD {
	// 	fmt.Println("val",val)
	// 	fmt.Println("key",key)
	// }

	// 递归
	// a := 10
	// res := fib.Fib(a)
	// f := fib.Fac(a)
	// fmt.Printf("%d的阶乘是: %d\n", a, f)
	// fmt.Printf("%d的斐波那契数列: %d", a, res)

	// 语言类型转换
	// 类型转换用于将一种数据类型的变量转换为另外一种类型的变量

	// var sum int = 17
  // var count int = 5
  // var mean float32
	// mean = float32(sum)/float32(count)
  // fmt.Printf("mean 的值为: %f\n",mean)

	//go 不支持隐式转换类型
	// var a int64
	// var b int32
	// a = 3
	// b = int32(a) // 必须要转换32
	// fmt.Println("b",b)


	//接口
	// 它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

	// 错误处理
	// 内置的错误接口提供了非常简单的错误处理机制。error类型是一个接口类型
	// panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误，
	// panic 用于主动抛出错误，recover 用来捕获 panic 抛出的错误。

	//go并发 通道缓存

	c := make(chan int,10) // 有缓冲，类似于为异步的
	go godemo.Fibonacci(cap(c),c)

	for v:= range c {
		fmt.Println("out:", time.Now())
    fmt.Println(v)
	}
}