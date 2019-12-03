### **变量err， 测试多返回值函数的错误**

*介绍* ：
Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败。当不使用 true 或 false 的时候，也可以使用一个 error 类型的变量来代替作为第二个返回值。
* 成功执行的话，error 的值为 nil，否则就会包含相应的错误信息（Go 语言中的错误类型为 error: var err error）。这样一来，就很明显需要用一个 if 语句来测试执行结果；由于其符号的原因，这样的形式又称之为 comma,ok 模式（pattern）。
* 当没有错误发生时，代码继续运行就是唯一要做的事情，所以 if 语句块后面不需要使用 else 分支。

#### **示例1**
```
value, err := pack1.Function1(param1)
if err != nil {
	fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
	return err
}
// 未发生错误，继续执行：
```
* 由于本例的函数调用者属于 main 函数，所以程序会直接停止运行。
*如果我们想要在错误发生的同时终止程序的运行，我们可以使用 os 包的 Exit 函数：*
```
if err != nil {
	fmt.Printf("Program stopping with error %v", err)
	os.Exit(1) // 此处的退出代码 1 可以使用外部脚本获取到
}
```
#### **示例2**
可以将错误的获取放置在 if 语句的初始化部分：
```
if err := file.Chmod(0664); err != nil {
	fmt.Println(err)
	return err
}
```

