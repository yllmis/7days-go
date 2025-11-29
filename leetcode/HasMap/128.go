package main

import (
	"fmt"
	"slices"
)

func longestConsecutive(nums []int) int {
	var lens, ans int
	// step1: 哈希表键存数组数字，值存个数（无用）
	num := make(map[int]int, 0)
	// step2: 遍历数组，存进哈希表
	for _, value := range nums {
		num[value]++
	}

	num1 := make([]int, 0)
	// step3：根据键排序，获得最长序列
	for k := range num {
		num1 = append(num1, k)
	}
	slices.Sort(num1)

	fmt.Println(num1)
	for i, _ := range num1 {

		if i+1 < len(num1) && num1[i] == num1[i+1]-1 {
			lens++
		} else {
			fmt.Println(lens)
			if ans < (lens + 1) {
				ans = lens + 1
			}
			lens = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

}
