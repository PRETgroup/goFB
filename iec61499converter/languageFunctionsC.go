package iec61499converter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kiwih/goFB/iec61499converter/iec61499"
)

//CECCTransition is used with getCECCTransitionCondition to return results to the template
type CECCTransition struct {
	IfCond    string
	AssEvents []string
}

//getCECCTransitionCondition returns the C "if" condition to use in state machine next state logic and associated events
// returns "full condition", "associated events"
func getCECCTransitionCondition(block iec61499.FB, iec61499trans string) CECCTransition {
	var events []string

	re1 := regexp.MustCompile("([<>=!]+)")          //for capturing operators
	re2 := regexp.MustCompile("([a-zA-Z0-9_<>=]+)") //for capturing variable and event names and operators
	isNum := regexp.MustCompile("^[0-9.]+$")

	retVal := iec61499trans

	//rename AND and OR
	retVal = strings.Replace(retVal, "AND", "&&", -1)
	retVal = strings.Replace(retVal, "OR", "||", -1)

	//re1: add whitespace around operators
	retVal = re1.ReplaceAllStringFunc(retVal, func(in string) string {
		return " " + in + " "
	})

	//re2: add "me->" where appropriate
	retVal = re2.ReplaceAllStringFunc(retVal, func(in string) string {
		if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ContainsAny(in, "!><=") || strings.ToLower(in) == "true" || strings.ToLower(in) == "false" {
			//no need to make changes, these aren't variables or events
			return in
		}

		if isNum.MatchString(in) {
			//no need to make changes, it is a numerical value of some sort
			return in
		}

		//check to see if it is an input event
		if block.EventInputs != nil {
			for _, event := range block.EventInputs.Events {
				if event.Name == in {
					events = append(events, "me->inputEvents.event."+event.Name)
					return "me->inputEvents.event." + event.Name //FUTURE WORK: Consider use of pointers and offsets to minimise memory footprint for events? i.e. "*(me->inputEvents." + in + " + ev_offset)"
				}
			}
		}

		//check to see if it is an output event
		if block.EventOutputs != nil {
			for _, event := range block.EventOutputs.Events {
				if event.Name == in {
					return "me->outputEvents.event." + event.Name
				}
			}
		}

		//check to see if it is input data
		if block.InputVars != nil {
			for _, Var := range block.InputVars.Variables {
				if Var.Name == in {
					return "me->" + in
				}
			}
		}

		//check to see if it is output data
		if block.OutputVars != nil {
			for _, Var := range block.OutputVars.Variables {
				if Var.Name == in {
					return "me->" + in
				}
			}
		}

		//check to see if it is internal var
		if block.BasicFB != nil && block.BasicFB.InternalVars != nil {
			for _, Var := range block.BasicFB.InternalVars.Variables {
				if Var.Name == in {
					return "me->" + in
				}
			}
		}

		//else, return it (no idea what else to do!) - it might be a function call or strange text constant
		return in
	})

	//tidy the whitespace
	retVal = strings.Replace(retVal, "  ", " ", -1)

	return CECCTransition{IfCond: retVal, AssEvents: events}
}

//locationType = 0 (Var)
//locationType = 1 (destination)
//locationType = 2 (source)
//don't use this function, use one of the helper functions
func renameCLocation(in string, locationType int) string {
	if strings.Contains(in, ".") {
		//it comes from a child FB
		if locationType == 1 {
			return strings.Replace(in, ".", ".inputEvents.event.", 1) //events are in a sub struct
		} else if locationType == 2 {
			return strings.Replace(in, ".", ".outputEvents.event.", 1) //events are in a sub struct
		}

		return in

	}
	//it comes from the parent FB, which means input/output is swapped
	if locationType == 1 {
		return "outputEvents.event." + in
	}
	if locationType == 2 {
		return "inputEvents.event." + in
	}
	return in

}

func renameCEventDestinationLocation(in string) string {
	return renameCLocation(in, 1)
}

