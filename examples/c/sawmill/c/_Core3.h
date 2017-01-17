// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block _Core3
#ifndef _CORE3_H_
#define _CORE3_H_

#include "fbtypes.h"

//This is a CFB, so we need the #includes for the child blocks embedded here
#include "ArgoTx.h"
#include "SawmillMessageHandler.h"
#include "SawmillModule.h"


//this block had no input events


//this block had no output events


struct _Core3 {
    //input events
	

    //output events
	

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	struct ArgoTx tx;
	struct SawmillMessageHandler messageHandler;
	struct SawmillModule sawmill;
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	
};

//all FBs get a preinit function
int _Core3_preinit(struct _Core3 *me);

//all FBs get an init function
int _Core3_init(struct _Core3 *me);

//all FBs get a run function
void _Core3_run(struct _Core3 *me);

//composite/resource/device FBs get sync functions
void _Core3_syncEvents(struct _Core3 *me);
void _Core3_syncData(struct _Core3 *me);

#endif