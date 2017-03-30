// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block TrainHA
#ifndef TRAINHA_H_
#define TRAINHA_H_

#include "fbtypes.h"


#include "cvode/cvode.h"
#include "nvector/nvector_serial.h"
#include "cvode/cvode_dense.h"
#include "sundials/sundials_dense.h"
#include "sundials/sundials_types.h"


//This is a BFB, so we need an enum type for the state machine
enum TrainHA_states { STATE_TrainHA_slow_mode_1_setup_0, STATE_TrainHA_slow_mode_1, STATE_TrainHA_fast_mode, STATE_TrainHA_slow_mode_2, STATE_TrainHA_fast_mode_setup_0, STATE_TrainHA_slow_mode_2_setup_0 };


union TrainHAInputEvents {
	struct {
		UDINT tick;
	} event;
	
};


union TrainHAOutputEvents {
	struct {
		UDINT update;
	} event;
	
};


typedef struct {
    //input events
	union TrainHAInputEvents inputEvents;

    //output events
	union TrainHAOutputEvents outputEvents;

    //input vars
	LREAL delta;
    LREAL Vs;
    LREAL Vf;
    
    //output vars
	LREAL x;
    
	//any internal vars (BFBs only)
    BOOL Variable1;
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum TrainHA_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	//this block uses cvode
	void *cvode_mem;
	N_Vector ode_solution;
	realtype T0;
	

	
} TrainHA_t;

//all FBs get a preinit function
int TrainHA_preinit(TrainHA_t  *me);

//all FBs get an init function
int TrainHA_init(TrainHA_t  *me);

//all FBs get a run function
void TrainHA_run(TrainHA_t  *me);



#endif
