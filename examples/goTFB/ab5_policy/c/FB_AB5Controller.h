// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block AB5Controller
#ifndef AB5CONTROLLER_H_
#define AB5CONTROLLER_H_

#include "fbtypes.h"
#include "util.h"



//This is a BFB with states, so we need an enum type for the state machine
enum AB5Controller_states { STATE_AB5Controller_unknown };


//This is an FB with policies, so we need an enum type for the state machine of each policy

enum AB5Controller_policy_AB5_states { POLICY_STATE_AB5Controller_AB5_s0, POLICY_STATE_AB5Controller_AB5_s1 };


union AB5ControllerInputEvents {
	struct {
		UDINT A;
	} event;
	
};


union AB5ControllerOutputEvents {
	struct {
		UDINT B;
	} event;
	
};


typedef struct {
    //input events
	union AB5ControllerInputEvents inputEvents;

    //output events
	union AB5ControllerOutputEvents outputEvents;

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum AB5Controller_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	
	
	

	//this block has policies
	enum AB5Controller_policy_AB5_states _policy_AB5_state;
	//output vars
	DTIMER v;
	
	
	

} AB5Controller_t;

//all FBs get a preinit function
int AB5Controller_preinit(AB5Controller_t  *me);

//all FBs get an init function
int AB5Controller_init(AB5Controller_t  *me);

//all FBs get a run function
void AB5Controller_run(AB5Controller_t  *me);



#endif