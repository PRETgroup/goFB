package tfbparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kiwih/goFB/iec61499"
)

type ParseTest struct {
	Name   string
	Input  string
	Output []iec61499.FB //if applicable
	Err    error         //if applicable
}

var d = iec61499.DebugInfo{
	SourceLine: -1,
	SourceFile: "testing",
}

var basicTests = []ParseTest{
	{
		Name: "simple typo 1",
		Input: `basicFB testBlock;
				interface of asdasd {}`,
		Err: ErrUndefinedFB,
	},
	{
		Name:  "simple typo 2",
		Input: `dadasdasd`,
		Err:   ErrUnexpectedValue,
	},
	{
		Name: "simple typo 3",
		Input: `basicFB testBlock1, , testBlock3;
				interface of testBlock2 {}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "simple typo 4",
		Input: `basicFB testBlock1, testBlock2;
				interface of testBlock2;`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "simple typo 5",
		Input: `basicFB testBlock1, testBlock2;
				interface of testBlock2 {`,
		Err: ErrUnexpectedEOF,
	},
	{
		Name: "simple typo 6",
		Input: `basicFB testBlock1, testBlock2;
				interface of testBlock2 {}
				architecture of testBlock2 {`,
		Err: ErrUnexpectedEOF,
	},
	{
		Name: "simple typo 7",
		Input: `basicFB testBlock1;
				architecture of asdasd {}`,
		Err: ErrUndefinedFB,
	},
	{
		Name: "simple typo 8",
		Input: `basicFB testBlock1, testBlock2;
				architecture testBlock2 {}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "missing word 1",
		Input: `basicFB testBlock1;
				interface testBlock1 {}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "empty interface 1",
		Input: `basicFB testBlock;
				interface of testBlock {}`,
		Output: []iec61499.FB{*iec61499.NewBasicFB("testBlock")},
		Err:    nil,
	},
	{
		Name: "empty interfaces 1",
		Input: `basicFB testBlock1, testBlock2, testBlock3;
				interface of testBlock2 {}`,
		Output: []iec61499.FB{*iec61499.NewBasicFB("testBlock1"), *iec61499.NewBasicFB("testBlock2"), *iec61499.NewBasicFB("testBlock3")},
		Err:    nil,
	},
}

func fbSliceEqual(a []iec61499.FB, b []iec61499.FB) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if same := iec61499.FBsEqual(a[i], b[i]); same != true {
			return false
		}
	}
	return true
}

func runParseTests(t *testing.T, pTests []ParseTest) {
	for i, test := range pTests {
		out, err := ParseString(fmt.Sprintf("Test[%d]", i), test.Input)
		if err != nil && test.Err == nil {
			t.Errorf("Test[%d](%s): Error '%s' occurred when it shouldn't have", i, test.Name, err.Error())
		} else if err == nil && test.Err != nil {
			t.Errorf("Test[%d](%s): Error didn't occur and it should have been '%s'", i, test.Name, test.Err.Error())
		} else if err != nil && test.Err != nil {
			if err.Err.Error() != test.Err.Error() {
				t.Errorf("Test[%d](%s): Error codes don't match (it was '%s', should have been '%s')", i, test.Name, err.Error(), test.Err.Error())
			}
		} else if err == nil && test.Err == nil {
			if len(out) != len(test.Output) {
				t.Errorf("Test[%d](%s): Output length of FB slices don't match! (was '%v' should have been '%v')", i, test.Name, len(out), len(test.Output))
			} else {
				if !fbSliceEqual(out, test.Output) {
					t.Errorf("Test[%d](%s): Outputs don't match!", i, test.Name)

					bytes, _ := json.MarshalIndent(test.Output, "", "\t")
					//fmt.Printf("\n\nDesired:\n%s", bytes)

					ioutil.WriteFile("test_desired.out.json", bytes, 0644)

					bytes, _ = json.MarshalIndent(out, "", "\t")
					ioutil.WriteFile("test_actual.out.json", bytes, 0644)

				}
			}
		}

	}
}

func TestParseBasics(t *testing.T) {
	runParseTests(t, basicTests)
}
