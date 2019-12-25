### **Go语言专项练习 - 函数**

#### **练习-1：** 可变参数的使用

> 写一个函数add 支持1个或多个int相加，并返回相加结果
> 写一个函数concat，支持1个或多个string拼接，并返回结果

#### 答案 - 1.1

```
package main

import (. "fmt")

	func Add(a int,b...int)int{
		sum:=a
		for i:=0;i<len(b);i++{
		sum+=b[i]
		}
	return sum
	}
	
func main(){
Println(Add(1,3,5,7,9))
}
```
*输出结果：* `25`

#### 答案 - 1.2
```
package main

import (. "fmt")

	func concat(a string, arg ...string) (result string) {
    result = a
    for i := 0; i < len(arg); i++ {
        result += arg[i]
    }
    return
}
	
func main(){
    res:=concat("Hello"," ","Golang!")
    Println(res)
}
```
*输出结果：* `Hello Golang`

#### **练习-2：**

>  统计一段字符串，中文，字符，数字，空格，出现的次数。

#### **练习-3**
> 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。


