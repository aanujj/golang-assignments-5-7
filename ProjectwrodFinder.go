package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = " " + "lazy cat jumps again and again and again"

func main() {

	words := strings.Fields(corpus)
	query := os.Args[1:]
queries:
	for _, q := range query {

		for i, w := range words {

			switch q {
			case "again", "and":
				break queries
			}

			if q == w {
				fmt.Printf("%v is : %v\n", i+1, q)

				break queries
			}

		}
	}

}
