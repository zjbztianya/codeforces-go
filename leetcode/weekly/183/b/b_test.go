// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`"1101"`}, {`"10"`}, {`"1"`}}
	exampleOuts := [][]string{{`6`}, {`1`}, {`0`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	exampleIns = append(exampleIns, []string{`"111"`})
	exampleOuts = append(exampleOuts, []string{`1`})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, numSteps, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-183/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/