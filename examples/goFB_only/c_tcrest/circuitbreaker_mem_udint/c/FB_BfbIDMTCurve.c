// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for BfbIDMTCurve
#include "FB_BfbIDMTCurve.h"


/* BfbIDMTCurve_preinit() is required to be called to 
 * initialise an instance of BfbIDMTCurve. 
 * It sets all I/O values to zero.
 */
int BfbIDMTCurve_preinit(BfbIDMTCurve_t  *me) {
	

	//reset the input events
	me->inputEvents.event.tick = 0;
	//reset the input events
	me->inputEvents.event.i_measured = 0;
	//reset the input events
	me->inputEvents.event.i_set_change = 0;
	
	//reset the output events
	me->outputEvents.event.unsafe = 0;
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	me->cnt = 0;
	me->max = 0.0;
	me->k = 100;
	me->b = 135;
	me->a = 1;
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	me->_state = STATE_BfbIDMTCurve_INIT;
	me->_trigger = true;
	
	
	

	

	return 0;
}

/* BfbIDMTCurve_init() is required to be called to 
 * set up an instance of BfbIDMTCurve. 
 * It passes around configuration data.
 */
int BfbIDMTCurve_init(BfbIDMTCurve_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void BfbIDMTCurve_ResetCnt(BfbIDMTCurve_t  *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
me->cnt = 0;
}

void BfbIDMTCurve_UpdateCnt(BfbIDMTCurve_t  *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
me->cnt = me->cnt + 1;
}

void BfbIDMTCurve_UpdateMax(BfbIDMTCurve_t  *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
me->max = (me->k * me->b) / ((me->i / me->i_set) - 1);
}



/* BfbIDMTCurve_run() executes a single tick of an
 * instance of BfbIDMTCurve according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will already be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void BfbIDMTCurve_run(BfbIDMTCurve_t  *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.unsafe = 0;
	
	
	

	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_BfbIDMTCurve_INIT:
			if(true) {
				
				me->_state = STATE_BfbIDMTCurve_SAFE;
				me->_trigger = true;
			};
			break;
		case STATE_BfbIDMTCurve_SAFE:
			if(me->inputEvents.event.tick && (me->i > me->i_set)) {
				
				me->_state = STATE_BfbIDMTCurve_COUNT;
				me->_trigger = true;
			};
			break;
		case STATE_BfbIDMTCurve_COUNT:
			if(me->inputEvents.event.tick && (me->cnt >= me->max)) {
				
				me->_state = STATE_BfbIDMTCurve_UNSAFE;
				me->_trigger = true;
			} else if(me->inputEvents.event.tick && (me->i < me->i_set)) {
				
				me->_state = STATE_BfbIDMTCurve_SAFE;
				me->_trigger = true;
			};
			break;
		case STATE_BfbIDMTCurve_UNSAFE:
			if(me->inputEvents.event.tick && (me->i < me->i_set)) {
				
				me->_state = STATE_BfbIDMTCurve_SAFE;
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
		case STATE_BfbIDMTCurve_INIT:
			#ifdef PRINT_STATES
				printf("BfbIDMTCurve: [Entered State INIT]\n");
			#endif
			
			break;
		case STATE_BfbIDMTCurve_SAFE:
			#ifdef PRINT_STATES
				printf("BfbIDMTCurve: [Entered State SAFE]\n");
			#endif
			BfbIDMTCurve_ResetCnt(me);
			
			break;
		case STATE_BfbIDMTCurve_COUNT:
			#ifdef PRINT_STATES
				printf("BfbIDMTCurve: [Entered State COUNT]\n");
			#endif
			BfbIDMTCurve_UpdateCnt(me);
			BfbIDMTCurve_UpdateMax(me);
			
			break;
		case STATE_BfbIDMTCurve_UNSAFE:
			#ifdef PRINT_STATES
				printf("BfbIDMTCurve: [Entered State UNSAFE]\n");
			#endif
			me->outputEvents.event.unsafe = 1;
			
			break;
		
		default: 
			break;
		}
	}

	me->_trigger = false;

	

	//Ensure input events are cleared
	me->inputEvents.event.tick = 0;
	me->inputEvents.event.i_measured = 0;
	me->inputEvents.event.i_set_change = 0;
	
}


