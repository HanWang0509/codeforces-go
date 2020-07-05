// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[[1,0,1],[1,1,0],[1,1,0]]`, 
			`13`,
		},
		{
			`[[0,1,1,0],[0,1,1,1],[1,1,1,0]]`, 
			`24`,
		},
		{
			`[[1,1,1,1,1,1]]`, 
			`21`,
		},
		{
			`[[1,0,1],[0,1,0],[1,0,1]]`, 
			`5`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numSubmat, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-196/problems/count-submatrices-with-all-ones/
