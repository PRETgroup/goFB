// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for DoorController
#include "DoorController.h"


/* DoorController_preinit() is required to be called to 
 * initialise an instance of DoorController. 
 * It sets all I/O values to zero.
 */
int DoorController_preinit(struct DoorController *me) {
	//if there are input events, reset them
	me->inputEvents.events[0] = 0;
	
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_DoorController_E_Stop;
	

	return 0;
}

/* DoorController_init() is required to be called to 
 * set up an instance of DoorController. 
 * It passes around configuration data.
 */
int DoorController_init(struct DoorController *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



/* DoorController_run() executes a single tick of an
 * instance of DoorController according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void DoorController_run(struct DoorController *me) {
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_DoorController_E_Stop:
			if(me->inputEvents.event.EmergencyStopChanged && ( ! me->EmergencyStop)) {
				me->_state = STATE_DoorController_Await;
				me->_trigger = true;
			};
			break;
		case STATE_DoorController_Run:
			if(me->inputEvents.event.EmergencyStopChanged && (me->EmergencyStop)) {
				me->_state = STATE_DoorController_E_Stop;
				me->_trigger = true;
			} else if(me->inputEvents.event.ReleaseDoorOverride || me->inputEvents.event.BottlingDone) {
				me->_state = STATE_DoorController_Run;
				me->_trigger = true;
			};
			break;
		case STATE_DoorController_Await:
			if(me->inputEvents.event.ReleaseDoorOverride || me->inputEvents.event.BottlingDone) {
				me->_state = STATE_DoorController_Run;
				me->_trigger = true;
			};
			break;
		
		}
	}

	//state output logic
	if(me->_trigger == true) {
		switch(me->_state) {
		case STATE_DoorController_E_Stop:
			break;

		case STATE_DoorController_Run:
			me->outputEvents.event.DoorReleaseCanister = 1;
			break;

		case STATE_DoorController_Await:
			break;

		
		}
	}

	me->_trigger = false;
}

//no algorithms were present for this function block


