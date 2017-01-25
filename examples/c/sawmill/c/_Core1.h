// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block _Core1
#ifndef _CORE1_H_
#define _CORE1_H_

#include "fbtypes.h"

//This is a CFB, so we need the #includes for the child blocks embedded here
#include "SawmillModule.h"
#include "SawmillMessageHandler.h"
#include "ArgoTx.h"


//this block had no input events


//this block had no output events


typedef struct {
    //input events
	

    //output events
	

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	SawmillModule_t sawmill;
	SawmillMessageHandler_t messageHandler;
	ArgoTx_t tx;
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	
} _Core1_t;

//all FBs get a preinit function
int _Core1_preinit(_Core1_t *me);

//all FBs get an init function
int _Core1_init(_Core1_t *me);

//all FBs get a run function
void _Core1_run(_Core1_t *me);

//composite/resource/device FBs get sync functions
void _Core1_syncEvents(_Core1_t *me);
void _Core1_syncData(_Core1_t *me);

#endif
