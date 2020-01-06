### **Go语言中的map**

#### **一，map概念 - 声明、初始化和 make**

#### **1.1 概念**
* map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。

* map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。

* map 是引用类型，可以使用如下声明：
```go
var map1 map[keytype]valuetype
var map1 map[string]int
```
（`[keytype]`和 `valuetype` 之间允许有空格，但是 gofmt 移除了空格）

#### **1.2 声明、初始化**

【**重要**】

[1] - 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。

[2] - 未初始化的 map 的值是 nil。

[3] - key 可以是任意可以用 == 或者！= 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key 。但是指针和接口类型可以
* 含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的
* 如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

[4] - value 可以是任意类型的；通过使用空接口类型，我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言

> 为了说明值可以是任意类型的，这里给出了一个使用`func() int`作为值的 map
```go
package main
import "fmt"

func main() {
    mf := map[int]func() int{
        1: func() int { return 10 },
        2: func() int { return 20 },
        5: func() int { return 50 },
    }
    fmt.Println(mf)
}
```
*输出结果:*
`map[1:0x10903be0 5:0x10903ba0 2:0x10903bc0] // 整形都被映射到函数地址`


[5] - map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以*如果你很在乎性能的话还是建议用切片来解决问题*。

[6] - map 也可以用函数作为自己的值，这样就可以用来做分支结构：key 用来选择要执行的函数。

[7] - 如果 key1 是 map1 的 key，那么 map1[key1] 就是对应 key1 的值，就如同数组索引符号一样
* 数组可以视为一种简单形式的 map，key 是从 0 开始的整数
* key1 对应的值可以通过赋值符号来设置为 val1：`map1[key1] = val1`
    * 令 `v := map1[key1]` 可以将 key1 对应的值赋值为 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。

[8] - 常用的 len(map1) 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。

##### *_示例 1.2 - 1:_*
```go
package main
import "fmt"

func main() {
    var mapLit map[string]int
    //var mapCreated map[string]float32
    var mapAssigned map[string]int

    mapLit = map[string]int{"one": 1, "two": 2}
    mapCreated := make(map[string]float32)
    mapAssigned = mapLit

    mapCreated["key1"] = 4.5
    mapCreated["key2"] = 3.14159
    mapAssigned["two"] = 3

    fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
    fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
    fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
    fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}
```
*输出结果：*
```go
Map literal at "one" is: 1
Map created at "key2" is: 3.14159
Map assigned at "two" is: 3
Mpa literal at "ten" is: 0
```
> mapLit 说明了 map literals 的使用方法： map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。

[9] - map 是**引用类型**的： 内存用 make 方法来分配。

[10] - map 的初始化：`var map1 = make(map[keytype]valuetype)`。

* 或者简写为：`map1 := make(map[keytype]valuetype)`。

* 上面例子中的 mapCreated 就是用这种方式创建的：`mapCreated := make(map[string]float32)`。
    * 相当于：`mapCreated := map[string]float32{}`。
* mapAssigned 也是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值

[11] - **不要使用 new，永远用 make 来构造 map**
* 【注意】： 如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：`gomapCreated := new(map[string]float32)`

接下来当我们调用：mapCreated["key1"] = 4.5 的时候，编译器会报错：
`invalid operation: mapCreated["key1"] (index of type *map[string]float32).`

#### **1.3 map 容量**

和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。

但是你也可以选择标明 map 的初始容量 capacity，就像这样：`make(map[keytype]valuetype, cap)`
> 例如:
`map2 := make(map[string]float32, 100)`

* 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，**对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明**。

```go
noteFrequency := map[string]float32 {
    "C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
    "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
```

#### **1.4 用切片作为 map 的值**

既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？例如，当我们要处理 unix 机器上的所有进程，以父进程（pid 为整形）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题。

> 这里有一些定义这种 map 的例子：
```go
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)
```

#### **二，测试键值对是否存在及删除元素**

#### **2.1 测试键值**

测试 map1 中是否存在 key1：

我们已经见过可以使用 `val1 = map1[key1]` 的方法获取 key1 对应的值 val1。如果 map 中不存在 key1，val1 就是一个值类型的空值。

这就会给我们带来困惑了：现在我们没法区分到底是 key1 不存在还是它对应的 value 就是空值。

为了解决这个问题，我们可以这么用：`val1, isPresent = map1[key1]`

isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent为true；如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。

如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：

```go
_, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false
```

或者和 if 混合使用：

```go
if _, ok := map1[key1]; ok {
	// ...
}
```

#### **2.2 删除元素**

从 map1 中删除 key1：

直接 `delete(map1, key1)` 就可以。

如果 key1 不存在，该操作不会产生错误。

##### *_示例 2.2 - 1:_*

```go
package main
import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
```

*输出结果：*
```go
	The value of "Beijing" in map1 is: 20
	Is "Paris" in map1 ?: false
	Value is: 0
    map1 does not contain Washington
```