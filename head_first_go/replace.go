package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	//replacer := strings.NewReplacer("#", "o")
	//fixed := replacer.Replace(broken)
	fixed := strings.NewReplacer("#", "o").Replace(broken)
	fmt.Println(fixed)
}
