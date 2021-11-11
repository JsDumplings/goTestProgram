// 斐波那函数
package fib
import ("fmt")
var res int
func Fac(x int)(int) {
	if x > 1 {
		res += x * (x - 1)
		x--
		fmt.Println("res",res)
		Fac(x)
	}
	return res
}
func Fib(x int)(int)  {
	if x > 2 {
		return Fib(x-2) + Fib(x - 1)
	}else{
		return x 
	}
}