package eca

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
	"github.com/PRETgroup/goFB/iec61499/fbexamples"
)

var trainCtrlChains = []EventChain{
	EventChain{InputName: "SysReady",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "init", Destination: "idle", Condition: "SysReady"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "RiChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)}, EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_0", Destination: "i_allow_train_entrance_s", Condition: "busyS == false && DsPrs == false"}, OutputEvents: []string{"WChange", "SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)}, EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_0", Destination: "i_allow_train_entrance_n", Condition: "busyN == false && DnPrs == false"}, OutputEvents: []string{"WChange", "SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)}, EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_0", Destination: "idle", Condition: "true"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "RnChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "idle", Destination: "n_allow_train_exit_0", Condition: "RnChange && RnReq == true"}, OutputEvents: []string{"WChange", "SChange"}}}}},
	EventChain{InputName: "RsChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "idle", Destination: "s_allow_train_exit_0", Condition: "RsChange && RsReq == true"}, OutputEvents: []string{"WChange", "SChange"}}}}},
	EventChain{InputName: "DwiChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_s", Destination: "i_allow_train_entrance_s_passed_first_signal", Condition: "DwiChange && DwiPrs == false"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_n", Destination: "i_allow_train_entrance_n_passed_first_signal", Condition: "DwiChange && DwiPrs == false"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "DwoChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "n_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "DwoChange && DwoPrs == false"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "s_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "DwoChange && DwoPrs == false"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "DwnChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "n_allow_train_exit_0", Destination: "n_allow_train_exit_passed_first_signal", Condition: "DwnChange && DwnPrs == false"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_n_passed_first_signal", Destination: "idle", Condition: "DwnChange && DwnPrs == false"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "DwsChange",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "s_allow_train_exit_0", Destination: "s_allow_train_exit_passed_first_signal", Condition: "DwsChange && DwsPrs == false"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_s_passed_first_signal", Destination: "idle", Condition: "DwsChange && DwsPrs == false"}, OutputEvents: []string{"SChange"}}}}},
	EventChain{InputName: "DnChange",
		OutputTraces: []EventTrace(nil)},
	EventChain{InputName: "DsChange",
		OutputTraces: []EventTrace(nil)},
	EventChain{InputName: "abort",
		OutputTraces: []EventTrace{
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "n_allow_train_exit_0", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "n_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "s_allow_train_exit_0", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "s_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_s", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_n", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}},
			EventTrace{EventTraceStep{InboundTransition: iec61499.ECTransition{Source: "i_allow_train_entrance_n_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{"SChange"}}}}}}

var trainCtrlEventSources = []InstanceConnection{
	InstanceConnection{InstanceID: 3, PortName: "RiChange"},
	InstanceConnection{InstanceID: 3, PortName: "RnChange"},
	InstanceConnection{InstanceID: 3, PortName: "RsChange"},
	InstanceConnection{InstanceID: 3, PortName: "DwiChange"},
	InstanceConnection{InstanceID: 3, PortName: "DwoChange"},
	InstanceConnection{InstanceID: 3, PortName: "DwnChange"},
	InstanceConnection{InstanceID: 3, PortName: "DwsChange"},
	InstanceConnection{InstanceID: 3, PortName: "DnChange"},
	InstanceConnection{InstanceID: 3, PortName: "DsChange"},
	InstanceConnection{InstanceID: 3, PortName: "SoChange"},
	InstanceConnection{InstanceID: 3, PortName: "SysReady"},
}

func TestDeriveBFBEventChainSet(t *testing.T) {
	trainStationFBs := make([]iec61499.FB, 0)

	for _, FBT := range fbexamples.EventTrainStationFBT {
		fb := iec61499.FB{}
		if err := xml.Unmarshal([]byte(FBT), &fb); err != nil {
			t.Fatal("Couldn't unmarshal test TrainCtrl XML:", err.Error())
		}
		trainStationFBs = append(trainStationFBs, fb)
	}

	analysedChains, err := DeriveBFBEventChainSet(trainStationFBs[6])
	if err != nil {
		t.Fatal("Error when deriving chains (there should have been no error):", err.Error())
	}

	if len(trainCtrlChains) != len(analysedChains) {
		t.Fatalf("Provided chains len != analysed chain len, %d != %d", len(trainCtrlChains), len(analysedChains))
	}

	//find each analysed chain inside provided chains
	for i := 0; i < len(analysedChains); i++ {
		found := false
		for j := 0; j < len(trainCtrlChains); j++ {
			if analysedChains[i].InputName == trainCtrlChains[j].InputName {
				found = true
				//make sure the names have the same number of traces
				if len(analysedChains[i].OutputTraces) != len(trainCtrlChains[j].OutputTraces) {
					t.Fatalf("Trace for name %s length mismatch #1 (%d != %d)", analysedChains[i].InputName, len(analysedChains[i].OutputTraces), len(trainCtrlChains[j].OutputTraces))
				}

				//make sure each trace is identical
				for k := 0; k < len(analysedChains[i].OutputTraces); k++ {
					if len(analysedChains[i].OutputTraces[k]) != len(trainCtrlChains[j].OutputTraces[k]) {
						t.Fatalf("Trace %s.%d has length mismatch #2 (%d != %d)", analysedChains[i].InputName, k, len(analysedChains[i].OutputTraces[k]), len(trainCtrlChains[j].OutputTraces[k]))
					}

					for l := 0; l < len(analysedChains[i].OutputTraces[k]); l++ {
						if analysedChains[i].OutputTraces[k][l].GetECStateName() != trainCtrlChains[j].OutputTraces[k][l].GetECStateName() {
							t.Fatalf("Trace %s.%d has data mismatch #3 (%#v != %#v)", analysedChains[i].InputName, k, analysedChains[i].OutputTraces[k], trainCtrlChains[j].OutputTraces[k])
						}

						if !reflect.DeepEqual(analysedChains[i].OutputTraces[k][l].OutputEvents, trainCtrlChains[j].OutputTraces[k][l].OutputEvents) {
							t.Fatalf("Trace %s.%d has data mismatch #4 (%#v != %#v)", analysedChains[i].InputName, k, analysedChains[i].OutputTraces[k], trainCtrlChains[j].OutputTraces[k])
						}
					}
				}

				break
			}
		}
		if found == false {
			t.Fatalf("Couldn't find chain for name %s", analysedChains[i].InputName)
		}
	}
}

func TestDeriveBFBEventChainSet_SelfLoopError(t *testing.T) {
	trainStationFBs := make([]iec61499.FB, 0)

	for _, FBT := range fbexamples.EventTrainStationFBT {
		fb := iec61499.FB{}
		if err := xml.Unmarshal([]byte(FBT), &fb); err != nil {
			t.Fatal("Couldn't unmarshal test TrainCtrl XML:", err.Error())
		}
		trainStationFBs = append(trainStationFBs, fb)
	}

	trainCtrlFB := trainStationFBs[6]

	//add a bad (infinite loop) transition
	trainCtrlFB.BasicFB.Transitions = append(trainCtrlFB.BasicFB.Transitions, iec61499.ECTransition{
		Source:      "i_allow_train_entrance_s",
		Destination: "i_allow_train_entrance_0",
		Condition:   "true",
	})

	_, err := DeriveBFBEventChainSet(trainCtrlFB)
	if err == nil {
		t.Fatal("Error: an instantaneous self-loop should have been detected")
	}
}

func TestListSIFBEventSources(t *testing.T) {
	trainStationFBs := make([]iec61499.FB, 0)

	for _, FBT := range fbexamples.EventTrainStationFBT {
		fb := iec61499.FB{}
		if err := xml.Unmarshal([]byte(FBT), &fb); err != nil {
			t.Fatal("Couldn't unmarshal test TrainCtrl XML:", err.Error())
		}
		trainStationFBs = append(trainStationFBs, fb)
	}

	trainInstG, err := CreateInstanceGraph(trainStationFBs, "Top")
	if err != nil {
		t.Fatal("There was an error, and there shouldn't have been (#1)")
	}

	eventSources, err := ListSIFBEventSources(trainInstG, trainStationFBs)
	if err != nil {
		t.Fatal("There was an error, and there shouldn't have been (#2)")
	}

	if !reflect.DeepEqual(eventSources, trainCtrlEventSources) {
		t.Fatal("The derived event sources for trainCtrl did not match what was expected.")
	}
}
