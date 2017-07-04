package iec61499

import "errors"

//TranslateHFBtoBFB will take a Hybrid Function Block and translate it to its equivalent
// Basic Function Block
//It operates according to the algorithm specified in [TODO: Paper link]
func (f *FB) TranslateHFBtoBFB() error {
	if f.HybridFB == nil {
		return errors.New("TranslateHFBtoBFB can only be called on a HybridFB")
	}

	return errors.New("Not yet implemented")
}
