// This file is generated by FBC.

#include "ArgoRx.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void ArgoRxinit(ArgoRx* me)
{
    me->_entered = false;
    me->_output.events = 0;
    me->chan = mp_create_qport(me->ChanId, SINK, sizeof(INT), 1);
	if(me->chan == NULL) {
		while(1);
	} 
}

/* Function block execution function */
void ArgoRxrun(ArgoRx* me)
{
    me->_output.events = 0;

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
		me->_Data = me->Data;
		me->_output.event.DataPresent = 1;
	}

}

