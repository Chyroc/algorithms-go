package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_914(t *testing.T) {
	test.Runs(t, hasGroupsSizeX, []*test.Case{
		{Input: `[1,2,3,4,4,3,2,1]`, Output: `true`},
		{Input: `[1,1,1,2,2,2,3,3]`, Output: `false`},
		{Input: `[1]`, Output: `false`},
		{Input: `[1,1]`, Output: `true`},
		{Input: `[1,1,2,2,2,2]`, Output: `true`},
	})
}
