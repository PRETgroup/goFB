// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block Load
#ifndef LOAD_H_
#define LOAD_H_

#include "fbtypes.h"
#include "util.h"



//This is a BFB with states, so we need an enum type for the state machine
enum Load_states { STATE_Load_reset, STATE_Load_update };




union LoadInputEvents {
	struct {
		UDINT Tick;
		UDINT Dp12Change;
		UDINT DpeExternalChange;
	} event;
	
};


union LoadOutputEvents {
	struct {
		UDINT DpeChange;
	} event;
	
};


typedef struct {
    //input events
	union LoadInputEvents inputEvents;

    //output events
	union LoadOutputEvents outputEvents;

    //input vars
	LREAL Dp12;
    LREAL DpeExternal;
    
    //output vars
	LREAL Dpe;
    
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum Load_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	
	
	

	

} Load_t;

//all FBs get a preinit function
int Load_preinit(Load_t  *me);

//all FBs get an init function
int Load_init(Load_t  *me);

//all FBs get a run function
void Load_run(Load_t  *me);



#endif