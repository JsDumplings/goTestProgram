package godemo

import (
	"fmt"
	"time"
)

func Fibonacci(n int, c chan int) {
	x, y := 0, 1
	
	for i :=0; i < n; i++ {
			c <- x
			fmt.Println("in:",time.Now())
			// time.Sleep(100)
			x, y = y, x+y
	}
	
	close(c)
}