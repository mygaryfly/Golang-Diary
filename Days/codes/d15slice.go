//【练习2.1 - 1：】fibonacci_funcarray.go: 练习主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片
package main

import "fmt"
var fibonacci = [7]int{1,1,2,3,5,8,13}

func main(){

	fmt.Printf("fx(4) =%d\n ",fx(4))

}
	func fx(index int)(slice []int){

			slice = fibonacci[0:index]
		
		return slice
	}