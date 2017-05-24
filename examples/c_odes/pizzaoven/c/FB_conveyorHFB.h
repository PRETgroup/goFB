// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block conveyorHFB
#ifndef CONVEYORHFB_H_
#define CONVEYORHFB_H_

#include "fbtypes.h"
#include "util.h"


#include "cvode/cvode.h"
#include "nvector/nvector_serial.h"
#include "cvode/cvode_dense.h"
#include "sundials/sundials_dense.h"
#include "sundials/sundials_types.h"


//This is a BFB, so we need an enum type for the state machine
enum conveyorHFB_states { STATE_conveyorHFB_Start, STATE_conveyorHFB_Qoff, STATE_conveyorHFB_Qon, STATE_conveyorHFB_enterQon0, STATE_conveyorHFB_enterQoff0, STATE_conveyorHFB_enterQoff1 };


union conveyorHFBInputEvents {
	struct {
		UDINT Start;
		UDINT Off;
		UDINT On;
		UDINT MChange;
	} event;
	
};


union conveyorHFBOutputEvents {
	struct {
		UDINT XChange;
		UDINT DChange;
	} event;
	
};


typedef struct {
    //input events
	union conveyorHFBInputEvents inputEvents;

    //output events
	union conveyorHFBOutputEvents outputEvents;

    //input vars
	LREAL DeltaTime;
    LREAL M;
    
    //output vars
	LREAL X;
    LREAL D;
    
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum conveyorHFB_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	//this block uses cvode
	void *cvode_mem;
	N_Vector ode_solution;
	realtype T0;
	realtype Tnext;
	realtype Tcurr;
	int solveInProgress;
	

	
} conveyorHFB_t;

//all FBs get a preinit function
int conveyorHFB_preinit(conveyorHFB_t  *me);

//all FBs get an init function
int conveyorHFB_init(conveyorHFB_t  *me);

//all FBs get a run function
void conveyorHFB_run(conveyorHFB_t  *me);



#endif