package main

import "fmt"

const (
	one = iota*2 + 1 // укажите здесь формулу с iota
	three
	five
	seven
	nine
	eleven
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
