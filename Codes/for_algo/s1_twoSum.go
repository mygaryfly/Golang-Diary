package main

import (
	"fmt"
)

var nums = []int{2, 7, 11, 15}
var target = 9

func main() {

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	res := []int{}
	for _, v1 := range nums {
		for _, v2 := range nums {
			if (v1 + v2) == target {
				res = []int{v1, v2}
				break
			}
		}
	}
	fmt.Println(res)
	return res

}
