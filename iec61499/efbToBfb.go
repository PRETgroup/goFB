package iec61499

import "errors"

//TranslateEFBtoBFB will take an Enforcer Function Block and compile it to its equivalent
// Basic Function Block
//It operates according to the algorithm specified in [TODO: Paper link]
func (f *FB) TranslateEFBtoBFB() error {
	if f.EnforceFB == nil {
		return errors.New("TranslateEFBtoBFB can only be called on an EnforceFB")
	}

	return errors.New("Not yet implemented")
}
