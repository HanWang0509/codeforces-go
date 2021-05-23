// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,3,2]`, `6`, 
			`1`,
		},
		{
			`[1,3,2]`, `2.7`, 
			`3`,
		},
		{
			`[1,3,2]`, `1.9`, 
			`-1`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minSpeedOnTime, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-242/problems/minimum-speed-to-arrive-on-time/
