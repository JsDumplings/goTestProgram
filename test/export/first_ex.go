package export
var Total_sum int = 0
func Add(x,y int)(int) {
	Total_sum = (x + y + 1)
	return x + y
}
