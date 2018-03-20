package stconverter

//STInstruction is the container interface for the sequence instructions
//we use it to ensure that types passed into STInstruction are types that we can understand
//(just as a sanity check)
type STInstruction interface {
	IsInstruction() bool
	// GetAssignment() *STAssignment
	// GetIfElsIfElse() *STIfElsIfElse
	// GetSwitchCase() *STSwitchCase
	// GetForLoop() *STForLoop
	// GetWhileLoop() *STWhileLoop
	// GetRepeatLoop() *STRepeatLoop
}

//IsInstruction meets the STInstruction interface
func (s STAssignment) IsInstruction() bool {
	return true
}

//IsInstruction meets the STInstruction interface
func (s STIfElsIfElse) IsInstruction() bool {
	return true
}

//IsInstruction meets the STInstruction interface
func (s STSwitchCase) IsInstruction() bool {
	return true
}

//IsInstruction meets the STInstruction interface
func (s STForLoop) IsInstruction() bool {
	return true
}

//IsInstruction meets the STInstruction interface
func (s STWhileLoop) IsInstruction() bool {
	return true
}

//IsInstruction meets the STInstruction interface
func (s STRepeatLoop) IsInstruction() bool {
	return true
}

// func (s *STAssignment) GetAssignment() *STAssignment {
// 	return s
// }

// func (s *STAssignment) GetIfElsIfElse() *STIfElsIfElse {
// 	return nil
// }

// func (s *STAssignment) GetSwitchCase() *STSwitchCase {
// 	return nil
// }

// func (s *STAssignment) GetForLoop() *STForLoop {
// 	return nil
// }

// func (s *STAssignment) GetWhileLoop() *STWhileLoop {
// 	return nil
// }
