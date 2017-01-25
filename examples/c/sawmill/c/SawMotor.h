// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block SawMotor
#ifndef SAWMOTOR_H_
#define SAWMOTOR_H_

#include "fbtypes.h"

//This is a BFB, so we need an enum type for the state machine
enum SawMotor_states { STATE_SawMotor_Start, STATE_SawMotor_Run, STATE_SawMotor_Stop };


union SawMotorInputEvents {
	struct {
		UDINT CommandChange : 1;
	} event;
	UDINT events[1];
};


//this block had no output events


typedef struct {
    //input events
	union SawMotorInputEvents inputEvents;

    //output events
	

    //input vars
	BOOL Run;
    
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum SawMotor_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	
} SawMotor_t;

//all FBs get a preinit function
int SawMotor_preinit(SawMotor_t *me);

//all FBs get an init function
int SawMotor_init(SawMotor_t *me);

//all FBs get a run function
void SawMotor_run(SawMotor_t *me);




#endif
