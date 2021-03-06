package leetcode

/*
 * [696] Count Binary Substrings
 *
 * https://leetcode.com/problems/count-binary-substrings/description/
 *
 * algorithms
 * Easy (51.05%)
 * Total Accepted:    19.8K
 * Total Submissions: 38.8K
 * Testcase Example:  '"00110"'
 *
 * Give a string s, count the number of non-empty (contiguous) substrings that
 * have the same number of 0's and 1's, and all the 0's and all the 1's in
 * these substrings are grouped consecutively.
 *
 * Substrings that occur multiple times are counted the number of times they
 * occur.
 *
 * Example 1:
 *
 * Input: "00110011"
 * Output: 6
 * Explanation: There are 6 substrings that have equal number of consecutive
 * 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
 * Notice that some of these substrings repeat and are counted the number of
 * times they occur.
 * Also, "00110011" is not a valid substring because all the 0's (and 1's) are
 * not grouped together.
 *
 *
 *
 * Example 2:
 *
 * Input: "10101"
 * Output: 4
 * Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal
 * number of consecutive 1's and 0's.
 *
 *
 *
 * Note:
 * s.length will be between 1 and 50,000.
 * s will only consist of "0" or "1" characters.
 *
 */

/*


* 给一个字符串，里面只有0和1，求成对出现的个数（相同数字需要相邻）(696)
  * 例如："00110011"包含： "0011", "01", "1100", "10", "0011", and "01"，6个
  * 首先统计相邻的数字相同的个数，比如"00110011" -> "2，2，2，2" group
  * 循环这个（`k<n-1`）
    * 如果group[i] > group[i+1]，个数加group[i+1]，否则加group[i]

注意：

* 字符串的长度再1-50000之间
* 字符串只包含0和1

*/

func countBinarySubstrings(s string) int {
	var group []int

	for k, v := range s {
		if k == 0 {
			group = append(group, 1)
		} else if v == rune(s[k-1]) {
			group[len(group)-1]++
		} else {
			group = append(group, 1)
		}
	}

	sum := 0
	for k, v := range group {
		if k == len(group)-1 {
			continue
		}

		if v <= group[k+1] {
			sum += v
		} else {
			sum += group[k+1]
		}
	}

	return sum
}
