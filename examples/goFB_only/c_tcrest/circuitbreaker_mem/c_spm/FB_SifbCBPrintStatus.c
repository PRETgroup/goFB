// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for SifbCBPrintStatus
#include "FB_SifbCBPrintStatus.h"


/* SifbCBPrintStatus_preinit() is required to be called to 
 _SPM * initialise an instance of SifbCBPrintStatus. 
 _SPM * It sets all I/O values to zero.
 _SPM */
int SifbCBPrintStatus_preinit(SifbCBPrintStatus_t  _SPM *me) {
	

	//reset the input events
	me->inputEvents.event.StatusUpdate = 0;
	
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	// me->_state = STATE_SifbCBPrintStatus_Start;
	// me->_trigger = true;
	
	
	

	

	return 0;
}

/* SifbCBPrintStatus_init() is required to be called to 
 _SPM * set up an instance of SifbCBPrintStatus. 
 _SPM * It passes around configuration data.
 _SPM */
int SifbCBPrintStatus_init(SifbCBPrintStatus_t  _SPM *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void SifbCBPrintStatus_PrintService(SifbCBPrintStatus_t  _SPM *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
// OFF 0
// RUNNING 1
// WEIGHT -1
// LASER -2
// STALL -3
// SPEED -4

HEX = me->St1 || (me->St2 << 4) || (me->St3 << 4); 
/*
printf("CB 1: ");
switch(me->St1) {
	case 0:
		printf("CLOSED       ");
		break;
	case 1:
		printf("OPEN         ");
		break;
	default:
		printf("UNKNOWN      ");
		break;
}
printf("\t");

printf("CB 2: ");
switch(me->St2) {
	case 0:
		printf("CLOSED       ");
		break;
	case 1:
		printf("OPEN         ");
		break;
	default:
		printf("UNKNOWN      ");
		break;
}
printf("\t");

printf("CB 3: ");
switch(me->St3) {
	case 0:
		printf("CLOSED       ");
		break;
	case 1:
		printf("OPEN         ");
		break;
	default:
		printf("UNKNOWN      ");
		break;
}
printf("\n");
*/

}



/* SifbCBPrintStatus_run() executes a single tick of an
 _SPM * instance of SifbCBPrintStatus according to synchronous semantics.
 _SPM * Notice that it does NOT perform any I/O - synchronisation
 _SPM * will need to be done in the parent.
 _SPM * Also note that on the first run of this function, trigger will already be set
 _SPM * to true, meaning that on the very first run no next state logic will occur.
 _SPM */
void SifbCBPrintStatus_run(SifbCBPrintStatus_t  _SPM *me) {
	//if there are output events, reset them
	
	
	
	

	
	// //next state logic
	// if(me->_trigger == false) {
	// 	switch(me->_state) {
	// 	case STATE_SifbCBPrintStatus_Start:
	// 		if(true) {
				
	// 			me->_state = STATE_SifbCBPrintStatus_Update;
	// 			me->_trigger = true;
	// 		};
	// 		break;
	// 	case STATE_SifbCBPrintStatus_Update:
	// 		if(me->inputEvents.event.StatusUpdate) {
				
	// 			me->_state = STATE_SifbCBPrintStatus_Update;
	// 			me->_trigger = true;
	// 		};
	// 		break;
		
	// 	default: 
	// 		break;
	// 	}
	// }

	// //state output logic
	// if(me->_trigger == true) {
	// 	switch(me->_state) {
	// 	case STATE_SifbCBPrintStatus_Start:
	// 		#ifdef PRINT_STATES
	// 			printf("SifbCBPrintStatus: [Entered State Start]\n");
	// 		#endif
			
	// 		break;
	// 	case STATE_SifbCBPrintStatus_Update:
	// 		#ifdef PRINT_STATES
	// 			printf("SifbCBPrintStatus: [Entered State Update]\n");
	// 		#endif
			SifbCBPrintStatus_PrintService(me);
			
	// 		break;
		
	// 	default: 
	// 		break;
	// 	}
	// }

	// me->_trigger = false;

	

	//Ensure input events are cleared
	me->inputEvents.event.StatusUpdate = 0;
	
}