func renameCEventSourceLocation(in string) string {
	return renameCLocation(in, 2)
}

//This finds a data connection source based on a destination source (destinations can only ever have one source for data connections in iec61499)
func findSourceDataName(conns []iec61499.Connection, destChildName string, destVarName string) string {
	for _, conn := range conns {
		if destChildName != "" {
			if conn.Destination == destChildName+"."+destVarName {
				return renameCLocation(conn.Source, 0)
			}
		} else {
			if conn.Destination == destVarName {
				return renameCLocation(conn.Source, 0)
			}
		}
	}
	return ""
}

//Ths finds event connection source(s) based on a destination source (events can have multiple sources for event connections in iec61499)
func findSourcesEventName(conns []iec61499.Connection, destChildName string, destEventName string) []string {
	var sources []string
	for _, conn := range conns {
		if destChildName != "" {
			if conn.Destination == destChildName+"."+destEventName {
				sources = append(sources, renameCLocation(conn.Source, 2))
			}
		} else {
			if conn.Destination == destEventName {
				sources = append(sources, renameCLocation(conn.Source, 2))
			}
		}
	}
	return sources
}

//findDestsEventName will find the destination(s) of connections given a connection source (one source can go to many destinations)
func findDestsEventName(conns []iec61499.Connection, sourceChildName string, sourceEventName string) []string {
	var dests []string
	for _, conn := range conns {
		if sourceChildName != "" {
			if conn.Source == sourceChildName+"."+sourceEventName {
				dests = append(dests, renameCLocation(conn.Destination, 1))
			}
		} else {
			if conn.Source == sourceEventName {
				dests = append(dests, renameCLocation(conn.Destination, 1))
			}
		}
	}
	return dests
}

//nextPossibleECCStates will find all states in the ECC in this basicFB that the current state feeds into
func nextPossibleECCStates(basicFB iec61499.BasicFB, curState iec61499.ECState) []iec61499.ECState {
	var nextStates []iec61499.ECState

out:
	for _, conn := range basicFB.Transistions {
		if conn.Source == curState.Name {
			for _, state := range basicFB.States {
				if state.Name == conn.Destination {
					nextStates = append(nextStates, state)
					continue out
				}
			}
		}
	}

	return nextStates
}

//findAlgorithmFromName searches a basicFB for an algorithm with the matching name
//if it does not find one, it returns an empty algorithm
func findAlgorithmFromName(basicFB iec61499.BasicFB, name string) iec61499.Algorithm {
	for _, alg := range basicFB.Algorithms {
		if alg.Name == name {
			return alg
		}
	}
	return iec61499.Algorithm{}
}

//blockNeedsCvode will return true if any algorithms inside the BFB need CVODE
func blockNeedsCvode(b iec61499.FB) bool {
	if b.BasicFB == nil {
		return false //only BFBs can need cvode
	}

	for i := 0; i < len(b.BasicFB.Algorithms); i++ {
		if algorithmNeedsCvode(b.BasicFB.Algorithms[i]) || algorithmNeedsCvodeInit(b.BasicFB.Algorithms[i]) {
			return true
		}
	}

	return false
}

//algorithmNeedsCvode/Init will return true if comment is ODE or ODE_init
//ODE is used to tick an ODE
//ODE_init is used to set ODE to some value at some time
func algorithmNeedsCvode(a iec61499.Algorithm) bool {
	return a.Comment == "ODE"
}
func algorithmNeedsCvodeInit(a iec61499.Algorithm) bool {
	return a.Comment == "ODE_init"
}

func stateIsCvodeSetup(b iec61499.BasicFB, s iec61499.ECState) bool {
	for _, action := range s.ECActions {
		if algorithmNeedsCvodeInit(findAlgorithmFromName(b, action.Algorithm)) {
			return true
		}
	}
	return false
}

//CvodeInit is used in templates when generating code from Cvode_init algorithms
type CvodeInit struct {
	OdeFName string
	Initials []OdeVar
}

