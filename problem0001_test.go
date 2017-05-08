/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-07 20:26:21
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-08 01:32:18
*/

package projecteuler_test

import (
	"testing"

	"github.com/pwarmuzStvns/projecteuler"
)

func TestProblem0001(t *testing.T) {
	testCases := []struct {
		limits  int
		answers int
	}{
		{1000, 233168},
		{10, 23},
		{10, 33},
	}
	var theory, programmatic int
	for _, values := range testCases {
		theory, programmatic = projecteuler.Problem0001(values.limits)
		if theory != values.answers || programmatic != values.answers {
			t.Error(
				"For:", values.limits,
				"Want:", values.answers,
				"Got theory:", theory,
				"Got programmatic:", programmatic)
		}
	}
}
