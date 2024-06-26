// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for ArgoTx
#include "FB_ArgoTx.h"


/* ArgoTx_preinit() is required to be called to 
 * initialise an instance of ArgoTx. 
 * It sets all I/O values to zero.
 */
int ArgoTx_preinit(ArgoTx_t _SPM *me) {
	//if there are input events, reset them
	me->inputEvents.events[0] = 0;
	
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//if there are input vars with default values, set them
	me->ChanId = 1;
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_ArgoTx_Start;
	

	return 0;
}

/* ArgoTx_init() is required to be called to 
 * set up an instance of ArgoTx. 
 * It passes around configuration data.
 */
int ArgoTx_init(ArgoTx_t _SPM *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	me->chan = mp_create_qport(me->ChanId, SOURCE, sizeof(INT), 1);
	if(me->chan == NULL) {
		return 1;
	}

	return 0;
}



/* ArgoTx_run() executes a single tick of an
 * instance of ArgoTx according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void ArgoTx_run(ArgoTx_t _SPM *me) {
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	if(me->inputEvents.event.DataPresent) {
		HEX = me->Data;
		*((volatile _SPM INT *)me->chan->write_buf) = me->Data;
		me->Success = mp_nbsend(me->chan);
		me->outputEvents.event.SuccessChanged = 1;
	}

	me->_trigger = false;
}
//no algorithms were present for this function block


