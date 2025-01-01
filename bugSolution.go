func myFunc(a []int) {b := make([]int, len(a)); copy(b, a); for i := range b {b[i]++}; fmt.Println(b)}
func main() {x := []int{1, 2, 3}; myFunc(x); fmt.Println(x)}