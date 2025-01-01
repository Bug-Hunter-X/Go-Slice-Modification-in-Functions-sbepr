func myFunc(a []int) {for i := range a {a[i]++}}
func main() {x := []int{1, 2, 3}; myFunc(x); fmt.Println(x)}