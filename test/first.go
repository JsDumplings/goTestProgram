package main
import (
	"fmt"
	// "./export"
	// "./arr"
	// "./point"
)
func main() {
	// sum:= export.Add(1,2)
	// fmt.Println(sum,export.Total_sum)

	// 数组
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
	var prt0 *int
	var prt **int
	a:= 1
	prt0 = &a
	prt = &prt0
	fmt.Println(a,*prt0,**prt)
}