package leetcode

/*
> https://leetcode.com/problems/two-sum/description/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

给定一个数组和一个数，从数组里面找出两个数，他们的和是这个数，返回这两个数字的index（答案唯一，同一个数只能使用一次）

Example.
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

*/

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func func1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func func2(nums []int, target int) []int {
	var mapInt = make(map[int]int, len(nums))
	for k, v := range nums {
		mapInt[v] = k
	}

	for k, v := range nums {
		if j, ok := mapInt[target-v]; ok && k != j {
			return []int{k, j}
		}
	}

	return []int{}
}

func func3(nums []int, target int) []int {
	var mapInt = make(map[int]int, len(nums))
	for k, v := range nums {
		if j, ok := mapInt[target-v]; ok {
			return []int{k, j}
		}
		mapInt[v] = k
	}

	return []int{}
}

func Test_1(t *testing.T) {
	test.Runs(t, func1, []*test.Case{
		{Input: `[2, 7, 11, 15]\n9`, Output: "[0, 1]"},
	})
	test.Runs(t, func2, []*test.Case{
		{Input: `[2, 7, 11, 15]\n9`, Output: "[0, 1]"},
	})
	test.Runs(t, func3, []*test.Case{
		{Input: `[2, 7, 11, 15]\n9`, Output: "[1, 0]"},
	})
}
