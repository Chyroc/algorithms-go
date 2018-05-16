package leetcode

import (
	"fmt"
	"sort"

	"github.com/Chyroc/algorithms-go/lib"
)

/*
> https://leetcode.com/problems/subsets-ii/description/
> Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).


* 给一个重复元素数据集set，求这个set的所有子集（不能重复），如[1,2,1,3]返回[[] [3] [2] [3 2] [1] [3 1] [2 1] [3 2 1]]
  * 遍历递归
    * 基本数和78是一样的
    * 不过跳过：前面数据是相同的，但是后面在不同位置取了相同的数据
*/

func backtrackWithDup(list *[][]int, tempList []int, nums []int, start int) {
	*list = append(*list, lib.ArrayDeepCopy(tempList))
	for i := start; i < len(nums); i++ {
		// skip duplicates
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		backtrackWithDup(list, append(tempList, nums[i]), nums, i+1)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var list [][]int
	backtrackWithDup(&list, []int{}, nums, 0)
	return list
}

func Example_90() {
	fmt.Printf("%v\n", subsetsWithDup([]int{}))
	fmt.Printf("%v\n", subsetsWithDup([]int{1}))
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 2}))
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 2, 3}))
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 2, 3, 1}))

	// Output:
	// [[]]
	// [[] [1]]
	// [[] [1] [1 2] [2]]
	// [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	// [[] [1] [1 1] [1 1 2] [1 1 2 3] [1 1 3] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}
