package iec61499

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

var aeiFB = *Must(
	NewBasicFB("AEIPolicy").
		AddEventInputNames([]string{"AS", "VS"}, d).
		AddEventOutputNames([]string{"AP", "VP"}, d).
		AddDataInputs([]string{"AEI_ns"}, []string{}, "ulint", "", "900000000", d)).
	AddPolicy("AEI").
	AddPFBState("s1", d).
	AddPFBState("s2", d).
	AddPFBDataInternals([]string{"tAEI"}, "DTIMER", "", "", d).
	AddPFBTransition("s1", "s2", "( VS or VP )", []PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
	AddPFBTransition("s2", "s1", "( AS or AP )", nil, d).
	AddPFBTransition("s2", "violation", "( tAEI > AEI_ns )", nil, d)

var aeiFBSplit = *Must(
	NewBasicFB("AEIPolicy").
		AddEventInputNames([]string{"AS", "VS"}, d).
		AddEventOutputNames([]string{"AP", "VP"}, d).
		AddDataInputs([]string{"AEI_ns"}, []string{}, "ulint", "", "900000000", d)).
	AddPolicy("AEI").
	AddPFBState("s1", d).
	AddPFBState("s2", d).
	AddPFBDataInternals([]string{"tAEI"}, "DTIMER", "", "", d).
	AddPFBTransition("s1", "s2", "VS", []PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
	AddPFBTransition("s1", "s2", "VP", []PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
	AddPFBTransition("s2", "s1", "AS", nil, d).
	AddPFBTransition("s2", "s1", "AP", nil, d).
	AddPFBTransition("s2", "violation", "( tAEI > AEI_ns )", nil, d)

var ab5 = *NewBasicFB("AB5Policy").
	AddEventInputNames([]string{"A"}, d).
	AddEventOutputNames([]string{"B"}, d).
	AddPolicy("AB5").
	AddPFBState("s0", d).
	AddPFBState("s1", d).
	AddPFBDataInternals([]string{"v"}, "DTIMER", "", "", d).
	AddPFBTransition("s0", "s0", "( !A and !B )", []PFBExpression{{VarName: "v", Value: "0"}}, d).
	AddPFBTransition("s0", "s1", "( A and !B )", []PFBExpression{{VarName: "v", Value: "0"}}, d).
	AddPFBTransition("s0", "violation", "( ( !A and B ) or ( A and B ) )", nil, d).
	AddPFBTransition("s1", "s1", "( !A and !B and v < 5 )", nil, d).
	AddPFBTransition("s1", "s0", "( !A and B )", nil, d).
	AddPFBTransition("s1", "violation", "( ( v >= 5 ) or ( A and B ) or ( A and !B ) )", nil, d)

func TestDeriveInputAEIPolicy(t *testing.T) {
	enfPol, err := MakePFBEnforcer(ab5.InterfaceList, ab5.Policies[0])
	if err != nil {
		t.Fatalf("Error occured and shouldn't have:%v\n", err.Error())
	}
	s, _ := json.MarshalIndent(enfPol, "", "\t")
	ioutil.WriteFile(fmt.Sprintf("test_actual.out.json"), s, 0644)

}

// func TestPoliciesSTGuards(t *testing.T) {

// }
