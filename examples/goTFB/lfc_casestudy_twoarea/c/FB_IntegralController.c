// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for IntegralController
#include "FB_IntegralController.h"


/* IntegralController_preinit() is required to be called to 
 * initialise an instance of IntegralController. 
 * It sets all I/O values to zero.
 */
int IntegralController_preinit(IntegralController_t  *me) {
	

	//reset the input events
	me->inputEvents.event.Tick = 0;
	//reset the input events
	me->inputEvents.event.DfChange = 0;
	//reset the input events
	me->inputEvents.event.Dp12Change = 0;
	
	//reset the output events
	me->outputEvents.event.DprefChange = 0;
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	me->StepSize = 0.1;
	me->Ki = 0.5;
	me->B = 0.425;
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	me->_state = STATE_IntegralController_reset;
	me->_trigger = true;
	
	
	

	

	return 0;
}

/* IntegralController_init() is required to be called to 
 * set up an instance of IntegralController. 
 * It passes around configuration data.
 */
int IntegralController_init(IntegralController_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void IntegralController_IntegralControllerTick(IntegralController_t  *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
me->Dpref = me->Dpref - (me->Ki * (me->Df * me->B + me->Dp12)) * me->StepSize;

}



/* IntegralController_run() executes a single tick of an
 * instance of IntegralController according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will already be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void IntegralController_run(IntegralController_t  *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.DprefChange = 0;
	
	
	

	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_IntegralController_reset:
			if(me->inputEvents.event.Tick) {
				
				me->_state = STATE_IntegralController_update;
				me->_trigger = true;
			};
			break;
		case STATE_IntegralController_update:
			if(me->inputEvents.event.Tick) {
				
				me->_state = STATE_IntegralController_update;
				me->_trigger = true;
			};
			break;
		
		default: 
			break;
		}
	}

	//state output logic
	if(me->_trigger == true) {
		switch(me->_state) {
		case STATE_IntegralController_reset:
			#ifdef PRINT_STATES
				printf("IntegralController: [Entered State reset]\n");
			#endif
			
			break;
		case STATE_IntegralController_update:
			#ifdef PRINT_STATES
				printf("IntegralController: [Entered State update]\n");
			#endif
			me->outputEvents.event.DprefChange = 1;
			IntegralController_IntegralControllerTick(me);
			
			break;
		
		default: 
			break;
		}
	}

	me->_trigger = false;

	

	//Ensure input events are cleared
	me->inputEvents.event.Tick = 0;
	me->inputEvents.event.DfChange = 0;
	me->inputEvents.event.Dp12Change = 0;
	
}


