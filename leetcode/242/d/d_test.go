// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[-1,2,-3,4,-5]`, 
			`5`,
		},
		{
			`[7,-6,5,10,5,-2,-6]`, 
			`13`,
		},
		{
			`[-10,-12]`, 
			`-22`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, stoneGameVIII, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-242/problems/stone-game-viii/
