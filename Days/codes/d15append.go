//【练习2.4 - 1】 给定切片 sl，将一个 []byte 数组追加到 sl 后面。写一个函数 `Append(slice, data []byte) []byte`，该函数在 sl 不能存储更多数据的时候自动扩容。
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

//练习 4.2 - 1

//给定 slice s[]int 和一个 int 类型的因子factor，扩展 s 使其长度为 len(s) * factor。
package main

import "fmt"

var s = []int{2,3,4}

func main(){
	fx(6)
}
	func fx(factor int){
		newS:=make([]int,len(s)*(factor-1))
		s=append(s,newS...)
		fmt.Print(s)
	}
//练习 4.2 - 2

//用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。

//练习 4.2 - 3

//写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。

//练习 4.2 - 4

//写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除
