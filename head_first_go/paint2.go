package main
import "fmt"
func paintNeeded(width float64, height float64){
	area := height * width;
	fmt.Printf("%6.2f liters needed\n", area/10.0)
}
func main() {
	paintNeeded(45, 10)
	paintNeeded(10.0, 5.6)
}

