
#### **数组练习**
练习 1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。
> 答案：
```
package main
import "fmt"

func main() {
	var a,b,c,d = 11,22,33,44
	fmt.Println("a address = ",&a)
	fmt.Println("b address = ",&b)
	fmt.Println("c address = ",&c)
	fmt.Println("d address = ",&d)
	
	var arr = [4]int{a,b,c,d}
	//var ptr [4]*int
	
	for i,x:=range arr {
		fmt.Printf("数组arr[%d]赋值打印：%p\n",i+1,&x)
	}
}
```
*输出结果：*
```
a address =  0x11810098
b address =  0x1181009c
c address =  0x118100b0
d address =  0x118100b4
数组arr[1]赋值打印：0x118100b8
数组arr[2]赋值打印：0x118100b8
数组arr[3]赋值打印：0x118100b8
数组arr[4]赋值打印：0x118100b8
```
练习 2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。

练习 3：fibonacci_array.go: 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。