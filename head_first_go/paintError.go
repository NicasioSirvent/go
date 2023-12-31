package main
import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid!", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid!", height)
	}
	return width * height / 10.0, nil
}

func main() {
	amount, err := paintNeeded(-5.3, +3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
}
