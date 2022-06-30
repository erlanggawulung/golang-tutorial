package main
import ("fmt")

func main() {
	var a, b float64 = 0.1, 0.2
	var c float64
	c = a + b

	fmt.Printf("a + b = %.5f\n", c)
	fmt.Printf("0.3 == %.5f: %v\n", c, 0.3 == c)
	fmt.Printf("0.1 == %.5f: %v\n", a, 0.1 == a)
}