package iec61499

import "errors"

/* an example of a conversion

policyFB AEIPolicy;
interface of AEIPolicy {
	in event AS, VS; //in here means that they're going from PLANT to CONTROLLER
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

SHOULD BECOME

basicFB AEIPolicyEnforcer;
interface of AEIPolicyEnforcer {
	in event AS_poci_in, VS_poci_in; //PLANT OUT CONTROLLER IN, IN/OUT w.r.t. ENFORCER
	out event AS_poci_out, VS_poci_out;

	in event AP_pico_in, VP_pico_in; //PLANT IN CONTROLLER OUT, IN/OUT w.r.t. ENFORCER
	out event AP_pico_out, VP_pico_out;
}
architecture of AEIPolicyEnforcer {
	internals {
		ulint AEI_ns := 900000000;
		dtimer tAEI; //DTIMER increases in DISCRETE TIME continuously
	}

	//ISSUE: the various signals need to always be propogated. Unless we set this up with a number of blocks, this is going to be a mess.
	//(I.e. a parent CFB, then a block to resolve signal conflicts, etc.)

	states {

		s1 {
			run in 'ST' `tAEI := tAEI + 1;` //DTIMERS added to EVERY STATE that they are not INCREMENTED IN
			-> s2_e on (VS_poci_in or VP_pico_in);
		}
		s2_e {								//ENTRY states for EVERY PFBTRANSITION with EXPRESSIONS
			run in 'ST' `tAEI := 0;`		//DTIMERS might be RESET in an ENTRY STATE
			-> s1_e on (AS_poci_in or AP_pico_in);
			-> s2_recovery_0 on (tAEI > AEI_ns);	//VIOLATION transitions go to RECOVERY states (unique for each transition)
			-> s2;
		}
		s2 {
			run in 'ST' `tAEI := tAEI + 1;`
			-> s1_e on (AS_poci_in or AP_pico_in);
			-> s2_recovery_0 on (tAEI > AEI_ns);	//VIOLATION transitions go to RECOVERY states (unique for each transition)
		}

		s2_recovery_0 {
			emit AP_pico_out;					//The RECOVERY algorithm
			//run in 'ST' `AP_pico_out`
		}
	}
}

*/

//TranslatePFBtoBFB will take a Policy Function Block and compile it to its equivalent
// Basic Function Block
//It operates according to the algorithm specified in [TODO: Paper link]
func (f *FB) TranslatePFBtoBFB() error {
	// if f.PolicyFB == nil {
	// 	return errors.New("TranslatePFBtoBFB can only be called on an PolicyFB")
	// }

	return errors.New("Not yet implemented")
}
