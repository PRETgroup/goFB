package iec61499

import "errors"

/* an example of a conversion

//event A is from Plant
//event B is from Controller
//A and B cannot happen simultaneously, A and B alternate starting with an A. B should be true within 5 ticks of B.

policyFB AB5Policy;
interface of AB5Policy {
	in event A;  //in here means that they're going from PLANT to CONTROLLER
	out event B; //out here means that they're going from CONTROLLER to PLANT
}

architecture of AB5Policy {
	internals {
		dtimer v;
	}

	states {
		s0 {														//first state is initial, and represents "We're waiting for an A"
			-> s0 on (!A and !B): v := 0;							//if we receive neither A nor B, do nothing
			-> s1 on (A and !B): v := 0;							//if we receive an A only, head to state s1
			-> violation on ((!A and B) or (A and B));				//if we receive a B, or an A and a B (i.e. if we receive a B) then VIOLATION
		}

		s1 {														//s1 is "we're waiting for a B, and it needs to get here within 5 ticks"
			-> s1 on (!A and !B and v < 5);							//if we receive nothing, and we aren't over-time, then we do nothing
			-> s0 on (!A and B);									//if we receive a B only, head to state s0
			-> violation on ((v >= 5) or (A and B) or (A and !B));	//if we go overtime, or we receive another A, then VIOLATION
		}
	}
}

CAN DERIVE TO INPUT POLICY (Remember, replace all outputs and things that depend on outputs with TRUE)

policyFB AB5Policy_INPUT;
interface of AB5Policy_INPUT {
	in event A;  //in here means that they're going from PLANT to CONTROLLER
}

architecture of AB5Policy_INPUT {
	internals {
		dtimer v;
	}

	states {
		s0 {
			-> s0 on (!A and (true)): v := 0;
			-> s1 on (A and (true)): v := 0;
			-> violation on ((!A and (true)) or (A and (true)));
		}

		s1 {
			-> s1 on (!A and (true) and v < 5);
			-> s0 on (!A and (true));
			-> violation on ((v >= 5) or (A and (true)) or (A and (true)));
		}
	}
}

IS EQUIVALENT TO

policyFB AB5Policy_INPUT;
interface of AB5Policy_INPUT {
	in event A;  //in here means that they're going from PLANT to CONTROLLER
}

architecture of AB5Policy_INPUT {
	internals {
		dtimer v;
	}

	states {
		s0 {
			-> s0 on (!A): v := 0;
			-> s1 on (A): v := 0;
			-> violation on (!A or A);
		}

		s1 {
			-> s1 on (!A and v < 5);
			-> s0 on (!A);
			-> violation on ((v >= 5) or (A));
		}
	}
}

Now remember, the enforcer for input will only make a change if an input _cannot_ result in a safe state.

In this case, in s0, every input (of A) can only result in a safe state, as both s0 is state and s1 has the potential to be safe
(I.e. we don't have to take the transition to the violation, so why would we?

In s1, if we receive an A, we're always going to end up in a violation. Stink!
So, what we'll do instead is get our enforcer to prevent the transition from occuring by instead picking a non-violating transition.
Both other transitions suggest !A, so that's what we'll do instead.

In s1, if v >= 5, that doesn't mean a violation _has_ to occur, because a !A could also have a non-violation.

WE CONVERT THESE TO ENFORCERS

every violation potential must be associated with a non-violating transition
// -> s0 on (!A): v := 0;
// -> s1 on (A): v := 0;
// -> violation on (!A or A);
s0 {
	if(!A) { 			//violation potential "(!A)"
						//auto-selected non-violating transition "-> s0 on (!A)", no edits required
	}

	if(A) {				//violation potential
						//auto-selected non-violating transition "-> s1 on (A)", no edits required
	}
}

// -> s1 on (!A and v < 5);
// -> s0 on (!A);
// -> violation on ((v >= 5) or (A));
s1 {
	if(v >= 5) { 		//violation potential "(v >= 5)"
		A = 0;			//auto-selected non-violating transition "-> s0 on (!A)", edit might be required
	}
	if(A) { 			//violation potential "(A)"
		A = 0;			//auto-selected non-violating transition "-> s0 on (!A)", edit might be required
	}
}

OUTPUT:

// -> s0 on (!A and !B): v := 0;
// -> s1 on (A and !B): v := 0;
// -> violation on ((!A and B) or (A and B));
s0 {
	//perform edits first
	if(!A and B) {		//violation potential "(!A and B)"
		B = 0;			//auto-selected non-violating transition "-> s0 on (!A and !B)", edit might be required
	}
	if(A and B) {		//violation potential "(A and B)"
		B = 0;			//auto-selected non-violating transition "-> s1 on (A and !B)", edit might be required
	}

	//now advance state
	if(!A and !B) {
		state = s0;
		v = 0;
	}
	if(A and !B) {
		state = s1;
		v = 0;
	}
}

// -> s1 on (!A and !B and v < 5);
// -> s0 on (!A and B);
// -> violation on ((v >= 5) or (A and B) or (A and !B));
s1 {
	//perform edits first
	if(v >= 5) {		//violation potential "(v >= 5)"
		B = 1;			//auto selected non-time non-violating transition "-> s0 on (!A and B)", edit might be required
	}
	if(A and B) {		//violation potential "(A and B)"
		B = 0;			//auto selected non-time non-violating transition "-> s0 on (!A and B)", edit might be required
	}
	if(A and !B) {

	}
}


*/

