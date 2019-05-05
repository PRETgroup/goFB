// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for SifbAmmeter
#include "FB_SifbAmmeter.h"


/* SifbAmmeter_preinit() is required to be called to 
 _SPM * initialise an instance of SifbAmmeter. 
 _SPM * It sets all I/O values to zero.
 _SPM */
int SifbAmmeter_preinit(SifbAmmeter_t  _SPM *me) {
	

	
	//reset the output events
	me->outputEvents.event.i_measured = 0;
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	//set any internal vars with default values
	
	
	
	
	

	
	
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	// me->_state = STATE_SifbAmmeter_Start;
	// me->_trigger = true;
	
	
	

	

	return 0;
}

/* SifbAmmeter_init() is required to be called to 
 _SPM * set up an instance of SifbAmmeter. 
 _SPM * It passes around configuration data.
 _SPM */
int SifbAmmeter_init(SifbAmmeter_t  _SPM *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	

	return 0;
}



//algorithms

void SifbAmmeter_update_amms(SifbAmmeter_t  _SPM *me) {
//PROVIDED CODE: this algorithm was provided in an algorithm's text field
int sw_ammh = ((SWITCHES & 0b1000) != 0);

if(sw_ammh == 1 && (int)(me->i) != 300) {
	//switch is pressed
	me->i = 300.0;
	me->outputEvents.event.i_measured = 1;
}
if(sw_ammh == 0 && (int)(me->i) != 5) {
	//switch is pressed
	me->i = 5.0;
	me->outputEvents.event.i_measured = 1;
}
}



/* SifbAmmeter_run() executes a single tick of an
 _SPM * instance of SifbAmmeter according to synchronous semantics.
 _SPM * Notice that it does NOT perform any I/O - synchronisation
 _SPM * will need to be done in the parent.
 _SPM * Also note that on the first run of this function, trigger will already be set
 _SPM * to true, meaning that on the very first run no next state logic will occur.
 _SPM */
void SifbAmmeter_run(SifbAmmeter_t  _SPM *me) {
	//if there are output events, reset them
	
	me->outputEvents.event.i_measured = 0;
	
	
	

	
	// //next state logic
	// if(me->_trigger == false) {
	// 	switch(me->_state) {
	// 	case STATE_SifbAmmeter_Start:
	// 		if(true) {
				
	// 			me->_state = STATE_SifbAmmeter_Update;
	// 			me->_trigger = true;
	// 		};
	// 		break;
	// 	case STATE_SifbAmmeter_Update:
	// 		if(true) {
				
	// 			me->_state = STATE_SifbAmmeter_Update;
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
	// 	case STATE_SifbAmmeter_Start:
	// 		#ifdef PRINT_STATES
	// 			printf("SifbAmmeter: [Entered State Start]\n");
	// 		#endif
			
	// 		break;
	// 	case STATE_SifbAmmeter_Update:
	// 		#ifdef PRINT_STATES
	// 			printf("SifbAmmeter: [Entered State Update]\n");
	// 		#endif
			SifbAmmeter_update_amms(me);
			
	// 		break;
		
	// 	default: 
	// 		break;
	// 	}
	// }

	// me->_trigger = false;

	

	//Ensure input events are cleared
	
}


