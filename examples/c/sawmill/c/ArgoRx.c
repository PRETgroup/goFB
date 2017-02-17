// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for ArgoRx
#include "ArgoRx.h"


/* ArgoRx_preinit() is required to be called to 
 * initialise an instance of ArgoRx. 
 * It sets all I/O values to zero.
 */
int ArgoRx_preinit(ArgoRx_t *me) {
	//if there are input events, reset them

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
	me->_state = STATE_ArgoRx_Start;
	
	return 0;
}

/* ArgoRx_init() is required to be called to 
 * set up an instance of ArgoRx. 
 * It passes around configuration data.
 */
int ArgoRx_init(ArgoRx_t *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	
	me->chan = mp_create_qport(me->ChanId, SINK, sizeof(INT), 1);
	if(me->chan == NULL) {
		return 1;
	} 
	return 0;
}



/* ArgoRx_run() executes a single tick of an
 * instance of ArgoRx according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void ArgoRx_run(ArgoRx_t *me) {
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	if(me->needToAck == true) {
		if(!mp_nback(me->chan)) {
			return;
		}	
	}
	me->needToAck = false;

	if(mp_nbrecv(me->chan)) {
		me->needToAck = true;
		me->Data = *((volatile INT _SPM*)me->chan->read_buf);
		
		//printf("chan %i recieved %i\n", me->ChanId, me->Data);

		if(mp_nback(me->chan)) {
			me->needToAck = false;
		}
		
		me->outputEvents.event.DataPresent = 1;
	}
}
//no algorithms were present for this function block


