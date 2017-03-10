// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for InjectorMotorController
#include "InjectorMotorController.h"


/* InjectorMotorController_preinit() is required to be called to 
 * initialise an instance of InjectorMotorController. 
 * It sets all I/O values to zero.
 */
int InjectorMotorController_preinit(InjectorMotorController_t  *me) {
	//if there are input events, reset them
	me->inputEvents.event.InjectorArmFinishedMovement = 0;
	me->inputEvents.event.EmergencyStopChanged = 0;
	me->inputEvents.event.ConveyorStoppedForInject = 0;
	me->inputEvents.event.PumpFinished = 0;
	
	//if there are output events, reset them
	me->outputEvents.event.StartPump = 0;
	me->outputEvents.event.InjectDone = 0;
	me->outputEvents.event.InjectorPositionChanged = 0;
	me->outputEvents.event.InjectRunning = 0;
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_InjectorMotorController_MoveArmUp;
	

	return 0;
}

/* InjectorMotorController_init() is required to be called to 
 * set up an instance of InjectorMotorController. 
 * It passes around configuration data.
 */
int InjectorMotorController_init(InjectorMotorController_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void InjectorMotorController_SetArmDownPosition(InjectorMotorController_t  *me) {
me->InjectorPosition = 255;
//printf("Injector: Set Injector Arm to Down position\n");
}

void InjectorMotorController_SetArmUpPosition(InjectorMotorController_t  *me) {
//printf("Injector: Set injector arm to up position\n");
me->InjectorPosition = 0;

}

void InjectorMotorController_Algorithm1(InjectorMotorController_t  *me) {
//printf("lalalala\n");
}



/* InjectorMotorController_run() executes a single tick of an
 * instance of InjectorMotorController according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void InjectorMotorController_run(InjectorMotorController_t  *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.StartPump = 0;
	me->outputEvents.event.InjectDone = 0;
	me->outputEvents.event.InjectorPositionChanged = 0;
	me->outputEvents.event.InjectRunning = 0;
	

	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_InjectorMotorController_MoveArmUp:
			if(me->inputEvents.event.InjectorArmFinishedMovement) {
				me->_state = STATE_InjectorMotorController_Await_Bottle;
				me->_trigger = true;
				
			};
			break;
		case STATE_InjectorMotorController_Await_Bottle:
			if(me->inputEvents.event.ConveyorStoppedForInject) {
				me->_state = STATE_InjectorMotorController_MoveArmDown;
				me->_trigger = true;
				
			};
			break;
		case STATE_InjectorMotorController_MoveArmDown:
			if(me->inputEvents.event.InjectorArmFinishedMovement) {
				me->_state = STATE_InjectorMotorController_Await_Pumping;
				me->_trigger = true;
				
			};
			break;
		case STATE_InjectorMotorController_Await_Pumping:
			if(me->inputEvents.event.PumpFinished) {
				me->_state = STATE_InjectorMotorController_MoveArmUp;
				me->_trigger = true;
				
			};
			break;
		
		}
	}

	//state output logic
	if(me->_trigger == true) {
		switch(me->_state) {
		case STATE_InjectorMotorController_MoveArmUp:
			InjectorMotorController_SetArmUpPosition(me);
			me->outputEvents.event.InjectorPositionChanged = 1;
			break;

		case STATE_InjectorMotorController_Await_Bottle:
			me->outputEvents.event.InjectDone = 1;
			break;

		case STATE_InjectorMotorController_MoveArmDown:
			InjectorMotorController_SetArmDownPosition(me);
			me->outputEvents.event.InjectorPositionChanged = 1;
			me->outputEvents.event.InjectRunning = 1;
			break;

		case STATE_InjectorMotorController_Await_Pumping:
			me->outputEvents.event.StartPump = 1;
			break;

		
		}
	}

	me->_trigger = false;
}


