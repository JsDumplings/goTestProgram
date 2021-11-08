package main

import "fmt"

var(
    e int
    a string
)

func testType() {
    var a string
    var b int
    var c bool
    var d []int
    var e *int
    var f map[string] int
    var g chan int
    var h func(string) int
    var m error // error 是接口
    fmt.Println(a,b,c,d,e,f,g,h,m)
}

func autotype() {
    // 只能在函数体中使用:=
    a = "郭英武"
    a = "电脑"
    b := 1
    c := [2]int{1,2}
    // var f int
    _,f := 1,2 //_ 值1被抛弃
    fmt.Println(a,b,c,f)
}

func ret()(int,int,string) {
 a,b,c := 1,2,"str"
 return a,b,c 
}

func getVal(){
    _,b,c := ret()
    fmt.Println(b,c)
}

/* 
常量部分
只能使用内置函数
iota //特殊常量，可以被编译器修改的常量
iota 在 const关键字出现时将被重置为 0
const 中每新增一行常量声明将使 iota 计数一次
iota 可理解为 const 语句块中的行索引
在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。
**/

const (
    a1 = 1
    b2 = "String"
    c3 = len(b2)
    a2 = iota 
    a3 = iota
    a4 = "值"
    a5 = iota
)

func main() {
    getVal()
    var g1 int = 4
    var prt *int 
    prt = &g1
    fmt.Println("prt",prt,"*prt",*prt)
    g1d := &g1 // & 变量地址
    fmt.Println(a1,b2,c3,a2,a3,a4,a5,g1d)
}

func main()  {
    fmt.Println("第二个main函数")
}