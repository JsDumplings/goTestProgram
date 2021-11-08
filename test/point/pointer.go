package point
// 指针指向 值的内存地址
func Content()(*int) {
	var ip *int
	a := 4
	ip = &a
	return ip
}