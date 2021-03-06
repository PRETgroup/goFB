// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Composite Function Block for top
#include "FB_top.h"

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


/* top_preinit() is required to be called to 
 * initialise an instance of top. 
 * It sets all I/O values to zero.
 */
int top_preinit(top_t  *me) {
	

	
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	
	
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(load_preinit(&me->load) != 0) {
		return 1;
	}
	
	if(controller_preinit(&me->contr) != 0) {
		return 1;
	}
	
	
	

	
	
	

	

	return 0;
}

/* top_init() is required to be called to 
 * set up an instance of top. 
 * It passes around configuration data.
 */
int top_init(top_t  *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	//sync config for load (of Type load) 
			//sync config for contr (of Type controller) 
			

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(load_init(&me->load) != 0) {
		return 1;
	}
	if(controller_init(&me->contr) != 0) {
		return 1;
	}
	
	

	return 0;
}



/* top_syncOutputEvents() synchronises the output events of an
 * instance of top as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void top_syncOutputEvents(top_t  *me) {
	//first, for all cfb children, call this same function
	//then, for all connections that are connected to an output on the parent, run their run their copy
	}

/* top_syncInputEvents() synchronises the input events of an
 * instance of top as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void top_syncInputEvents(top_t  *me) {
	//first, we explicitly synchronise the children
	
	me->load.inputEvents.event.on = me->contr.outputEvents.event.on; 
	
	me->load.inputEvents.event.off = me->contr.outputEvents.event.off; 
	
	me->load.inputEvents.event.fault = me->contr.outputEvents.event.fault; 
	

	//then, call this same function on all cfb children
	
}

/* top_syncOutputData() synchronises the output data connections of an
 * instance of top as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void top_syncOutputData(top_t  *me) {
	//for all composite function block children, call this same function
	
	
	//for data that is sent from child to this CFB (me), always copy (event controlled copies will be resolved at the next level up) //TODO: arrays!?
	
	
}

/* top_syncInputData() synchronises the input data connections of an
 * instance of top as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void top_syncInputData(top_t  *me) {
	//for all basic function block children, perform their synchronisations explicitly
	
	//sync for load (of type load) which is a BFB
	
	
	//sync for contr (of type controller) which is a BFB
	
	
	
	//for all composite function block children, call this same function
	
	
}


/* top_run() executes a single tick of an
 * instance of top according to synchronise semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * is done using the _syncX functions at this (and any higher) level.
 */
void top_run(top_t  *me) {
	
	
	load_run(&me->load);
	
	controller_run(&me->contr);
	
}





