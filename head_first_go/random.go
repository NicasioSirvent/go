package main
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10 - guesses, "guesses left")
		fmt.Print("Make a Guess: ")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		input = strings.TrimSpace(input)
		guess, error := strconv.Atoi(input)
		if error != nil {
			log.Fatal(error)
		}

		if guess < target {
			fmt.Println("Opps, Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Opps, Your guess was HIGH.")
		} else {
			fmt.Println("Good Job! You guessed it!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("You Failed, the correct number was", target)
	}
}
