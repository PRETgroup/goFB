package eca

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/PRETgroup/goFB/iec61499"
)

const trainCtrlFBT = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="TrainCtrl" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs>
			<Event Name="SysReady" Comment=""></Event>
			<Event Name="RiChange" Comment="">
				<With Var="RiReq"></With>
			</Event>
			<Event Name="RnChange" Comment="">
				<With Var="RnReq"></With>
			</Event>
			<Event Name="RsChange" Comment="">
				<With Var="RsReq"></With>
			</Event>
			<Event Name="DwiChange" Comment="">
				<With Var="DwiPrs"></With>
			</Event>
			<Event Name="DwoChange" Comment="">
				<With Var="DwoPrs"></With>
			</Event>
			<Event Name="DwnChange" Comment="">
				<With Var="DwnPrs"></With>
			</Event>
			<Event Name="DwsChange" Comment="">
				<With Var="DwsPrs"></With>
			</Event>
			<Event Name="DnChange" Comment="">
				<With Var="DnPrs"></With>
			</Event>
			<Event Name="DsChange" Comment="">
				<With Var="DsPrs"></With>
			</Event>
			<Event Name="abort" Comment=""></Event>
		</EventInputs>
		<EventOutputs>
			<Event Name="SChange" Comment="">
				<With Var="SiGrn"></With>
				<With Var="SnGrn"></With>
				<With Var="SsGrn"></With>
			</Event>
			<Event Name="WChange" Comment="">
				<With Var="WiDvrg"></With>
				<With Var="WoDvrg"></With>
				<With Var="WnDvrg"></With>
				<With Var="WsDvrg"></With>
			</Event>
		</EventOutputs>
		<InputVars>
			<VarDeclaration Name="RiReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RnReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RsReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InputVars>
		<OutputVars>
			<VarDeclaration Name="SiGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</OutputVars>
	</InterfaceList>
	<BasicFB>
		<InternalVars>
			<VarDeclaration Name="busyN" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="busyS" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InternalVars>
		<ECC>
			<ECState Name="init" Comment="" x="" y=""></ECState>
			<ECState Name="idle" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="ClrSignals"></ECAction>
			</ECState>
			<ECState Name="n_allow_train_exit_0" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNExit"></ECAction>
			</ECState>
			<ECState Name="n_allow_train_exit_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNExitHalf"></ECAction>
			</ECState>
			<ECState Name="s_allow_train_exit_0" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSExit"></ECAction>
			</ECState>
			<ECState Name="s_allow_train_exit_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSExitHalf"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_0" Comment="" x="" y=""></ECState>
			<ECState Name="i_allow_train_entrance_s" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSEntrance"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_s_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetEntranceHalf"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_n" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNEntrance"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_n_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetEntranceHalf"></ECAction>
			</ECState>
			<ECTransition Source="init" Destination="idle" Condition="SysReady" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="n_allow_train_exit_0" Condition="RnChange &amp;&amp; RnReq == true" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="s_allow_train_exit_0" Condition="RsChange &amp;&amp; RsReq == true" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="i_allow_train_entrance_0" Condition="RiChange &amp;&amp; RiReq == true" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_0" Destination="n_allow_train_exit_passed_first_signal" Condition="DwnChange &amp;&amp; DwnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_0" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_passed_first_signal" Destination="idle" Condition="DwoChange &amp;&amp; DwoPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_0" Destination="s_allow_train_exit_passed_first_signal" Condition="DwsChange &amp;&amp; DwsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_0" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_passed_first_signal" Destination="idle" Condition="DwoChange &amp;&amp; DwoPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="i_allow_train_entrance_s" Condition="busyS == false &amp;&amp; DsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="i_allow_train_entrance_n" Condition="busyN == false &amp;&amp; DnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="idle" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s" Destination="i_allow_train_entrance_s_passed_first_signal" Condition="DwiChange &amp;&amp; DwiPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s_passed_first_signal" Destination="idle" Condition="DwsChange &amp;&amp; DwsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n" Destination="i_allow_train_entrance_n_passed_first_signal" Condition="DwiChange &amp;&amp; DwiPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n_passed_first_signal" Destination="idle" Condition="DwnChange &amp;&amp; DwnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
		</ECC>
		<Algorithm Name="ClrSignals" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: ClrSignals\n&#34;);&#xA;        me-&gt;SiGrn = false;&#xA;        me-&gt;SnGrn = false;&#xA;        me-&gt;SsGrn = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNEntrance" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNEntrance\r\n&#34;);&#xA;        me-&gt;SiGrn = true;&#xA;        me-&gt;WiDvrg = false;&#xA;        me-&gt;WnDvrg = false;&#xA;        me-&gt;busyN = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSEntrance" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSEntrance\r\n&#34;);&#xA;        me-&gt;SiGrn = true;&#xA;        me-&gt;WiDvrg = true;&#xA;        me-&gt;WsDvrg = false; //this is an error!&#xA;        me-&gt;busyS = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetEntranceHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetEntranceHalf\n&#34;);&#xA;        me-&gt;SiGrn = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNExit" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNExit\n&#34;);&#xA;        me-&gt;SnGrn = true;&#xA;        me-&gt;WnDvrg = true;&#xA;        me-&gt;WoDvrg = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNExitHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNExitHalf\n&#34;);&#xA;        me-&gt;SnGrn = false;&#xA;        me-&gt;busyN = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSExit" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSExit\n&#34;);&#xA;        me-&gt;SsGrn = true;&#xA;        me-&gt;WsDvrg = false;&#xA;        me-&gt;WoDvrg = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSExitHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSExitHalf\n&#34;);&#xA;        me-&gt;SsGrn = false;&#xA;        me-&gt;busyS = false;&#xA;    "></Other>
		</Algorithm>
	</BasicFB>
</FBType>`

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

func TestDeriveBFBEventChainSet(t *testing.T) {
	trainCtrlFB := iec61499.FB{}
	if err := xml.Unmarshal([]byte(trainCtrlFBT), &trainCtrlFB); err != nil {
		t.Fatal("Couldn't unmarshal test TrainCtrl XML:", err.Error())
	}

	analysedChains, err := DeriveBFBEventChainSet(trainCtrlFB)
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
	trainCtrlFB := iec61499.FB{}
	if err := xml.Unmarshal([]byte(trainCtrlFBT), &trainCtrlFB); err != nil {
		t.Fatal("Couldn't unmarshal test TrainCtrl XML:", err.Error())
	}

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
