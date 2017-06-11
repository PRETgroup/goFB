package tfbparser

import (
	"testing"

	"github.com/kiwih/goFB/iec61499"
)

var interfaceTests = []ParseTest{
	{
		Name: "simple events 1",
		Input: `basicFB testBlock;
				interface of testBlock {
					in event inEvent;
					out event outEvent;
				}`,
		Output: []iec61499.FB{*iec61499.NewBasicFB("testBlock").
			AddEventInputNames([]string{"inEvent"}, d).
			AddEventOutputNames([]string{"outEvent"}, d)},
		Err: nil,
	},
	{
		Name: "simple events 2",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						out event outEvent;
					}`,
		Output: []iec61499.FB{*iec61499.NewBasicFB("testBlock").
			AddEventInputNames([]string{"inEvent"}, d).
			AddEventOutputNames([]string{"outEvent"}, d)},
		Err: nil,
	},
	{
		Name: "events typo 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						out outEvent;
					}`,
		Err: ErrInvalidType,
	},
	{
		Name: "events typo 2",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						out event outEvent
					}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "data typo 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in asdasd inData with inEvent;
						out event outEvent;
					}`,
		Err: ErrInvalidType,
	},
	{
		Name: "data input 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in bool inData with inEvent;
						out event outEvent;
					}`,
		Output: []iec61499.FB{*iec61499.Must(iec61499.NewBasicFB("testBlock").AddEventInputNames([]string{"inEvent"}, d).AddEventOutputNames([]string{"outEvent"}, d).AddDataInputs([]string{"inData"}, []string{"inEvent"}, "bool", "", "", d))},
		Err:    nil,
	},
	{
		Name: "data input 2",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in bool[3] inData with inEvent;
						out event outEvent;
					}`,
		Output: []iec61499.FB{*iec61499.Must(iec61499.NewBasicFB("testBlock").AddEventInputNames([]string{"inEvent"}, d).AddEventOutputNames([]string{"outEvent"}, d).AddDataInputs([]string{"inData"}, []string{"inEvent"}, "bool", "3", "", d))},
		Err:    nil,
	},
	{
		Name: "data input array typo 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in bool[3 inData with inEvent;
						out event outEvent;
					}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "data input 3",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in bool[3] initial [0,1,0] inData with inEvent; 
						out event outEvent;
					}`,
		Output: []iec61499.FB{*iec61499.Must(iec61499.NewBasicFB("testBlock").AddEventInputNames([]string{"inEvent"}, d).AddEventOutputNames([]string{"outEvent"}, d).AddDataInputs([]string{"inData"}, []string{"inEvent"}, "bool", "3", "[0,1,0]", d))},
		Err:    nil,
	},
	{
		Name: "data default typo 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in bool[3] initial 0,1,0 inData with inEvent;
						out event outEvent;
					}`,
		Err: ErrUnexpectedValue,
	},
	{
		Name: "data default in/out 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in lreal inData with inEvent;
						out event outEvent;
						out lreal outData with outEvent;
					}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewBasicFB("testBlock").
						AddEventInputNames([]string{"inEvent"}, d).
						AddEventOutputNames([]string{"outEvent"}, d).
						AddDataInputs([]string{"inData"}, []string{"inEvent"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData"}, []string{"outEvent"}, "lreal", "", "", d)),
		},
		Err: nil,
	},
	{
		Name: "data default in/out 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent1, inEvent2;
						in lreal inData1, inData2 with inEvent1, inEvent2;
						out event outEvent1, outEvent2;
						out lreal outData1, outData2 with outEvent1, outEvent2;
					}`,
		Output: []iec61499.FB{
			*iec61499.Must(
				iec61499.Must(
					iec61499.NewBasicFB("testBlock").
						AddEventInputNames([]string{"inEvent1", "inEvent2"}, d).
						AddEventOutputNames([]string{"outEvent1", "outEvent2"}, d).
						AddDataInputs([]string{"inData1", "inData2"}, []string{"inEvent1", "inEvent2"}, "lreal", "", "", d)).
					AddDataOutputs([]string{"outData1", "outData2"}, []string{"outEvent1", "outEvent2"}, "lreal", "", "", d)),
		},
		Err: nil,
	},
	{
		Name: "data bad triggers 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in lreal inData with inEvent;
						in event inEvent;
						out event outEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrUndefinedEvent,
	},
	{
		Name: "data bad triggers 2",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in lreal inData with inEvent;
						out event outEvent;
						out lreal outData with inEvent;
					}`,
		Err: ErrUndefinedEvent,
	},
	// {
	// 	Name: "data bad triggers 3",
	// 	Input: `compositeFB testBlock;
	// 				interface of testBlock {
	// 					in event inEvent;
	// 					in lreal inData with inEvent;
	// 					out event outEvent;
	// 					out lreal outData with outEvent;
	// 				}`,
	// 	Err: iec61499.ErrOnlyBasicFBsGetTriggers,
	// },
	{
		Name: "data bad triggers 4",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in lreal inData with asdasdasdasd;
						out event outEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrUndefinedEvent,
	},
	{
		Name: "data bad triggers 5",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in lreal inData with inEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrUndefinedEvent,
	},
	{
		Name: "Events don't get triggers 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in event inEvent2 with inEvent;
						out event outEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrUnexpectedAssociation,
	},
	{
		Name: "Events don't get arrays 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in event[3] inEvent2;
						out event outEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrInvalidIOMeta,
	},
	{
		Name: "Events don't get initials 1",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in event inEvent2;
						out event initial 1 outEvent;
						out lreal outData with outEvent;
					}`,
		Err: ErrInvalidIOMeta,
	},
	// {
	// 	Name: "Name in use 1",
	// 	Input: `basicFB testBlock;
	// 				interface of testBlock {
	// 					in event inEvent;
	// 					in event inEvent;
	// 					out event outEvent;
	// 					out lreal outData with outEvent;
	// 				}`,
	// 	Err: ErrNameAlreadyInUse,
	// },
	{
		Name: "Unexpected EOF",
		Input: `basicFB testBlock;
					interface of testBlock {
						in event inEvent;
						in event inEvent2;
						out int initial`,
		Err: ErrUnexpectedValue,
	},
}

func TestParseStringInterface(t *testing.T) {
	runParseTests(t, interfaceTests)
}
