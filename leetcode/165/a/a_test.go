// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`[[0,0],[2,0],[1,1],[2,1],[2,2]]`}, {`[[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]`}, {`[[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]]`}, {`[[0,0],[1,1]]`}}
	sampleOuts := [][]string{{`"A"`}, {`"B"`}, {`"Draw"`}, {`"Pending"`}}
	if err := testutil.RunLeetCodeFunc(t, tictactoe, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
