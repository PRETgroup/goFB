package tfbparser

import (
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
)

var efbArchitectureTests = []ParseTest{
	{
		Name: "missing brace after s1",
		Input: `policyFB testBlock;
				interface of testBlock{
				}
				architecture of testBlock {
					states {
						s1 

					}
				}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "AEIPolicy",
		Input: `policyFB AEIPolicy;
				interface of AEIPolicy {
					in event AS, VS; //in here means that they're going from PLANT to CONTROLLER
					out event AP, VP;//out here means that they're going from CONTROLLER to PLANT
				
					in ulint AEI_ns := 900000000;
				}
				architecture of AEIPolicy {
					internals {
						dtimer tAEI; //DTIMER increases in DISCRETE TIME continuously
					}
				
					//P3: AS or AP must be true within AEI after a ventricular event VS or VP.
				
					states {
						s1 {
							//-> <destination> [on guard] [: output expression][, output expression...] ;
							-> s2 on (VS or VP): tAEI := 0;
						}
				
						s2 {
							-> s1 on (AS or AP);
							-> violation on (tAEI > AEI_ns);
						}
					} 
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.NewPolicyFB("AEIPolicy").
					AddEventInputNames([]string{"AS", "VS"}, d).
					AddEventOutputNames([]string{"AP", "VP"}, d).
					AddDataInputs([]string{"AEI_ns"}, []string{}, "ulint", "", "900000000", d)).
				AddPFBState("s1", d).
				AddPFBState("s2", d).
				AddPFBDataInternals([]string{"tAEI"}, "DTIMER", "", "", d).
				AddPFBTransition("s1", "s2", "( VS || VP )", []iec61499.PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
				AddPFBTransition("s2", "s1", "( AS || AP )", nil, d).
				AddPFBTransition("s2", "violation", "( tAEI > AEI_ns )", nil, d),
		},
	},
}

func TestParsePFBArchitecture(t *testing.T) {
	runParseTests(t, efbArchitectureTests)
}
