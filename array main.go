package main

import "fmt"

const (
	winter = 1
	yearly = 4
)

func main() {

	var books [yearly]string
	books[0] = "kafka's revenge "

	books[1] = " Marvel"

	books[2] = "Stay Golden"

	books[3] = books[0] + "second edition"

	fmt.Printf("books :  %#v\n", books)

	var (
		wbooks [winter]string
		sbooks [3]string
	)

	wbooks[0] = books[0]

	for i := range sbooks {
		sbooks[i] = books[i+1]
	}

	fmt.Println("wbooks : ", wbooks)
	fmt.Printf("summer books : %#v", sbooks)

}
