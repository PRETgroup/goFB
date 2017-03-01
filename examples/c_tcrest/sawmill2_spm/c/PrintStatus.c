// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for PrintStatus
#include "PrintStatus.h"


/* PrintStatus_preinit() is required to be called to 
 * initialise an instance of PrintStatus. 
 * It sets all I/O values to zero.
 */
int PrintStatus_preinit(PrintStatus_t _SPM *me) {
	//if there are input events, reset them
	me->inputEvents.event.StatusUpdate = 0;
	
	//if there are output events, reset them
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_PrintStatus_Start;
	

	return 0;
}

/* PrintStatus_init() is required to be called to 
 * set up an instance of PrintStatus. 
 * It passes around configuration data.
 */
int PrintStatus_init(PrintStatus_t _SPM *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



/* PrintStatus_run() executes a single tick of an
 * instance of PrintStatus according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void PrintStatus_run(PrintStatus_t _SPM *me) {
	//if there are output events, reset them
	
	

	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_PrintStatus_Start:
			if(true) {
				me->_state = STATE_PrintStatus_Run;
				me->_trigger = true;
				
			};
			break;
		case STATE_PrintStatus_Run:
			if(true) {
				me->_state = STATE_PrintStatus_Run;
				me->_trigger = true;
				
			};
			break;
		
		}
	}

	//state output logic
	if(me->_trigger == true) {
		switch(me->_state) {
		case STATE_PrintStatus_Start:
			break;

		case STATE_PrintStatus_Run:
			PrintStatus_PrintService(me);
			break;

		
		}
	}

	me->_trigger = false;
}
//algorithms

void PrintStatus_PrintService(PrintStatus_t _SPM *me) {
// OFF 0
// RUNNING 1
// WEIGHT -1
// LASER -2
// STALL -3
// SPEED -4

if(me->inputEvents.event.StatusUpdate) {
	printf("Saw 1: ");
	switch(me->Saw1Status) {
		case 0:
			printf("OFF          ");
			break;
		case 1:
			printf("RUNNING      ");
			break;
		case -1:
			printf("ERR:SAWDUST  ");
			break;
		case -2:
			printf("ERR:PROXIMITY");
			break;
		case -3:
			printf("ERR:STALL    ");
			break;
		case -4:
			printf("ERR:JITTER   ");
			break;
		default:
			printf("ERR:UNKNOWN  ");
	}
	printf("\t");

	printf("Saw 2: ");
	switch(me->Saw2Status) {
		case 0:
			printf("OFF          ");
			break;
		case 1:
			printf("RUNNING      ");
			break;
		case -1:
			printf("ERR:SAWDUST  ");
			break;
		case -2:
			printf("ERR:PROXIMITY");
			break;
		case -3:
			printf("ERR:STALL    ");
			break;
		case -4:
			printf("ERR:JITTER   ");
			break;
		default:
			printf("ERR:UNKNOWN  ");
	}
	printf("\t");

	printf("Saw 3: ");
	switch(me->Saw3Status) {
		case 0:
			printf("OFF          ");
			break;
		case 1:
			printf("RUNNING      ");
			break;
		case -1:
			printf("ERR:SAWDUST  ");
			break;
		case -2:
			printf("ERR:PROXIMITY");
			break;
		case -3:
			printf("ERR:STALL    ");
			break;
		case -4:
			printf("ERR:JITTER   ");
			break;
		default:
			printf("ERR:UNKNOWN  ");
	}
	printf("\n");

}
}



