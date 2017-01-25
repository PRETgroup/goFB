// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block SawmillLaser
#ifndef SAWMILLLASER_H_
#define SAWMILLLASER_H_

#include "fbtypes.h"

//This is a BFB, so we need an enum type for the state machine
enum SawmillLaser_states { STATE_SawmillLaser_Start };


//this block had no input events


union SawmillLaserOutputEvents {
	struct {
		UDINT LaserChanged : 1;
	} event;
	UDINT events[1];
};


typedef struct {
    //input events
	

    //output events
	union SawmillLaserOutputEvents outputEvents;

    //input vars
	
    //output vars
	BOOL LaserBroken;
    
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum SawmillLaser_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	
} SawmillLaser_t;

//all FBs get a preinit function
int SawmillLaser_preinit(SawmillLaser_t *me);

//all FBs get an init function
int SawmillLaser_init(SawmillLaser_t *me);

//all FBs get a run function
void SawmillLaser_run(SawmillLaser_t *me);




#endif
