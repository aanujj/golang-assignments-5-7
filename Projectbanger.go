package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	asking := os.Args[1]

	length := len(asking)

	final := strings.Repeat("*", length) + asking + strings.Repeat("*", length)

	final = strings.ToUpper(final)

	fmt.Println(final)
}

// hoya**** output
