package main

import "fmt"

func swap1(a, b int) {
	a, b = b, a
	fmt.Printf("swap1: a =%d , b=%d\n ", a, b)
}

func swap2(p1, p2 *int) {
	*p1, *p2 = *p2, *p1

}

func main() {
	a, b := 10, 20
	fmt.Printf("main: a =%d , b=%d\n ", a, b)
	swap1(a, b)
	swap2(&a, &b)
	fmt.Printf("swap2: a =%d , b=%d\n ", a, b)
}
