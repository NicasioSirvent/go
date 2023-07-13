package main
import (
	"bufio"
	"fmt"
	"os"
	"log"
)
func main() {
	fmt.Print("Enter a Grade: ")
	reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	if  err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
