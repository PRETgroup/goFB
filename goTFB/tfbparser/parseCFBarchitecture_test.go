package tfbparser

import (
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
)

var cfbArchitectureTests = []ParseTest{
	{
		Name: "basic arch 1, single instance",
		Input: `compositeFB testBlock;
				interface of testBlock {
					in event inEvent;
					in with inEvent lreal inData;
					out event outEvent;
					out with outEvent lreal outData;
				}
				architecture of testBlock {
					instance Incr incr;
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewCompositeFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddCFBInstances("Incr", []string{"incr"}, d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 2, two joined instances",
		Input: `compositeFB testBlock;
				interface of testBlock {
					in event inEvent;
					in with inEvent lreal inData;
					out event outEvent;
					out with outEvent lreal outData;
				}
				architecture of testBlock {
					instance Incr incr1, incr2;
					
					events {
						incr1.inEvent <- incr2.outEvent;
					}

					data {
						incr1.inData <- incr2.outData;
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewCompositeFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddCFBInstances("Incr", []string{"incr1", "incr2"}, d).
				AddCFBNetworkEventConns([]string{"incr2.outEvent"}, "incr1.inEvent", d).
				AddCFBNetworkDataConn("incr2.outData", "incr1.inData", d),
		},
		Err: nil,
	},
	{
		Name: "basic arch 3, two joined instances with param and double event source",
		Input: `compositeFB testBlock;
				interface of testBlock {
					in event inEvent;
					in with inEvent lreal inData;
					out event outEvent;
					out with outEvent lreal outData;
				}
				architecture of testBlock {
					instance Incr incr1, incr2;
					
					events {
						incr1.inEvent <- incr2.outEvent, incr1.outEvent;
					}

					data {
						incr1.inData <- incr2.outData;
						incr2.inData <- ` + "`" + `1.0` + "`" + `;
					}
				}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewCompositeFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)).
				AddCFBInstances("Incr", []string{"incr1", "incr2"}, d).
				AddCFBNetworkEventConns([]string{"incr2.outEvent", "incr1.outEvent"}, "incr1.inEvent", d).
				AddCFBNetworkDataConn("incr2.outData", "incr1.inData", d).
				AddCFBNetworkParameter("1.0", "incr2", "inData", d),
		},
		Err: nil,
	},
}

func TestParseCFBArchitecture(t *testing.T) {
	runParseTests(t, cfbArchitectureTests)
}
