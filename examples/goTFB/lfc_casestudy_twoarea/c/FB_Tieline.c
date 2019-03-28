// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for Tieline
#include "FB_Tieline.h"


/* Tieline_preinit() is required to be called to 
 * initialise an instance of Tieline. 
 * It sets all I/O values to zero.
 */
int Tieline_preinit(Tieline_t  *me) {
	

	//reset the input events
	me->inputEvents.event.Tick = 0;
	//reset the input events
	me->inputEvents.event.DfChange = 0;
	
	//reset the output events
	me->outputEvents.event.Dp12Change = 0;
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	me->StepSize = 0.1;
	me->T0 = 0.0707;
	me->Pi = 3.14159;
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	me->_state = STATE_Tieline_reset;
	me->_trigger = true;
	
	
	

	

	return 0;
}

/* Tieline_init() is required to be called to 
 * set up an instance of Tieline. 
 * It passes around configuration data.
 */
int Tieline_init(Tieline_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void Tieline_TielineTick(Tieline_t  *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
me->Dp12 = me->Dp12 - (((2 * me->Pi) * me->T0) * (me->Df1 - me->Df2)) * me->StepSize;

}



/* Tieline_run() executes a single tick of an
 * instance of Tieline according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will already be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void Tieline_run(Tieline_t  *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.Dp12Change = 0;
	
	
	

	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_Tieline_reset:
			if(me->inputEvents.event.Tick) {
				
				me->_state = STATE_Tieline_update;
				me->_trigger = true;
			};
			break;
		case STATE_Tieline_update:
			if(me->inputEvents.event.Tick) {
				
				me->_state = STATE_Tieline_update;
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
		case STATE_Tieline_reset:
			#ifdef PRINT_STATES
				printf("Tieline: [Entered State reset]\n");
			#endif
			
			break;
		case STATE_Tieline_update:
			#ifdef PRINT_STATES
				printf("Tieline: [Entered State update]\n");
			#endif
			me->outputEvents.event.Dp12Change = 1;
			Tieline_TielineTick(me);
			
			break;
		
		default: 
			break;
		}
	}

	me->_trigger = false;

	

	//Ensure input events are cleared
	me->inputEvents.event.Tick = 0;
	me->inputEvents.event.DfChange = 0;
	
}


