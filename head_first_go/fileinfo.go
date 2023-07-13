package main
import (
	"fmt"
	"log"
	"os"
)
func main() {
	//fileInfo, _ := os.Stat("apples.go")
	fileInfo, error := os.Stat("apples.go")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(fileInfo.Size())
}
