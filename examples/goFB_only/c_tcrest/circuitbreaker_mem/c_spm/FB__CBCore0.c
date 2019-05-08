// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Composite Function Block for _CBCore0
#include "FB__CBCore0.h"

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


/* _CBCore0_preinit() is required to be called to 
 * initialise an instance of _CBCore0. 
 * It sets all I/O values to zero.
 */
int _CBCore0_preinit(_CBCore0_t  _SPM *me) {
	

	
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	
	
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(ArgoRx_preinit(&me->cb3rx) != 0) {
		return 1;
	}
	
	if(ArgoRx_preinit(&me->cb2rx) != 0) {
		return 1;
	}
	
	if(ArgoRx_preinit(&me->cb1rx) != 0) {
		return 1;
	}
	
	if(SifbCBPrintStatus_preinit(&me->Reference) != 0) {
		return 1;
	}
	
	
	

	
	
	

	

	return 0;
}

/* _CBCore0_init() is required to be called to 
 * set up an instance of _CBCore0. 
 * It passes around configuration data.
 */
int _CBCore0_init(_CBCore0_t  _SPM *me) {
	//pass in any parameters on this level
	me->cb3rx.ChanId = 3;
	me->cb2rx.ChanId = 2;
	me->cb1rx.ChanId = 1;
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	//sync config for cb3rx (of Type ArgoRx) 
			//sync config for cb2rx (of Type ArgoRx) 
			//sync config for cb1rx (of Type ArgoRx) 
			//sync config for Reference (of Type SifbCBPrintStatus) 
			me->Reference.St1 = me->cb1rx.Data;
							me->Reference.St2 = me->cb2rx.Data;
							me->Reference.St3 = me->cb3rx.Data;
							

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(ArgoRx_init(&me->cb3rx) != 0) {
		return 1;
	}
	if(ArgoRx_init(&me->cb2rx) != 0) {
		return 1;
	}
	if(ArgoRx_init(&me->cb1rx) != 0) {
		return 1;
	}
	if(SifbCBPrintStatus_init(&me->Reference) != 0) {
		return 1;
	}
	
	

	return 0;
}



/* _CBCore0_syncOutputEvents() synchronises the output events of an
 * instance of _CBCore0 as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore0_syncOutputEvents(_CBCore0_t  _SPM *me) {
	//first, for all cfb children, call this same function
	//then, for all connections that are connected to an output on the parent, run their run their copy
	}

/* _CBCore0_syncInputEvents() synchronises the input events of an
 * instance of _CBCore0 as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore0_syncInputEvents(_CBCore0_t  _SPM *me) {
	//first, we explicitly synchronise the children
	
	me->Reference.inputEvents.event.StatusUpdate = me->cb3rx.outputEvents.event.DataPresent || me->cb2rx.outputEvents.event.DataPresent || me->cb1rx.outputEvents.event.DataPresent; 
	

	//then, call this same function on all cfb children
	
}

/* _CBCore0_syncOutputData() synchronises the output data connections of an
 * instance of _CBCore0 as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore0_syncOutputData(_CBCore0_t  _SPM *me) {
	//for all composite function block children, call this same function
	
	
	//for data that is sent from child to this CFB (me), always copy (event controlled copies will be resolved at the next level up) //TODO: arrays!?
	
	
}

/* _CBCore0_syncInputData() synchronises the input data connections of an
 * instance of _CBCore0 as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore0_syncInputData(_CBCore0_t  _SPM *me) {
	//for all basic function block children, perform their synchronisations explicitly
	
	//sync for cb3rx (of type ArgoRx) which is a BFB
	
	
	//sync for cb2rx (of type ArgoRx) which is a BFB
	
	
	//sync for cb1rx (of type ArgoRx) which is a BFB
	
	
	//sync for Reference (of type SifbCBPrintStatus) which is a BFB
	
	if(me->Reference.inputEvents.event.StatusUpdate == 1) {
		me->Reference.St1 = me->cb1rx.Data;
		me->Reference.St2 = me->cb2rx.Data;
		me->Reference.St3 = me->cb3rx.Data;
		
	} 
	
	
	//for all composite function block children, call this same function
	
	
}


/* _CBCore0_run() executes a single tick of an
 * instance of _CBCore0 according to synchronise semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * is done using the _syncX functions at this (and any higher) level.
 */
void _CBCore0_run(_CBCore0_t  _SPM *me) {
	
	
	ArgoRx_run(&me->cb3rx);
	
	ArgoRx_run(&me->cb2rx);
	
	ArgoRx_run(&me->cb1rx);
	
	SifbCBPrintStatus_run(&me->Reference);
	
}





