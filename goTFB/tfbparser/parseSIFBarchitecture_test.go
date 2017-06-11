package tfbparser

import (
	"testing"

	"github.com/kiwih/goFB/iec61499"
)

var sifbArchitectureTests = []ParseTest{
	{
		Name: "basic sifb arch 1",
		Input: `serviceFB testBlock;
				interface of testBlock {
					in event inEvent;
					in lreal inData with inEvent;
					out event outEvent;
					out lreal outData with outEvent;
				}
				architecture of testBlock {
					in "C";

					in_struct ` + "`" + `//0` + "`" + `;
					pre_init ` + "`" + `//1` + "`" + `;
					init ` + "`" + `//2` + "`" + `;
					run ` + "`" + `//3` + "`" + `;
					shutdown ` + "`" + `//4` + "`" + `;
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewServiceFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddSIFBParams("C", "//0", "//1", "//2", "//3", "//4", d),
		},
		Err: nil,
	},
}

func TestParseSIFBArchitecture(t *testing.T) {
	runParseTests(t, sifbArchitectureTests)
}
