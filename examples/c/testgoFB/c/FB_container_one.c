// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Composite Function Block for container_one
#include "FB_container_one.h"

//When running a composite block, note that you would call the functions in this order (and this is very important)
//_preinit(); 
//_init();
//do {
//	_syncOutputEvents();
//	_syncInputEvents();
//	_syncOutputData();
//	_syncInputData();
//	_run();
//} loop;


/* container_one_preinit() is required to be called to 
 * initialise an instance of container_one. 
 * It sets all I/O values to zero.
 */
int container_one_preinit(container_one_t  *me) {
	//if there are input events, reset them
	me->inputEvents.event.DataInChanged = 0;
	
	//if there are output events, reset them
	me->outputEvents.event.DataOutChanged = 0;
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(passforward_preinit(&me->inside) != 0) {
		return 1;
	}
	
	

	

	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	

	return 0;
}

/* container_one_init() is required to be called to 
 * set up an instance of container_one. 
 * It passes around configuration data.
 */
int container_one_init(container_one_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	//sync config for inside (of Type passforward) 
	
	
		me->inside.DataIn = me->DataIn;
	
	
		me->inside.printf_id = me->printf_id;
	

	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(passforward_init(&me->inside) != 0) {
		return 1;
	}
	
	

	return 0;
}



/* container_one_syncOutputEvents() synchronises the output events of an
 * instance of container_one as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void container_one_syncOutputEvents(container_one_t  *me) {
	//first, for all cfb children, call this same function
	
	
	//then, for all connections that are connected to an output on the parent, run their run their copy
	
	me->outputEvents.event.DataOutChanged = me->inside.outputEvents.event.DataOutChanged; 
	
}

/* container_one_syncInputEvents() synchronises the input events of an
 * instance of container_one as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void container_one_syncInputEvents(container_one_t  *me) {
	//first, we explicitly synchronise the children
	
	me->inside.inputEvents.event.DataInChanged = me->inputEvents.event.DataInChanged; 
	

	//then, call this same function on all cfb children
	
}

/* container_one_syncOutputData() synchronises the output data connections of an
 * instance of container_one as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void container_one_syncOutputData(container_one_t  *me) {
	//for all composite function block children, call this same function
	
	
	//for data that is sent from child to this CFB (me), always copy (event controlled copies will be resolved at the next level up) //TODO: arrays!?
	me->DataOut = me->inside.DataOut;
	
	
}

/* container_one_syncInputData() synchronises the input data connections of an
 * instance of container_one as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void container_one_syncInputData(container_one_t  *me) {
	//for all basic function block children, perform their synchronisations explicitly
	
	//sync for inside (of type passforward) which is a BFB
	
	if(me->inside.inputEvents.event.DataInChanged == 1) { 
		me->inside.DataIn = me->DataIn;
	} 
	
	
	//for all composite function block children, call this same function
	
	
}


/* container_one_run() executes a single tick of an
 * instance of container_one according to synchronise semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * is done using the _syncX functions at this (and any higher) level.
 */
void container_one_run(container_one_t  *me) {
	passforward_run(&me->inside);
	
}
