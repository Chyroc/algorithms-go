package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_650(t *testing.T) {
	test.Runs(t, minSteps, []*test.Case{
		{Input: "3", Output: "3"},
	})
}
