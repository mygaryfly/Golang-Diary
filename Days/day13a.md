### **函数 - 补充**

#### **一， 不定参函数-参数的传递**

##### 1.1 **如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。**

##### *_示例 - 1:_*
```
package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
```
*输出结果：*
```
The minimum is: 0
The minimum in the slice is: 1
```

##### 1.2 **一个接受变长参数的函数可以将这个参数作为其它函数的参数进行传递：**

##### *_示例 - 2:_*
```
func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) { }
func F3(s []string) { }
```
* 变长参数可以作为对应类型的 slice 进行二次传递。

> 但是如果变长参数的类型并不是都相同的呢？使用 5 个参数来进行传递并不是很明智的选择，有 2 种方案可以解决这个问题：

* 1, 使用结构
> 定义一个结构类型，假设它叫 Options，用以存储所有可能的参数：
```
type Options struct {
	par1 type1,
	par2 type2,
	...
}
```
函数 F1 可以使用正常的参数 a 和 b，以及一个没有任何初始化的 Options 结构： `F1(a, b, Options {}`)。如果需要对选项进行初始化，则可以使用 `F1(a, b, Options {par1:val1, par2:val2})`。

* 2, 使用空接口
> 如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数（详见第 11.9 节）。该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。一般而言我们会使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：
```
func typecheck(..,..,values … interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: …
			case float: …
			case string: …
			case bool: …
			default: …
		}
	}
}
```

#### **二， defer和追踪**

#### **2.1 defer**
*介绍：* 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
* 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 `finally`语句块，它一般用于释放某些已分配的资源。

##### *_示例 - 3:_*
```
package main
import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2() // defer将function2函数推迟执行
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}
```
*输出结果A：*
```
In Function1 at the top
In Function1 at the bottom!
Function2: Deferred until the end of the calling function!
```
> 我们将 defer 关键字去掉并对比输出结果。

*输出结果B：*
```
In function1 at the top
Function2: Deferred until the end of the calling function!In function1 at the bottom!
```
* *可见，若拿走defer关键字，函数function2按照原来的顺序被执行*

#### **2.1.1** - 使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0：
##### *_示例 - 4:_*
```
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
```
#### **2.1.2** - 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
##### *_示例 - 5:_*

```
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
```
*输出结果：* `4 3 2 1 0`

#### **2.1.3** - 关键字 defer 允许我们进行一些函数执行完成后的收尾工作

> 1, 关闭文件流
```
// open a file  
defer file.Close()
```
> 2, 解锁一个加锁的资源
```
mu.Lock()  
defer mu.Unlock() 
```
> 3, 打印最终报告
```
printHeader()  
defer printFooter()
```
> 4,关闭数据库链接
```
// open a database connection  
defer disconnectFromDB()
```

* 合理使用 defer 语句能够使得代码更加简洁。

##### *_示例 - 6:_*
> 以下代码模拟了上面描述的第 4 种情况：
```
package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}
```
*输出结果：*
```
ok, connected to db
Defering the database disconnect.
Doing some DB operations ...
Oops! some crash or network error ...
Returning from function here!
ok, disconnected from db
```
#### **2.2 使用 defer 语句实现代码追踪**

```
package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a() 
} 
// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）

func main() {
	b()
}
```
*输出结果：*
> 两次跟defer有关的执行打印，最终都被放到最后并按照后进先出的顺序执行。
```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

#### **2.3 使用 defer 语句来记录函数的参数与返回值**
> 下面的代码展示了另一种在调试时使用 defer 语句的手法
```
package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}() // %q: 带双引号的字符串
	return 7, io.EOF
}

func main() {
	func1("Go")
}
```
*输出结果：* `2019/12/20 16:55:26 func1("Go") = 7, EOF`

#### **三， 递归函数**

#### **3.1 递归函数的定义**
* 当一个函数在其函数体内调用自身，则称之为递归。最经典的例子便是计算斐波那契数列，即前两个数为1，从第三个数开始每个数均为前两个数之和。

>数列如下所示：
`1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …`

##### *_示例 - 7:_*
> 下面的程序可用于生成该数列
```
package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
```
*输出结果：*
```
fibonacci(0) is: 1
fibonacci(1) is: 1
fibonacci(2) is: 2
fibonacci(3) is: 3
fibonacci(4) is: 5
fibonacci(5) is: 8
fibonacci(6) is: 13
fibonacci(7) is: 21
fibonacci(8) is: 34
fibonacci(9) is: 55
fibonacci(10) is: 89
```

#### **3.2 递归函数的使用**

* 许多问题都可以使用优雅的递归来解决，比如说著名的快速排序算法

* 在使用递归函数时经常会遇到的一个重要问题就是栈溢出：
  一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为[**懒惰求值**](https://www.w3cschool.cn/functional_programm_for_rest/4zygvozt.html)的技术解决，在 Go 语言中，我们可以使用管道（channel）和 goroutine来实现。

* Go 语言中也可以使用相互调用的递归函数：
  多个函数之间相互调用形成闭环。因为 Go 语言编译器的特殊性，这些函数的声明顺序可以是任意的。
##### *_示例 - 8:_*
> 下面这个简单的例子展示了函数 odd 和 even 之间的相互调用
```
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
```
#### **3.3 递归函数 - 练习题**

* **练习 - 1**：重写本节中生成斐波那契数列的程序并返回两个命名返回值，即数列中的位置和对应的值，例如 5 与 4，89 与 10。
* **练习 - 2**：使用递归函数从 10 打印到 1。
* **练习 - 3**：实现一个输出前 30 个整数的阶乘的程序。
> n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。

> 然后，使用命名返回值来实现这个程序的第二个版本。

> 特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？
  最好的解决方案就是使用 big 包。

