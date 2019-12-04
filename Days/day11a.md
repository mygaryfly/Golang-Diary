### **stitch的高级应用**

#### **一，Type Switch**

**_介绍：_**
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

Type Switch 语法格式如下：

```
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    /* 你可以定义任意个数的case */
    default: /* 可选 */
       statement(s);
}
```

**示例**
```
package main

import "fmt"

func main() {
   var x interface{}
     
   switch i := x.(type) {
      case nil:  
         fmt.Printf(" x 的类型 :%T",i)                
      case int:  
         fmt.Printf("x 是 int 型")                      
      case float64:
         fmt.Printf("x 是 float64 型")          
      case func(int) float64:
         fmt.Printf("x 是 func(int) 型")                      
      case bool, string:
         fmt.Printf("x 是 bool 或 string 型" )      
      default:
         fmt.Printf("未知型")    
   }  
}
```

* 以上代码执行结果为：` x 的类型 :<nil> `

