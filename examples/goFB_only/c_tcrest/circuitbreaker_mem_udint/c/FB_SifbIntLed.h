// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block SifbIntLed
#ifndef SIFBINTLED_H_
#define SIFBINTLED_H_

#include "fbtypes.h"
#include "util.h"



//This is a BFB with states, so we need an enum type for the state machine
enum SifbIntLed_states { STATE_SifbIntLed_Start, STATE_SifbIntLed_Update };




union SifbIntLedInputEvents {
	struct {
		UDINT i_change;
	} event;
	
};


//this block had no output events


typedef struct {
    //input events
	union SifbIntLedInputEvents inputEvents;

    //output events
	

    //input vars
	INT i;
    
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	// enum SifbIntLed_states _state; //stores current state
	// BOOL _trigger; //indicates if a state transition has occured this tick
	
	
	

	

} SifbIntLed_t;

//all FBs get a preinit function
int SifbIntLed_preinit(SifbIntLed_t  *me);

//all FBs get an init function
int SifbIntLed_init(SifbIntLed_t  *me);

//all FBs get a run function
void SifbIntLed_run(SifbIntLed_t  *me);



#endif
