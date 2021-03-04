package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1]

	guess, _ := strconv.Atoi(args)

	for turn := 1; turn <= 5; turn++ {

		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("you won 👻")
			fmt.Println()
			return
		}

	}
	fmt.Println("you lost  💩")

}
