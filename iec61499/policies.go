package iec61499

import (
	"github.com/PRETgroup/goFB/goFB/stconverter"
)

//FBECCGuardToSTExpression converts a given FB's guard into a STExpression parsetree
func FBECCGuardToSTExpression(fb *FB, guard string) ([]stconverter.STInstruction, *stconverter.STParseError) {
	return stconverter.ParseString(fb.Name, guard)
	// if err != nil {
	// 	return nil, err
	// }

}

//SeparateViolationTransitions will scan a PolicyFB for violation transitions where the guard has multiple clauses separated by ORs
//Where found, it will break them into two transitions
func (p *PolicyFB) SeparateViolationTransitions() {

}
