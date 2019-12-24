###  **函数**

##### *小贴士 1：*
Go是编译型语言，所以函数编写的顺序是无关紧要的；鉴于可读性的需求，最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）

此外：好的程序是非常注意DRY原则的，即不要重复你自己（Don't Repeat Yourself），意思是执行特定任务的代码只能在程序里面出现一次。

#### **一，介绍**

Go 里面有三种类型的函数
* 普通的带有名字的函数
* 匿名函数或者lambda函数
* 方法（Methods)

##### **常见函数形式**
```
func test(){}                 //无参数，无返回值函数
func test( a int){}          //有参数，无返回值函数
func test(args...int)       //...表示这个是不定参数
func test() int { return 10}                  //有参数，有返回值函数

func ftest(){  

    func (s string){          //匿名函数，函数内部直接执行
        fmt.Println(s)
    }("hello world")

}

a := func(a int){             //函数    可以赋值变量                                                                    
    fmt.Println(a)
}
a(1000)

func test1() func(int) int{    //函数 作为返回值
    
    return func(a int) int{
        
        return  100
    }
}

func test3(a func() int){       //函数作为参数       闭包

        fmt.Println(a)               //函数作为参数  打印0x1093620  函数为引用类型
    fmt.Println(a())
}

e := 2222333
test3(func() int {                

    return  e
})
```
* 【小贴士-a】不定参数的位置一定只能位于形参中的最后一个！
* 【小贴士-b】固定参数一定要传参，不定参数根据需求来传参！
* 【小贴士-c】不定参数实参的传递
> 比如：
```
package main

import (. "fmt")

func main(){
    test0(1,2,3,4,5)
    }

    func test0(args...int){
		test1(args...) //全部元素传递给tets1
		}
		
    func test1(temp...int){
        for _,data:=range temp{ //设定temp集合的下标为返回空值，只返回集合体本身内容
			Println("data=",data)
		}
       
    }
```
*输出结果a：*
```
data= 1
data= 2
data= 3
data= 4
data= 5
```
> 如果test1改为：
`test1(args[:2]...) //抓取第一个参数到第三个参数且不包括第三个参数`
*输出结果b*
```
data= 1
data= 2
```
> 如果test1改为：
`test1(args[2:]...) //抓取第三个参数到最后一个函数，且包括第三个参数`
*输出结果c*
```
data= 3
data= 4
data= 5
```

#### *_1.1_* 匿名函数
* 匿名函数是指不需要定义函数名的一种函数实现方式。
* 在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。
* 匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必声明。

##### **定义匿名函数的方法**

第一种，先声明一个函数类型的变量，然后定义一个匿名函数

第二种，使用更简略的 ":=" 方式定义一个匿名函数

##### **匿名函数的使用方式**

**1.1.1** 在定义匿名函数的时候就可以直接使用（这种方式只使用一次）

```
package main
import (
    "fmt"
)
func main(){
    res1 := func (n1 int, n2 int) int {
        return n1 + n2
    }(10, 30)  //括号里的10,30 就相当于参数列表，分别对应n1和n2
    
    fmt.Println("res1=",res1)
}
```
*输出结果:* `res1= 40`

**1.1.2** 将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
```
package main
import (
    "fmt"
)
func main(){
    //将匿名函数fun 赋给变量test_fun
    //则test_fun的数据类型是函数类型，可以通过test_fun完成调用
    test_fun := func (n1 int, n2 int) int {
        return n1 - n2
    }

    res2 := test_fun(10, 30)
    res3 := test_fun(50, 30)
    fmt.Println("res2=", res2)
    fmt.Println("res3=", res3)
    fmt.Printf("%T", test_fun)
}
```
*输出结果：*
```
res2= -20
res3= 20
func(int, int) int
```
**1.1.3** 全局匿名函数:就是将匿名函数赋给一个全局变量，那么这个匿名函数在当前程序里可以使用

```
package main
import (
    "fmt"
)

//Test_fun 就是定义好的全局变量
//全局变量必须首字母大写
var (
    Test_fun = func (n1 int, n2 int) int {
        return n1 - n2
    }
)
func main(){
    val1 := Test_fun(9, 7)

    fmt.Println("val1=", val1)
}
```



##### **匿名函数应用**
1、匿名函数的变量为函数地址

2、直接创建匿名函数执行并返回结果

#### *_1.2_* Go语言中的方法 
[传送门](./day14.md)

#### **二，函数的使用**

#### *_2.1_* 除了main()、init()函数外，其它所有类型的函数都可以有参数与返回值。函数参数、返回值以及它们的类型被统称为函数签名。

##### *小贴士 2：* 函数可以没有参数或接受多个参数。
##### *示例 - 1*
> 在本例中，add 接受两个 int 类型的参数
```
package main

import "fmt"
//注意类型在变量名~之后
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
*输出结果:* `55`

* 【重要】当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

*_在本例中_* `x int, y int` 可以被缩写为`x, y int`

#### *_2.2_* 函数的调用

函数被调用的基本格式如下：
`pack1.Function(arg1, arg2, …, argn)`

Function 是 pack1 包里面的一个函数，括号里的是被调用函数的实参（argument）：这些值被传递给被调用函数的形参（parameter）。函数被调用的时候，这些实参将被复制（简单而言）然后传递给被调用函数。函数一般是在其他函数里面被调用的，这个其他函数被称为调用函数（calling function）。函数能多次调用其他函数，这些被调用函数按顺序（简单而言）执行，理论上，函数调用其他函数的次数是无穷的（直到函数调用栈被耗尽）。

##### *示例 - 2*
> 一个简单的函数调用其他函数的例子：
```
package main

