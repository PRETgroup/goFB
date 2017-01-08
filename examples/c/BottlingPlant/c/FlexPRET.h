// This file has been automatically generated by goFB and should not be edited by hand
// Transpiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block FlexPRET
#ifndef FLEXPRET_H_
#define FLEXPRET_H_

#include "fbtypes.h"

//This is a CFB, so we need the #includes for the child blocks embedded here
#include "IOManager.h"
#include "CanisterCounter.h"
#include "DoorController.h"
#include "ConveyorController.h"
#include "RejectArmController.h"
#include "InjectorPumpsController.h"
#include "InjectorMotorController.h"


//this block had no input events


//this block had no output events


struct FlexPRET {
    //input events
	

    //output events
	

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	struct IOManager IO;
	struct CanisterCounter CCounter;
	struct DoorController Door;
	struct ConveyorController Conveyor;
	struct RejectArmController RejectArm;
	struct InjectorPumpsController Pumps;
	struct InjectorMotorController Motor;
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	
};

//all FBs get an init function
void FlexPRET_init(struct FlexPRET *me);

//all FBs get a run function
void FlexPRET_run(struct FlexPRET *me);

//composite/resource/device FBs get sync functions
void FlexPRET_syncEvents(struct FlexPRET *me);
void FlexPRET_syncData(struct FlexPRET *me);

#endif