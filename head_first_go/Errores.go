package main
import (
	"errors"
	"fmt"
)
func main() {
	err := errors.New("Height cant be negative!!")
	fmt.Println(err.Error())
}