func main() {
    println("In main before calling greeting")
    greeting()
    println("In main after calling greeting")
}

func greeting() {
    println("In greeting: Hi!!!!!")
}
```
*输出结果：*
```
In main before calling greeting
In greeting: Hi!!!!!
In main after calling greeting
```

##### *小贴士 3：*
函数可以将其他函数调用作为它的参数，只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的
> 例如：
假设 f1 需要 3 个参数 `f1(a, b, c int)`，同时 f2 返回 3 个参数` f2(a, b int) (int, int, int)`，就可以这样调用 `f1：f1(f2(a, b))`。

##### *小贴士 4：*
函数重载（function overloading）指的是可以编写多个同名函数，只要它们拥有不同的形参与/或者不同的返回值，在 Go 里面函数重载是不被允许的。这将导致一个编译错误：
`funcName redeclared in this book, previous declaration at lineno`

#### *_2.3_* 函数的声明

如果需要申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体：

```func flushICache(begin, end uintptr) // implemented externally```

函数也可以以申明的方式被使用，作为一个函数类型，就像：

```type binOp func(int, int) int```

在这里，不需要函数体 `{}`
```
// 阐述函数本身作为变量的类型，来申明某变量是函数类型的变量

package main

import "fmt"

func Add (a,b int)int{
	return a+b
}

func Minus (c,d int)int{
	return c-d
}
// 函数也是一种类型，通过type给一个函数类型起名
type FuncType func(int,int)int // FuncTpye是一个函数类型，func后没有函数名，也没有函数体。

func main () {
	var result int
	result = Add(1,1) //函数传统的调用方式
	fmt.Println("加法result = ",result)
// 声明一个函数类型的变量名，它的名称叫fTest
	var fTest FuncType
	fTest = Minus // 是变量就可以赋值
	result = fTest(30,20) // 相当于 Minus(30,20)
	fmt.Println("减法result = ",result)
}
```
* 函数是一等值（first-class value）：它们可以赋值给变量，就像 `add := binOp` 一样。意思函数可以被当做一个值赋值给变量，或者说可以把一个变量定义为某函数。
>这个变量知道自己指向的函数的签名，所以给它赋一个具有不同签名的函数值是不可能的。
* 函数值（functions value）之间可以相互比较：如果它们引用的是相同的函数或者都是 nil 的话，则认为它们是相同的函数。函数不能在其它函数里面声明（不能嵌套），不过我们可以通过使用匿名函数来破除这个限制。

* 目前 Go 没有泛型（generic）的概念，也就是说它不支持那种支持多种类型的函数。不过在大部分情况下可以通过接口（interface），特别是空接口与类型选择（type switch）与/或者通过使用反射（reflection）来实现相似的功能。使用这些技术将导致代码更为复杂、性能更为低下，所以在非常注意性能的的场合，最好是为每一个类型单独创建一个函数，而且代码可读性更强。

#### *_2.3.2_* 函数的返回值

函数能够接收参数供自己使用，也可以返回零个或多个值（我们通常把返回多个值称为返回一组值）。多值返回是 Go 的一大特性，为我们判断一个函数是否正常执行提供了方便。

* **【重要】**：我们通过 `return` 关键字返回一组值。事实上，任何一个有返回值（单个或多个）的函数都必须以 `return` 或 `panic` 结尾。



#### *_2.3.2_* 多值返回
##### *示例 - 3*
让我们看三个函数：一个没有返回值，一个有一个返回值，一个有两个返回值。
```
func log(message string) {
}

func add(a int, b int) int {
}

func power(name string) (int, bool) {
}
```
> 我们可以像这样使用最后一个：
```
value, exists := power("goku")
if exists == false {
  // 处理错误情况
}
```
* 【重要】有时候，你仅仅关注其中一个返回值。这个情况下，你可以将其他的返回值赋值给空白符。这不仅仅是一个惯例。_ ，空白标识符，特殊在于实际上返回值并没有赋值。这让你可以一遍又一遍地使用 _ 而不用管它的类型
```
_, exists := power("goku")
if exists == false {
  // handle this error case
}
```
##### *示例 - 4* swap 函数返回了两个字符串。
> 函数可以返回任意数量的返回值。

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```
#### *_2.3.3_* 命名返回值

* Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
* 返回值的名称应当具有一定的意义，它可以作为文档使用。
* 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
* 【重要】直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

##### *示例 - 5* 函数的返回值被命名为sum
> 此时sum可以被看作为函数顶部的变量
```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
*输出结果* `7 10`

#### *_2.3.4_* 函数的变量
* var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

##### *示例 - 6* 函数体的内部变量和外部变量
>就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。
```
package main

import "fmt"

var c, python, java bool // 全局变量，函数体的外部变量

func main() {
	var i int // 局部变量，函数体的内部变量
	fmt.Println(i, c, python, java)
}
```
* **【值得注意的是】** 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。

* **零值**：没有明确初始值的变量声明会被赋予它们的 零值。
```
零值是：

    数值类型为 0，
    布尔类型为 false，
    字符串为 ""（空字符串）。
```
> 比如：
```
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```
*输出结果：* `0 0 false ""`

**【重要】** 类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法或 `var =` 表达式语法），变量的类型由右值推导得出。

> 比如：
```
package main

import "fmt"

func main() {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
```
*输出结果：*
```
i is of type int
f is of type float64
g is of type complex128
```


