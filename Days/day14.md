### **Go语言中的结构体和方法**

#### **一，Go语言中的结构体**
* Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
* 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
* 结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
```
· Title ：标题
· Author ： 作者
· Subject：学科
· ID：书籍ID
``` 
##### **1.1 - 定义结构体**

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
```
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```
一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
```
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
```
> **比如**
```
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}


func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```
*输出结果：*
```
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com Go 语言教程 6495407}
{Go 语言 www.runoob.com  0}
```
##### *示例 - 1*
```
package main

import "fmt"

type hello struct {

	a string
	b string
}

	func main(){
	
		fmt.Println(hello{a:"Hello",b:" World!"})
		fmt.Println(hello{a:"I am ",b:"Gary!"})
	
	}
```
*输出结果:*
```
{Hello  World!}
{I am  Gary!}
```
##### **1.2 - 访问结构体成员**

*_如果要访问结构体成员，需要使用点号`.`操作符，格式为：_* `结构体.成员名`

> 结构体类型变量使用 struct 关键字定义，实例如下：
```
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
```
*输出结果：*
```
Book 1 title : Go 语言
Book 1 author : www.runoob.com
Book 1 subject : Go 语言教程
Book 1 book_id : 6495407
Book 2 title : Python 教程
Book 2 author : www.runoob.com
Book 2 subject : Python 语言教程
Book 2 book_id : 6495700
```

##### **1.3 - 结构体作为参数传递给函数**
* 你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：
> 比如：
```
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook(Book1)

   /* 打印 Book2 信息 */
   printBook(Book2)
}
    /*将结构体类型作为参数传递给函数,打印上述book1和book2的内容*/
func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
``` 
*输出结果：*
```
Book title : Go 语言
Book author : www.runoob.com
Book subject : Go 语言教程
Book book_id : 6495407
Book title : Python 教程
Book author : www.runoob.com
Book subject : Python 语言教程
Book book_id : 6495700
```

##### **1.4 - 结构体指针**
* 你可以定义指向结构体的指针类似于其他指针变量，格式如下：
 `var struct_pointer *Books`
* 以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前
`struct_pointer = &Book1`
* 使用结构体指针访问结构体成员，使用 "." 操作符：
`struct_pointer.title`

*_接下来让我们使用结构体指针重写以上实例，代码如下：_*
> 从打印 Book1 和Book2信息这里开始替换,可以得到相同的输出结果
```
/* 打印 Book1 信息 */
   printBook(&Book1)

   /* 打印 Book2 信息 */
   printBook(&Book2)
}
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
```

#### **二，Go语言中的方法**

[传送门 - 深度解析1](https://www.cnblogs.com/flying1819/articles/8832447.html)

[传送门 - 深度解析2](https://studygolang.com/articles/11393)
##### **2.1 - 介绍**
* Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
```
func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/
}
```
* Go中虽没有class，但依旧有method
* 通过显示说明receiver来实现与某个类型的结合
* 只能为同一个包中的类型定义方法
* receiver可以是类型的值或者指针
* 不存在方法重载
* 可以使用值或指针来调用方法，编译器会自动完成转换
* 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数(Method Value vs. Method Expression)
* 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
* 类型别名不会拥有底层类型所附带的方法
* 方法可以调用结构中的非公开字段
* 【重要】同一个包下的任何类型都可以声明方法，只要它的类型既不是指针类型也不是接口类型。
> 下面定义一个结构体类型和该类型的一个方法：

```
package main

import (
   "fmt"  
)

/* 定义结构体 */
type Circle struct {
  radius float64
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```
*输出结果：*
`圆的面积 =  314`

##### *示例 - 2*
```
package main

import (
    "fmt"
)

type A struct {
    Name string
}

type B struct {
    Name string
}

func main() {
    a := A{}
    a.Print()
    b := B{}
    b.Print()
}

//编译器根据接收者的类型，来判断它是属于哪个方法
func (a A) Print() {
    //取一个变量a，a就是接收者，它的接收者的类型就是structA,Print就是方法的名称，参数在Print()的括号中定义
    //receiver就是这个函数的第一个接收者，而且是强制规定的，这个时候就变成了一个方法
    fmt.Println("A")
}
func (b B) Print() {
    fmt.Println("B")
}
```