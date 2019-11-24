#### 变量
在 Go 中，可以很轻松地将它们都声明为指针类型：
var a, b *int

其次，这种语法能够按照从左至右的顺序阅读，使得代码更加容易理解。

示例：

var a int
var b bool
var str string
你也可以改写成这种形式：

``` 
var (
	a int
	b bool
	str string
)
```
这种因式分解关键字的写法一般用于声明全局变量。当你在函数体内声明局部变量时，应使用简短声明语法 :=，例如：
                       
                       a := 1
##### 简短形式，使用 := 赋值操作符
简短新式的声明语句写上 var 关键字就显得有些多余了，比如我们可以将它们简写为 a := 50 或 b := false。

a 和 b 的类型（int 和 bool）将由编译器自动推断
##### 注意事项：
如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。

如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。

##### 练习一：
```
//定义全局变量a，写死为"G"
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
//这里引用局部变量a的值与全局变量没有冲突
   a := "O"
   print(a)
}
```
##### 练习二：
```
//本案例中直接定义全局变量a为string类型
package main

var a string

func main() {
//对a进行赋值
   a = "G"
   print(a)
   f1()
}

func f1() {
//对第二个函数f1中，局部变量a进行赋值
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}
```
