package iec61499

import "errors"

//TranslatePFBtoBFB will take an Policyr Function Block and compile it to its equivalent
// Basic Function Block
//It operates according to the algorithm specified in [TODO: Paper link]
func (f *FB) TranslatePFBtoBFB() error {
	if f.PolicyFB == nil {
		return errors.New("TranslatePFBtoBFB can only be called on an PolicyFB")
	}

	return errors.New("Not yet implemented")
}
