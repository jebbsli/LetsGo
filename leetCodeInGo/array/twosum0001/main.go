package main

import "fmt"

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

题目大意 #
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
*/

func twoSum1(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1 + num2 == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int  {
	m := make(map[int]int, 0)
	for i, num := range nums {
		another := target - num
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}

		m[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum1(nums, target)
	fmt.Printf("result: %v\n", result)

	result = twoSum2(nums, target)
	fmt.Printf("result: %v\n", result)

}
