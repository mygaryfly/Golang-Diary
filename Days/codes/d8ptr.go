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

/* 质数
package main

import (
"fmt"
)

var x int
var f bool

func main(){

	fmt.Scanf("%d\n",&x)
	fmt.Printf("%v\n",x)
	fmt.Printf("%v\n",f)
	judge(x,&f)
	fmt.Printf("%t\n",f)
}

func judge(a int,b *bool){
	b=&f
	a=x
	for i:=2;i<a;i++{
		if a%i==0 {
		*b = false
		break
		}else if a%i>0 {
			*b = true
		}
	}
}
*/
