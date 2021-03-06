// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for Ticksource
#include "FB_Ticksource.h"


/* Ticksource_preinit() is required to be called to 
 * initialise an instance of Ticksource. 
 * It sets all I/O values to zero.
 */
int Ticksource_preinit(Ticksource_t  *me) {
	

	
	//reset the output events
	me->outputEvents.event.TickIc = 0;
	//reset the output events
	me->outputEvents.event.TickGen = 0;
	//reset the output events
	me->outputEvents.event.TickLoad = 0;
	//reset the output events
	me->outputEvents.event.TickTie = 0;
	//reset the output events
	me->outputEvents.event.TickPrint = 0;
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	me->_state = STATE_Ticksource_reset;
	me->_trigger = true;
	
	
	

	

	return 0;
}

/* Ticksource_init() is required to be called to 
 * set up an instance of Ticksource. 
 * It passes around configuration data.
 */
int Ticksource_init(Ticksource_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//no algorithms were present for this function block


/* Ticksource_run() executes a single tick of an
 * instance of Ticksource according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will already be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void Ticksource_run(Ticksource_t  *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.TickIc = 0;
	me->outputEvents.event.TickGen = 0;
	me->outputEvents.event.TickLoad = 0;
	me->outputEvents.event.TickTie = 0;
	me->outputEvents.event.TickPrint = 0;
	
	
	

	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_Ticksource_reset:
			if(true) {
				
				me->_state = STATE_Ticksource_update;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_print:
			if(true) {
				
				me->_state = STATE_Ticksource_tie;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_tie:
			if(true) {
				
				me->_state = STATE_Ticksource_ic;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_ic:
			if(true) {
				
				me->_state = STATE_Ticksource_load;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_load:
			if(true) {
				
				me->_state = STATE_Ticksource_gen;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_gen:
			if(true) {
				
				me->_state = STATE_Ticksource_print;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_update:
			if(true) {
				
				me->_state = STATE_Ticksource_update_load;
				me->_trigger = true;
			};
			break;
		case STATE_Ticksource_update_load:
			if(true) {
				
				me->_state = STATE_Ticksource_update;
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
		case STATE_Ticksource_reset:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State reset]\n");
			#endif
			
			break;
		case STATE_Ticksource_print:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State print]\n");
			#endif
			me->outputEvents.event.TickPrint = 1;
			
			break;
		case STATE_Ticksource_tie:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State tie]\n");
			#endif
			me->outputEvents.event.TickTie = 1;
			
			break;
		case STATE_Ticksource_ic:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State ic]\n");
			#endif
			me->outputEvents.event.TickIc = 1;
			
			break;
		case STATE_Ticksource_load:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State load]\n");
			#endif
			me->outputEvents.event.TickLoad = 1;
			
			break;
		case STATE_Ticksource_gen:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State gen]\n");
			#endif
			me->outputEvents.event.TickGen = 1;
			
			break;
		case STATE_Ticksource_update:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State update]\n");
			#endif
			me->outputEvents.event.TickGen = 1;
			me->outputEvents.event.TickTie = 1;
			me->outputEvents.event.TickIc = 1;
			me->outputEvents.event.TickPrint = 1;
			
			break;
		case STATE_Ticksource_update_load:
			#ifdef PRINT_STATES
				printf("Ticksource: [Entered State update_load]\n");
			#endif
			me->outputEvents.event.TickLoad = 1;
			
			break;
		
		default: 
			break;
		}
	}

	me->_trigger = false;

	

	//Ensure input events are cleared
	
}


