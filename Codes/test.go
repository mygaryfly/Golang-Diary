package main

import (
	"fmt"
<<<<<<< HEAD
	"strconv"
)

func main() {
	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)	  
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	} 
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
=======
)

func main() {
	power := 1000
	fmt.Printf("default power is %d\n", power)

	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}

//https://github.com/unknwon/go-fundamental-programming/blob/master/lectures/lecture1.md
>>>>>>> f7594bbeabc11cbfb341467a2c5ec2a6b0467b62
