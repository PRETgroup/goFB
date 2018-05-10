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
				policy of testBlock {
					states {
						s1 

					}
				}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "AEIPolicy",
		Input: `basicFB AEIPolicy;
				interface of AEIPolicy {
					in event AS, VS; //in here means that they're going from PLANT to CONTROLLER
					out event AP, VP;//out here means that they're going from CONTROLLER to PLANT
				
					in ulint AEI_ns := 900000000;
				}
				policy AEI of AEIPolicy {
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
				iec61499.NewBasicFB("AEIPolicy").
					AddEventInputNames([]string{"AS", "VS"}, d).
					AddEventOutputNames([]string{"AP", "VP"}, d).
					AddDataInputs([]string{"AEI_ns"}, []string{}, "ulint", "", "900000000", d)).
				AddPolicy("AEI").
				AddPFBState("s1", d).
				AddPFBState("s2", d).
				AddPFBDataInternals([]string{"tAEI"}, "DTIMER", "", "", d).
				AddPFBTransition("s1", "s2", "( VS or VP )", []iec61499.PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
				AddPFBTransition("s2", "s1", "( AS or AP )", nil, d).
				AddPFBTransition("s2", "violation", "( tAEI > AEI_ns )", nil, d),
		},
	},
	{
		Name: "AB5Policy",
		Input: `basicFB AB5Policy;
			interface of AB5Policy {
				in event A;  //in here means that they're going from PLANT to CONTROLLER
				out event B; //out here means that they're going from CONTROLLER to PLANT
			}
			
			policy AB5 of AB5Policy {
				internals {
					dtimer v;
				}
			
				states {
					s0 {														//first state is initial, and represents "We're waiting for an A"
						-> s0 on (!A and !B): v := 0;							//if we receive neither A nor B, do nothing
						-> s1 on (A and !B): v := 0;							//if we receive an A only, head to state s1
						-> violation on ((!A and B) or (A and B));				//if we receive a B, or an A and a B (i.e. if we receive a B) then VIOLATION
					}
			
					s1 {														//s1 is "we're waiting for a B, and it needs to get here within 5 ticks"
						-> s1 on (!A and !B and v < 5);							//if we receive nothing, and we aren't over-time, then we do nothing
						-> s0 on (!A and B);									//if we receive a B only, head to state s0
						-> violation on ((v >= 5) or (A and B) or (A and !B));	//if we go overtime, or we receive another A, then VIOLATION
					}
				}
			}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("AB5Policy").
				AddEventInputNames([]string{"A"}, d).
				AddEventOutputNames([]string{"B"}, d).
				AddPolicy("AB5").
				AddPFBState("s0", d).
				AddPFBState("s1", d).
				AddPFBDataInternals([]string{"v"}, "DTIMER", "", "", d).
				AddPFBTransition("s0", "s0", "( !A and !B )", []iec61499.PFBExpression{{VarName: "v", Value: "0"}}, d).
				AddPFBTransition("s0", "s1", "( A and !B )", []iec61499.PFBExpression{{VarName: "v", Value: "0"}}, d).
				AddPFBTransition("s0", "violation", "( ( !A and B ) or ( A and B ) )", nil, d).
				AddPFBTransition("s1", "s1", "( !A and !B and v < 5 )", nil, d).
				AddPFBTransition("s1", "s0", "( !A and B )", nil, d).
				AddPFBTransition("s1", "violation", "( ( v >= 5 ) or ( A and B ) or ( A and !B ) )", nil, d),
		},
	},
}

func TestParsePFBArchitecture(t *testing.T) {
	runParseTests(t, efbArchitectureTests)
}