//CvodeTick is used to store variables and ode functions when processing ODE states
type CvodeTick struct {
	Vars     []OdeVar //the internal vars such as X_dot
	EmitVars []OdeVar //the emitted vars such as Y
}

//OdeVar is a varname and a varvalue used to store variables and string functions
type OdeVar struct {
	VarName      string
	VarValue     string
	TriggerValue string //the trigger value (if any) (i.e. the crossing which might trigger a state change)
}

func (c CvodeInit) GetInitialValues() []OdeVar {
	return c.Initials
	// //check to see if it's just a number
	// if _, err := strconv.ParseFloat(c.Initial, 64); err == nil {
	// 	return c.Initial
	// }
	// //if not, add me-> to it
	// return "me->" + c.Initial
}

func parseOdeInitAlgo(initAlgo string) CvodeInit {
	lines := strings.Split(initAlgo, "\n")

	c := CvodeInit{}

	nameRegex := regexp.MustCompile(`ode[\s]+\=[\s]+([a-zA-Z0-9\_]+);`)
	for _, line := range lines {
		nameMatch := nameRegex.FindStringSubmatch(line)
		if len(nameMatch) == 2 {
			c.OdeFName = nameMatch[1]
		}
	}

	primeRegex := regexp.MustCompile(`([a-zA-Z0-9]+)_prime\s+=\s+([^;]+);`)
	for _, line := range lines {
		nameMatch := primeRegex.FindStringSubmatch(line)
		if len(nameMatch) == 3 {
			c.Initials = append(c.Initials, OdeVar{VarName: nameMatch[1], VarValue: nameMatch[2]})
		}
	}

	return c
}

//parseOdeRunAlgo takes an algorithm that specifies an ODE and parses it into a CvodeTick
// suitable for embedding into a template
func parseOdeRunAlgo(s string) CvodeTick {
	lines := strings.Split(s, "\n")

	c := CvodeTick{}

	dotRegex := regexp.MustCompile(`([a-zA-Z0-9]+)_dot\s+=\s+([^;]+);`)
	for _, line := range lines {
		nameMatch := dotRegex.FindStringSubmatch(line)
		if len(nameMatch) == 3 {
			c.Vars = append(c.Vars, OdeVar{VarName: nameMatch[1], VarValue: nameMatch[2]})
		}
	}

	emitRegex := regexp.MustCompile(`([a-zA-Z0-9\_]+)\s+=\s+([^;]+);`)
	for _, line := range lines {
		nameMatch := emitRegex.FindStringSubmatch(line)
		if len(nameMatch) == 3 {
			//by searching for things that can include "_" in the regex then ignoring them if they do
			//we prevent an annoying case where x_trigger = 7 will make a variable called trigger = 7
			if !strings.Contains(nameMatch[1], "_") {
				c.EmitVars = append(c.EmitVars, OdeVar{VarName: nameMatch[1], VarValue: nameMatch[2]})
			}
		}
	}

	triggerRegex := regexp.MustCompile(`([a-zA-Z0-9]+)_trigger\s+=\s+([^;]+);`)
	for _, line := range lines {
		nameMatch := triggerRegex.FindStringSubmatch(line)
		if len(nameMatch) == 3 {
			found := false
			for i := 0; i < len(c.Vars); i++ { //apply trigger to Vars that match the trigger var name
				if c.Vars[i].VarName == nameMatch[1] {
					c.Vars[i].TriggerValue = nameMatch[2]
					found = true
				}
				break
			}
			if !found {
				panic("Couldn't find variable " + nameMatch[1] + " for _trigger")
			}
		}
	}

	return c
}

//fixOdeVarNameInF is used in the _f functions of ODEs because we can't refer to the relevant variables as me->whatever in this
//  function and this function only, instead, we have to refer to them using the NV_Ith_S(ode_solution, index) notation
func fixOdeVarNameInF(curEval string, name string, index int) string {
	return strings.Replace(curEval, "me->"+name, fmt.Sprintf("NV_Ith_S(ode_solution, %v)", index), -1)
}
