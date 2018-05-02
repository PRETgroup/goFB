package tfbparser

import (
	"github.com/PRETgroup/goFB/iec61499"
)

var hfbArchitectureTests = []ParseTest{
	{
		Name: "basic hfb arch 1",
		Input: `hybridFB Conveyor;
		interface of Conveyor {
			in event convOn;
			in event convOff;
			in lreal deltaTime; //compulsory input for hybridFBs.
			in lreal maxSpeed;
			out event dChange;
			out with dChange lreal initial 0 d;
		}

		architecture of Conveyor {
			internal lreal initial 0 x;
			
			locations {
				l_start {
					-> l_off on true, run ` + "`" + `x_prime = 0; d_prime = 0;` + "`" + `, emit dChange;
				}
				l_off {
					invariant ` + "`" + `0 <= x && x <= m` + "`" + `;
					run ` + "`" + `x_dot = x / 5` + "`" + `;
					run incD;
					emit dChange;
					-> l_on on convOn, run nextL, emit dChange;
				}
				l_on {
					invariant ` + "`" + `0 <= x` + "`" + `; //equivalent to invariant 0 <= x && x <= m
					invariant ` + "`" + `x <= m` + "`" + `; //
					run ` + "`" + `x_dot = (m - x) / 5` + "`" + `;
					run incD;
					emit dChange;
					-> l_off on convOff, run nextL, emit dChange;
				}
			}

			algorithm incD ` + "`" + `d_dot = x;` + "`" + `;
			algorithm nextL ` + "`" + `x_prime = x; d_prime = d;` + "`" + `;
				
		}`,
		//TODO: I've just remembered that I haven't written the parser for the transition line run/emits
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

// func TestParseHFBArchitecture(t *testing.T) {
// 	runParseTests(t, hfbArchitectureTests)
// }
