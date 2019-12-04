package main

import (
	"fmt"
)

func main() {
	power := 1000
	fmt.Printf("default power is %d\n", power)

	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}

//https://github.com/unknwon/go-fundamental-programming/blob/master/lectures/lecture1.md