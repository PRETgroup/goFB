// This file has been automatically generated by goFB and should not be edited by hand
// Transpiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for InjectorMotorController
#include "InjectorMotorController.h"




/* InjectorMotorController_init() is required to be called to 
 * initialise an instance of InjectorMotorController. 
 * It sets all I/O values to zero.
 */
void InjectorMotorController_init(struct InjectorMotorController *me) {
	//if there are input events, reset them
	me->inputEvents.events[0] = 0;
	
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs only), call this same function on them
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_InjectorMotorController_MoveArmUp;
	
}



/* InjectorMotorController_run() executes a single tick of an
 * instance of InjectorMotorController according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void InjectorMotorController_run(struct InjectorMotorController *me) {
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
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

//algorithms

void InjectorMotorController_SetArmDownPosition(struct InjectorMotorController *me) {
me->InjectorPosition = 255;
//printf("Injector: Set Injector Arm to Down position\n");
}

void InjectorMotorController_SetArmUpPosition(struct InjectorMotorController *me) {
printf("Injector: Set injector arm to up position\n");
me->InjectorPosition = 0;

}

void InjectorMotorController_Algorithm1(struct InjectorMotorController *me) {
printf("lalalala\n");
}


