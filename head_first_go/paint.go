package main
import "fmt"
func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%6.2f liters needed\n", area/10.0)
}
func main() {
	paintNeeded(100,24.3)
	paintNeeded(5.5, 33.3)
	paintNeeded(5.0, 10.0)
}