//TranslatePFBtoSIFB will take a Policy Function Block and compile it to its enforcer as a
// Service Interface Function Block
//It operates according to the algorithm specified in [TODO: Paper link]
func (f *FB) TranslatePFBtoSIFB() error {
	if f.PolicyFB == nil {
		return errors.New("TranslatePFBtoBFB can only be called on an PolicyFB")
	}

	return errors.New("Not yet implemented")
}

/*


policyFB AEIPolicy;
interface of AEIPolicy {
	in event AS, VS; //in here means that they're going from PLANT to CONTROLLER, and the EnforcerI is derived from this
	out event AP, VP;//out here means that they're going from CONTROLLER to PLANT
}
architecture of AEIPolicy {
	internals {
		ulint AEI_ns := 900000000;
		dtimer tAEI; //DTIMER increases in DISCRETE TIME continuously
	}

	//P3: AS or AP must be true within AEI after a ventricular event VS or VP.

	states {
		s1 {
			//-> <destination> [on guard] [: output expression][, output expression...] ;
			-> s2 on (VS or VP): tAEI := 0;
		}

		s2 {
			-> s1 on (AS or AP);
			-> violation on (tAEI > AEI_ns);
		}
	}
}

SHOULD BECOME TWO BLOCKS

serviceFB AEIPolicyEnforcerI;
interface of AEIPolicyEnforcerI {
	in event AS_poci, VS_poci; //PLANT OUT CONTROLLER IN, IN/OUT w.r.t. ENFORCER
	out event AS_poci_prime, VS_poci_prime;
}

architecture of AEIPolicyEnforcerI {

}




serviceFB AEIPolicyEnforcerO;
interface of AEIPolicyEnforcerO {
	in event AS_poci_prime, VS_poci_prime; //PLANT OUT CONTROLLER IN, IN/OUT w.r.t. ENFORCER

	in event AP_pico, VP_pico; //PLANT IN CONTROLLER OUT, IN/OUT w.r.t. ENFORCER
	out event AP_pico_prime, VP_pico_prime;
}
architecture of AEIPolicyEnforcerI {
	internals {
		ulint AEI_ns := 900000000;
		dtimer tAEI; //DTIMER increases in DISCRETE TIME continuously
		int _STATE := 0;
	}

	init in 'ST' `

	`;

	run in 'ST' `
		IF WE GO TO SAFE I STATE
			_poci_outs <= _poci_ins
		ELSE
			_poci_outs <= modified(poci_ins)
		END IF

		RUN CONTROLLER

		IF WE GO TO SAFE STATE
			_pico_outs <= pico_ins
		ELSE
			_pico_outs <= modified(pico_outs)
		END IF

		RUN PLANT

		//CAN we pipeline this?

		//ie run every tick, where we take PRE

		//        controller... enforcer... plant
		//tick 0             <-          <-  a
		//tick 0             ->          ->

		//tick 1             <-     a    <-  b
		//tick 1             ->          ->

		//tick 2       a'    <-     b    <-  c
		//tick 2             ->          ->

		//tick 3       b'    <-     c    <-  d
		//tick 3             ->     a''  ->

		//tick 4       c'    <-     d    <-  e
		//tick 4             ->     b''  ->  a'''


		//MULTICYCLE (4 cycles per synchronous tick)

		//        controller... enforcer... plant
		//tick 0             <-          <-  a
		//tick 0             ->          ->

		//tick 1             <-     a    <-
		//tick 1             ->          ->

		//tick 2       a'    <-          <-
		//tick 2             ->          ->

		//tick 3             <-          <-
		//tick 3             ->     a''  ->

		//tick 4             <-          <-
		//tick 4             ->          ->  a'''
	`;

}

*/
