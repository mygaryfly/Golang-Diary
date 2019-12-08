### GO语言基础语法

> [GO基础语法](https://www.runoob.com/go/go-basic-syntax.html)

#### **语法结构**

* Go语言主要有四种类型的声明语句：var,	const,	type	和	func,	分别对应变量,	常量,	类型和函数实体 对象的声明
```
//	当前程序的包名(一个可执行程序只有一个main	包) 
//  一般建议package的名称和目录名保持一致 
    package	main
//	导入其它包
//	缺少或未使用的包,程序都无法编译通过 
    import "fmt"
//	通过 const 关键字来进行常量的定义 
    const number1 =10
//	通过 var 关键字来声明变量 
    var	number2	= 20
//	数组 
    var	number3	=[5]int{1,3,5,7,9}
//	集合 
    var	number4	= map[string]int{"Age1":18,"Age2":19,"Age2":20,}
//	一般类型声明 
    type number5 int
//	结构声明 
    type number6 struct{}
//	接口声明 
    type number7	interface{}
//	通过	
    func 关键字来进行函数的声明 
//	只有 package 名称为 main 的包才可以包含 main 函数 
    func main()	{
    fmt.Println("Hello	World") }
```

#### **表达式**
a++ / a-- 这些是表达式，不是值。因此在引用变量的时候：
不可以使用` var a = a++` 这样的写法。a++ / a--作为表达式可单独写在一行
* **举例**
```
func main(){
a:=1
a++
fmt.Println(a)
// 输出的结果为2
}
```