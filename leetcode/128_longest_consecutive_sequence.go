package leetcode

/*
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (39.03%)
 * Total Accepted:    155.8K
 * Total Submissions: 398.6K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers, find the length of the longest
 * consecutive elements sequence.
 *
 * Your algorithm should run in O(n) complexity.
 *
 * Example:
 *
 *
 * Input: [100, 4, 200, 1, 3, 2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 */

/*


* 给定一个无序数组，求连续元素的最大长度，要求时间复杂度 O(n)
  * 用一个map[int]int 实现
    * 存储连续数组的左右边界 -> 当前连续的长度
    * 有的时候元素只有一个，并且为了避免重复计算，已经计算过的要加入到m
*/

func getOrDefault(m map[int]int, key, d int) int {
	d2, ok := m[key]
	if ok {
		return d2
	} else {
		return d
	}
}

func longestConsecutive(nums []int) int {
	var res int
	var m = make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			continue
		}
		left := getOrDefault(m, n-1, 0)
		right := getOrDefault(m, n+1, 0)

		sum := left + right + 1
		res = max(res, sum)

		m[n] = sum
		m[n-left] = sum
		m[n+right] = sum
	}

	return res
}
