// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[3,5,1]`, 
			`true`,
		},
		{
			`[1,2,4]`, 
			`false`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, canMakeArithmeticProgression, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-196/problems/can-make-arithmetic-progression-from-sequence/
