package iec61499converter

import (
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

	re1 := regexp.MustCompile("([<>=!]+)")
	re2 := regexp.MustCompile("([a-zA-Z_<>=]+)")

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
		//check to see if it is an input event
		if block.EventInputs != nil {
			for _, event := range block.EventInputs.Events {
				if event.Name == in {
					events = append(events, "me->inputEvents.event."+event.Name)
					return "me->inputEvents.event." + event.Name //FUTURE WORK: Consider use of pointers and offsets to minimise memory footprint for events? i.e. "*(me->inputEvents." + in + " + ev_offset)"
				}
			}
		}
		//else, return it assuming input data
		return "me->" + in
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

//CvodeInit is used in templates when generating code from Cvode_init algorithms
type CvodeInit struct {
	OdeFName string
	Initials []InitialVar
}

type InitialVar struct {
	VarName  string
	VarValue string
}

func (c CvodeInit) GetInitialValues() []InitialVar {
	return c.Initials
	// //check to see if it's just a number
	// if _, err := strconv.ParseFloat(c.Initial, 64); err == nil {
	// 	return c.Initial
	// }
	// //if not, add me-> to it
	// return "me->" + c.Initial
}

func parseOdeInitAlgo(s string) CvodeInit {

	c := CvodeInit{}

	nameRegex := regexp.MustCompile(`ode[\s]+\=[\s]+([a-zA-Z0-9\_]+);`)

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		nameMatch := nameRegex.FindStringSubmatch(line)
		if len(nameMatch) == 2 {
			c.OdeFName = nameMatch[1]
		}
	}

	c.Initials = []InitialVar{{VarName: "ode_solution", VarValue: "x"}}

	return c
}
