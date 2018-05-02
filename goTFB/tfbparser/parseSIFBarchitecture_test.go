package tfbparser

import (
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
)

var sifbArchitectureTests = []ParseTest{
	{
		Name: "basic sifb arch 1",
		Input: `serviceFB testBlock;
				interface of testBlock {
					in event inEvent;
					in with inEvent lreal inData;
					out event outEvent;
					out with outEvent lreal outData;
				}
				architecture of testBlock {
					in "C";

					arbitrary ` + "`" + `//0` + "`" + `;
					in_struct ` + "`" + `//1` + "`" + `;
					pre_init ` + "`" + `//2` + "`" + `;
					init ` + "`" + `//3` + "`" + `;
					run ` + "`" + `//4` + "`" + `;
					shutdown ` + "`" + `//5` + "`" + `;
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewServiceFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddSIFBParams("C", "//0", "//1", "//2", "//3", "//4", "//5", d),
		},
		Err: nil,
	},
}

func TestParseSIFBArchitecture(t *testing.T) {
	runParseTests(t, sifbArchitectureTests)
}
