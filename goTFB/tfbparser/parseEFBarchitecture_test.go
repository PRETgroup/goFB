package tfbparser

import (
	"testing"
)

var efbArchitectureTests = []ParseTest{
	{
		Name: "missing brace after s1",
		Input: `enforceFB testBlock;
				interface of testBlock{
				}
				architecture of testBlock {
					states {
						s1 

					}
				}`,
		Err: ErrUnexpectedValue,
	},
}

func TestParseEFBArchitecture(t *testing.T) {
	runParseTests(t, efbArchitectureTests)
}
