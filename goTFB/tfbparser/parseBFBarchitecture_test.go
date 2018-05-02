package tfbparser

import (
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
)

var bfbArchitectureTests = []ParseTest{
	{
		Name: "basic arch 1",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent;
					in lreal inData with inEvent;
					out event outEvent;
					out lreal outData with outEvent;
				}
				architecture of testBlock {
					internal lint internalData;
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewBasicFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddBFBDataInternals([]string{"internalData"}, "lint", "", "", d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 2",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent; out event outEvent;
				}
				architecture of testBlock {
					internal byte internalData1, internalData2;
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent"}, d).
				AddBFBDataInternals([]string{"internalData1", "internalData2"}, "byte", "", "", d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 3",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent; out event outEvent;
				}
				architecture of testBlock {
					internals {
						byte internalData1, internalData2;
						int internalData3, internalData4;
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent"}, d).
				AddBFBDataInternals([]string{"internalData1", "internalData2"}, "byte", "", "", d).
				AddBFBDataInternals([]string{"internalData3", "internalData4"}, "int", "", "", d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 4",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent; out event outEvent;
				}
				architecture of testBlock {
					internals {
						byte[3] internalData1, internalData2;
						int internalData3, internalData4 := "word";
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent"}, d).
				AddBFBDataInternals([]string{"internalData1", "internalData2"}, "byte", "3", "", d).
				AddBFBDataInternals([]string{"internalData3", "internalData4"}, "int", "", "\"word\"", d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 5",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent; out event outEvent;
				}
				architecture of testBlock {
					internals {
						byte[3] internalData1, internalData2;
						int internalData3, internalData4 := 3;
					}
					internal lint[2] otherInternal := [0,1];
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent"}, d).
				AddBFBDataInternals([]string{"internalData1", "internalData2"}, "byte", "3", "", d).
				AddBFBDataInternals([]string{"internalData3", "internalData4"}, "int", "", "3", d).
				AddBFBDataInternals([]string{"otherInternal"}, "lint", "2", "[0,1]", d),
		},
		Err: nil,
	},
	{
		Name: "missing brace in arch 1",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent; out event outEvent;
				}
				architecture of testBlock {
					internals {
						byte[3] internalData1, internalData2;
						int internalData3, internalData4 := "word";
					
				}`,
		Err: ErrUnexpectedEOF, //missing closing brace
	},
	{
		Name: "association on internal 1",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent;
					in lreal inData with inEvent;
					out event outEvent;
					out lreal outData with outEvent;
				}
				architecture of testBlock {
					internal lint internalData with outEvent;
				}`,
		Err: ErrUnexpectedAssociation, //all associations are unexpected on internal data
	},
	{
		Name: "event on internal 1",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent;
					in lreal inData with inEvent;
					out event outEvent;
					out lreal outData with outEvent;
				}
				architecture of testBlock {
					internal event internalData;
				}`,
		Err: ErrInvalidType, //event is invalid type for internal data
	},
	{
		Name: "bad type on internal 1",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock {
					internals{
						asdasd internalData;
					} 
				}`,
		Err: ErrInvalidType, //asdasd is invalid type
	},
	// {
	// 	Name: "internal name used twice",
	// 	Input: `basicFB testBlock;
	// 			interface of testBlock {
	// 				in int nameInUse;
	// 			}
	// 			architecture of testBlock {
	// 				internals{
	// 					int nameInUse;
	// 				}
	// 			}`,
	// 	Err: ErrNameAlreadyInUse, //name used twice
	// },
	{
		Name: "missing initial value",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock {
					internals{
						int nameInUse :=;
					} 
				}`,
		Err: ErrUnexpectedValue, //initial should have quote marks
	},
	{
		Name: "bad array on internal",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock {
					internals{
						int[5 nameInUse;
					} 
				}`,
		Err: ErrUnexpectedValue, //bad array
	},
	{
		Name: "missing semicolon on internal",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock {
					internals{
						int[5] nameInUse
					} 
				}`,
		Err: ErrUnexpectedValue, //missing semicolon
	},
	{
		Name: "bad arch 1",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock;`,
		Err: ErrUnexpectedValue, //should be arch of x {} not arch of x ;
	},
	{
		Name: "bad type on internal 1",
		Input: `basicFB testBlock;
				interface of testBlock {
				}
				architecture of testBlock {
					internals{
						asdasd internalData;
					} 
				}`,
		Err: ErrInvalidType, //asdasd is invalid
	},
	{
		Name: "valid basic arch 1",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent;
				}
				architecture of testBlock {
					states {
						s1 {
							emit outEvent;
							
							-> s2 on inEvent;
						}
					}
					
					state s2 {
						-> s1 on inEvent;
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent"}, d).
				AddBFBState("s1", []iec61499.Action{{"", "outEvent", d}}, d).
				AddBFBState("s2", nil, d).
				AddBFBTransition("s1", "s2", "inEvent", d).
				AddBFBTransition("s2", "s1", "inEvent", d),
		},
	},
	{
		Name: "valid basic arch 2",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent1, outEvent2, outEvent3;
				}
				architecture of testBlock {
					states {
						s1 {
							emit outEvent1, outEvent2;
							
							-> s2 on inEvent;
						}
						s2 {
							run testAlg;
							emit outEvent3;

							-> s1 on inEvent;
						}
					}
					algorithms {
						testAlg in "C" ` + "`" + `//c code` + "`" + `;
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent1", "outEvent2", "outEvent3"}, d).
				AddBFBState("s1", []iec61499.Action{{"", "outEvent1", d}, {"", "outEvent2", d}}, d).
				AddBFBState("s2", []iec61499.Action{{"", "outEvent3", d}, {"testAlg", "", d}}, d).
				AddBFBTransition("s1", "s2", "inEvent", d).
				AddBFBTransition("s2", "s1", "inEvent", d).
				AddBFBAlgorithm("testAlg", "C", `//c code`, d),
		},
	},
	{
		Name: "valid basic arch 3",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent1, outEvent2;
				}
				architecture of testBlock {
					internal int internal1;
					internal int internal2;

					state s1 {emit outEvent1, outEvent2; -> s2 on inEvent; }
					state s2 {run testAlg1, testAlg2; -> s1 on inEvent; }

					algorithm testAlg1 in "C" ` + "`" + `//c code` + "`" + `;
					algorithm testAlg2 in "C" ` + "`" + `//more c code` + "`" + `;
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddBFBDataInternals([]string{"internal1", "internal2"}, "int", "", "", d).
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent1", "outEvent2"}, d).
				AddBFBState("s1", []iec61499.Action{{"", "outEvent1", d}, {"", "outEvent2", d}}, d).
				AddBFBState("s2", []iec61499.Action{{"testAlg1", "", d}, {"testAlg2", "", d}}, d).
				AddBFBTransition("s1", "s2", "inEvent", d).
				AddBFBTransition("s2", "s1", "inEvent", d).
				AddBFBAlgorithm("testAlg1", "C", `//c code`, d).
				AddBFBAlgorithm("testAlg2", "C", `//more c code`, d),
		},
	},
	{
		Name: "valid basic arch 4 (anon alg test)",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent1, outEvent2;
				}
				architecture of testBlock {
					internal int internal1;
					internal int internal2;

					state s1 {
						emit outEvent1, outEvent2; 
						
						-> s2 on inEvent; 
					}

					state s2 {
						run in "C" ` + "`" + `//c code` + "`" + `, testAlg2; 
						
						-> s1 on inEvent; 
					}

					algorithm testAlg2 in "C" ` + "`" + `//more c code` + "`" + `;
				}`,
		Output: []iec61499.FB{
			*iec61499.NewBasicFB("testBlock").
				AddBFBDataInternals([]string{"internal1", "internal2"}, "int", "", "", d).
				AddEventInputNames([]string{"inEvent"}, d).
				AddEventOutputNames([]string{"outEvent1", "outEvent2"}, d).
				AddBFBState("s1", []iec61499.Action{{"", "outEvent1", d}, {"", "outEvent2", d}}, d).
				AddBFBState("s2", []iec61499.Action{{"s2_alg0", "", d}, {"testAlg2", "", d}}, d).
				AddBFBTransition("s1", "s2", "inEvent", d).
				AddBFBTransition("s2", "s1", "inEvent", d).
				AddBFBAlgorithm("s2_alg0", "C", `//c code`, d).
				AddBFBAlgorithm("testAlg2", "C", `//more c code`, d),
		},
	},
	{
		Name: "invalid basic arch 1",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent1, outEvent2, outEvent3;
				}
				architecture of testBlock {
					states {
						s1 {
							emit outEvent1, outEvent2;
							
							-> s1 on inEvent;
						}
						s1 {
							run testAlg;
							emit outEvent3;

							-> s1 on inEvent;
						}
					}
					algorithms {
						testAlg in "C" ` + "`" + `//c code` + "`" + `;
					}
				}`,
		Err: ErrNameAlreadyInUse, //s1 is defined twice
	},
	// {
	// 	Name: "state not defined",
	// 	Input: `basicFB testBlock;
	// 			interface of testBlock{
	// 				in event inEvent;
	// 				out event outEvent1, outEvent2, outEvent3;
	// 			}
	// 			architecture of testBlock {
	// 				states {
	// 					s1 {
	// 						emit outEvent1, outEvent2;

	// 						-> s3 on inEvent;
	// 					}
	// 					s2 {
	// 						run testAlg;
	// 						emit outEvent3;

	// 						-> s1 on inEvent;
	// 					}
	// 				}
	// 				algorithms {
	// 					testAlg in "C" ` + "`" + `//c code` + "`" + `;
	// 				}
	// 			}`,
	// 	Err: iec61499.ErrUndefinedState, //s3 isn't defined
	// },
	{
		Name: "algorithm not surrounded by backticks",
		Input: `basicFB testBlock;
				interface of testBlock{
					in event inEvent;
					out event outEvent1, outEvent2, outEvent3;
				}
				architecture of testBlock {
					states {
						s1 {
							emit outEvent1, outEvent2;
							
							-> s1 on inEvent;
						}
						s2 {
							run testAlg;
							emit outEvent3;

							-> s1 on inEvent;
						}
					}
					algorithms {
						testAlg in 'C' //c code;
					}
				}`,
		Err: ErrUnexpectedValue, //alg isn't surrounded by backticks
	},
	// {
	// 	Name: "undefined output event",
	// 	Input: `basicFB testBlock;
	// 			interface of testBlock{
	// 				in event inEvent;
	// 				out event outEvent1, outEvent2, outEvent3;
	// 			}
	// 			architecture of testBlock {
	// 				states {
	// 					s1 {
	// 						emit outEvent1, outEvent4;

	// 						-> s2 on inEvent;
	// 					}
	// 					s2 {
	// 						run testAlg;
	// 						emit outEvent3;

	// 						-> s1 on inEvent;
	// 					}
	// 				}
	// 				algorithms {
	// 					testAlg in "C" ` + "`" + `//c code` + "`" + `;
	// 				}
	// 			}`,
	// 	Err: ErrUndefinedEvent, //outEvent4 doesn't exist
	// },
	// {
	// 	Name: "alg doesn't exist",
	// 	Input: `basicFB testBlock;
	// 			interface of testBlock{
	// 				in event inEvent;
	// 				out event outEvent1, outEvent2, outEvent3;
	// 			}
	// 			architecture of testBlock {
	// 				states {
	// 					s1 {
	// 						emit outEvent1, outEvent2;

	// 						-> s2 on inEvent;
	// 					}
	// 					s2 {
	// 						run testAlg2;
	// 						emit outEvent3;

	// 						-> s1 on inEvent;
	// 					}
	// 				}
	// 				algorithms {
	// 					testAlg in "C" ` + "`" + `//c code` + "`" + `;
	// 				}
	// 			}`,
	// 	Err: iec61499.ErrUndefinedAlgorithm, //testAlg2 doesn't exist
	// },
	{
		Name: "missing brace after s1",
		Input: `basicFB testBlock;
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

func TestParseBFBArchitecture(t *testing.T) {
	runParseTests(t, bfbArchitectureTests)
}
