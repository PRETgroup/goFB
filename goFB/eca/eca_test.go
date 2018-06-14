package eca

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
	"github.com/PRETgroup/goFB/iec61499/fbexamples"
)

var trainCtrlChains = map[string][]EventTrace{
	"abort": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "n_allow_train_exit_0", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "n_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "s_allow_train_exit_0", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "s_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_s", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_n", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_n_passed_first_signal", Destination: "idle", Condition: "abort"}, OutputEvents: []string{
					"SChange"}}}},
	"RsChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "idle", Destination: "s_allow_train_exit_0", Condition: "RsChange && RsReq == true"}, OutputEvents: []string{
					"WChange", "SChange"}}}},
	"DwoChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "n_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "DwoChange && DwoPrs == false"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "s_allow_train_exit_passed_first_signal", Destination: "idle", Condition: "DwoChange && DwoPrs == false"}, OutputEvents: []string{
					"SChange"}}}},
	"DwnChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "n_allow_train_exit_0", Destination: "n_allow_train_exit_passed_first_signal", Condition: "DwnChange && DwnPrs == false"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_n_passed_first_signal", Destination: "idle", Condition: "DwnChange && DwnPrs == false"}, OutputEvents: []string{
					"SChange"}}}},
	"DwiChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_s", Destination: "i_allow_train_entrance_s_passed_first_signal", Condition: "DwiChange && DwiPrs == false"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_n", Destination: "i_allow_train_entrance_n_passed_first_signal", Condition: "DwiChange && DwiPrs == false"}, OutputEvents: []string{
					"SChange"}}}},
	"DwsChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "s_allow_train_exit_0", Destination: "s_allow_train_exit_passed_first_signal", Condition: "DwsChange && DwsPrs == false"}, OutputEvents: []string{
					"SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_s_passed_first_signal", Destination: "idle", Condition: "DwsChange && DwsPrs == false"}, OutputEvents: []string{
					"SChange"}}}},
	"DnChange": []EventTrace{},
	"DsChange": []EventTrace{},
	"SysReady": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "init", Destination: "idle", Condition: "SysReady"}, OutputEvents: []string{
					"SChange"}}}},
	"RiChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)},
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_0", Destination: "i_allow_train_entrance_s", Condition: "busyS == false && DsPrs == false"}, OutputEvents: []string{
					"WChange", "SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)},
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_0", Destination: "i_allow_train_entrance_n", Condition: "busyN == false && DnPrs == false"}, OutputEvents: []string{
					"WChange", "SChange"}}},
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "idle", Destination: "i_allow_train_entrance_0", Condition: "RiChange && RiReq == true"}, OutputEvents: []string(nil)},
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "i_allow_train_entrance_0", Destination: "idle", Condition: "true"}, OutputEvents: []string{
					"SChange"}}}},
	"RnChange": []EventTrace{
		EventTrace{
			EventTraceStep{
				InboundTransition: iec61499.ECTransition{
					Source: "idle", Destination: "n_allow_train_exit_0", Condition: "RnChange && RnReq == true"}, OutputEvents: []string{
					"WChange", "SChange"}}}}}

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

	if !reflect.DeepEqual(analysedChains, trainCtrlChains) {
		t.Fatalf("The chains don't match!")
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
