// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block load
#ifndef LOAD_H_
#define LOAD_H_

#include "fbtypes.h"
#include "util.h"


#include "cvode/cvode.h"
#include "nvector/nvector_serial.h"
#include "cvode/cvode_dense.h"
#include "sundials/sundials_dense.h"
#include "sundials/sundials_types.h"


//This is a BFB with states, so we need an enum type for the state machine
enum load_states { STATE_load_INIT, STATE_load_L_OFF_E0, STATE_load_L_OFF_E1, STATE_load_L_OFF_E2, STATE_load_L_OFF, STATE_load_L_INRUSH_E0, STATE_load_L_INRUSH, STATE_load_L_NOM_E0, STATE_load_L_NOM, STATE_load_L_FAULT_E0, STATE_load_L_FAULT };




union loadInputEvents {
	struct {
		UDINT on;
		UDINT off;
		UDINT fault;
	} event;
	
};


union loadOutputEvents {
	struct {
		UDINT i_change;
	} event;
	
};


typedef struct {
    //input events
	union loadInputEvents inputEvents;

    //output events
	union loadOutputEvents outputEvents;

    //input vars
	
    //output vars
	LREAL i;
    
	//any internal vars (BFBs only)
    LREAL DeltaTime;
    LREAL i_fault;
    LREAL i_inrush;
    LREAL i_nom;
    LREAL t_up;
    LREAL t_down;
    LREAL t_fault;
    LREAL x;
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum load_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	//this block uses cvode
	void *cvode_mem;
	N_Vector ode_solution;
	realtype T0;
	realtype Tnext;
	realtype Tcurr;
	int solveInProgress;
	
	
	

	

} load_t;

//all FBs get a preinit function
int load_preinit(load_t  *me);

//all FBs get an init function
int load_init(load_t  *me);

//all FBs get a run function
void load_run(load_t  *me);



#endif
