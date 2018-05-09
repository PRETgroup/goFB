package iec61499

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/PRETgroup/goFB/goFB/stconverter"
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

func TestPoliciesSTGuards(t *testing.T) {
	stGuards, err := aeiFB.Policies[0].GetPFBSTTransitions()
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	for i, g := range stGuards {
		rets := traverse(g.STGuard)
		bytes, _ := json.MarshalIndent(rets, "", "\t")
		ioutil.WriteFile(fmt.Sprintf("test_actual.%v.out.json", i), bytes, 0644)
	}

}

func TestTraverse(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "(a or b) and (c or d)")
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	expr, ok := guard[0].(stconverter.STExpression)
	if !ok {
		t.Fatalf("Got an error and shouldn't have: couldn't cast expression")
	}
	rets := traverse(stconverter.STExpression(expr))
	bytes, _ := json.MarshalIndent(rets, "", "\t")
	ioutil.WriteFile("test_guard.out.json", bytes, 0644)
}
