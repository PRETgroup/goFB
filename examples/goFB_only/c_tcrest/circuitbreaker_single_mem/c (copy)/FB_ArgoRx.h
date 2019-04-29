// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block ArgoRx
#ifndef ARGORX_H_
#define ARGORX_H_

#include "fbtypes.h"

//This is a BFB, so we need an enum type for the state machine
enum ArgoRx_states { STATE_ArgoRx_Start };


//this block had no input events


union ArgoRxOutputEvents {
	struct {
		UDINT DataPresent : 1;
	} event;
	UDINT events[1];
};


typedef struct {
    //input events
	

    //output events
	union ArgoRxOutputEvents outputEvents;

    //input vars
	UDINT ChanId;
    
    //output vars
	INT Data;
    
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	enum ArgoRx_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick

		qpd_t* chan;
	BOOL needToAck;
	
} ArgoRx_t;

//all FBs get a preinit function
int ArgoRx_preinit(ArgoRx_t *me);

//all FBs get an init function
int ArgoRx_init(ArgoRx_t *me);

//all FBs get a run function
void ArgoRx_run(ArgoRx_t *me);




#endif
