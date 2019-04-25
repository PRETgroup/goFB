// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block BfbIDMTCurve
#ifndef BFBIDMTCURVE_H_
#define BFBIDMTCURVE_H_

#include "fbtypes.h"
#include "util.h"



//This is a BFB with states, so we need an enum type for the state machine
enum BfbIDMTCurve_states { STATE_BfbIDMTCurve_s_init, STATE_BfbIDMTCurve_s_safe, STATE_BfbIDMTCurve_s_count, STATE_BfbIDMTCurve_s_unsafe };




union BfbIDMTCurveInputEvents {
	struct {
		UDINT tick;
		UDINT i_measured;
		UDINT iSet_change;
	} event;
	
};


union BfbIDMTCurveOutputEvents {
	struct {
		UDINT unsafe;
	} event;
	
};


typedef struct {
    //input events
	union BfbIDMTCurveInputEvents inputEvents;

    //output events
	union BfbIDMTCurveOutputEvents outputEvents;

    //input vars
	USINT i;
    USINT iSet;
    
    //output vars
	
	//any internal vars (BFBs only)
    ULINT cnt;
    LREAL max;
    REAL k;
    REAL b;
    REAL a;
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum BfbIDMTCurve_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	
	
	

	

} BfbIDMTCurve_t;

//all FBs get a preinit function
int BfbIDMTCurve_preinit(BfbIDMTCurve_t  *me);

//all FBs get an init function
int BfbIDMTCurve_init(BfbIDMTCurve_t  *me);

//all FBs get a run function
void BfbIDMTCurve_run(BfbIDMTCurve_t  *me);



#endif
