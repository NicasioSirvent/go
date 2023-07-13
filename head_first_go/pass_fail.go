package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	//status := ""
	var status string
	fmt.Print("Enter a Grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	grade, error := strconv.ParseFloat(input, 64)
	if error != nil {
		log.Fatal(error)
	}

	if grade >= 60.0 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
