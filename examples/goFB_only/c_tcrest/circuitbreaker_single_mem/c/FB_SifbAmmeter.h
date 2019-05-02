// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block SifbAmmeter
#ifndef SIFBAMMETER_H_
#define SIFBAMMETER_H_

#include "fbtypes.h"
#include "util.h"



//This is a BFB with states, so we need an enum type for the state machine
enum SifbAmmeter_states { STATE_SifbAmmeter_Start, STATE_SifbAmmeter_Update };




//this block had no input events


union SifbAmmeterOutputEvents {
	struct {
		UDINT i_measured;
	} event;
	
};


typedef struct {
    //input events
	

    //output events
	union SifbAmmeterOutputEvents outputEvents;

    //input vars
	
    //output vars
	REAL i;
    
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	// enum SifbAmmeter_states _state; //stores current state
	// BOOL _trigger; //indicates if a state transition has occured this tick
	
	
	

	

} SifbAmmeter_t;

//all FBs get a preinit function
int SifbAmmeter_preinit(SifbAmmeter_t  *me);

//all FBs get an init function
int SifbAmmeter_init(SifbAmmeter_t  *me);

//all FBs get a run function
void SifbAmmeter_run(SifbAmmeter_t  *me);



#endif
