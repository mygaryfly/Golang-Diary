
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
> 答案-1
```
package main

import "fmt"

func main(){

	var array [16]int
	
	for i:=0;i<16;i++{
		
		array[i]=i
		
	}
	fmt.Println(array)
}
```

> 答案-2
```
package main

import "fmt"

func main(){

	var arr [16]int

	for  x:=range arr{
	
		 arr[x]=x
			
	}
	fmt.Println(arr)
}
```
*输出结果：* // 两种方式的输出结果是一样的
`[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]`
【踩过的坑】
* 1，`var arr [16]int`不能写在for循环里面。
* 2，Print语句要放在for循环外。

练习 3：fibonacci_array.go: 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。
> 答案 -1
```
package main //方法一，数组for循环后打印

import "fmt"

var fibo [50]int

func main(){
	
	for n:=0;n<50;n++{
		
		if n<=1 {
			fibo[n]=1
		} else {
			fibo[n]=fibo[n-1]+fibo[n-2]
		}
	}
	fmt.Println(fibo)

}
```
> 答案 -2
```
package main //方法二，递归函数

import "fmt"

var x int

var Fibo [50]int

func main (){
	
	for i:=0;i<len(Fibo);i++{
	
		Fibo[i]=fx(i)
		
		fmt.Printf("Fibo[%d] = %d\n",i,fx(i))
		}
}

func fx(x int)(y int){
	
	if x<=1 {
			y=1
	} else {
		y = fx(x-1)+fx(x-2)
	} 
	return
} 
```
【踩过的坑】
* 当函数命名和数组命名相同时，系统报错类似`cannot call non-function + 命名`，该现象是因函数和数组命名相同，导致其中一个的功能被屏蔽。main函数是最后运行的函数，内部调用Fibo数组前已经运行过Fibo函数，此时可以视为Fibo函数屏蔽了同名称的Fibo数组的功能。
