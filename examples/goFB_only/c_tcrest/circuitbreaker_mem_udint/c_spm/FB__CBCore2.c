// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Composite Function Block for _CBCore2
#include "FB__CBCore2.h"

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


/* _CBCore2_preinit() is required to be called to 
 * initialise an instance of _CBCore2. 
 * It sets all I/O values to zero.
 */
int _CBCore2_preinit(_CBCore2_t _SPM  *me) {
	

	
	
	//set any input vars with default values
	
	//set any output vars with default values
	
	
	
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(SifbTimer_preinit(&me->timer) != 0) {
		return 1;
	}
	
	if(ArgoTx_preinit(&me->tx) != 0) {
		return 1;
	}
	
	if(SifbManagementControls_preinit(&me->hmi) != 0) {
		return 1;
	}
	
	if(SifbIntLed_preinit(&me->led) != 0) {
		return 1;
	}
	
	if(CfbBreakerController_preinit(&me->cb) != 0) {
		return 1;
	}
	
	if(SifbAmmeter_preinit(&me->amm) != 0) {
		return 1;
	}
	
	if(SawmillMessageHandler_preinit(&me->msgh) != 0) {
		return 1;
	}
	
	
	

	
	
	

	

	return 0;
}

/* _CBCore2_init() is required to be called to 
 * set up an instance of _CBCore2. 
 * It passes around configuration data.
 */
int _CBCore2_init(_CBCore2_t _SPM  *me) {
	//pass in any parameters on this level
	me->tx.ChanId = 2;
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	
	
	//sync config for timer (of Type SifbTimer) 
			//sync config for tx (of Type ArgoTx) 
			me->tx.Data = me->msgh.TxData;
							//sync config for hmi (of Type SifbManagementControls) 
			//sync config for led (of Type SifbIntLed) 
			me->led.i = me->cb.b;
							//sync config for cb (of Type CfbBreakerController) 
			me->cb.i = me->amm.i;
							me->cb.i_set = me->hmi.i_set;
							//sync config for amm (of Type SifbAmmeter) 
			//sync config for msgh (of Type SawmillMessageHandler) 
			me->msgh.Message = me->cb.b;
							me->msgh.TxSuccess = me->tx.Success;
							

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(SifbTimer_init(&me->timer) != 0) {
		return 1;
	}
	if(ArgoTx_init(&me->tx) != 0) {
		return 1;
	}
	if(SifbManagementControls_init(&me->hmi) != 0) {
		return 1;
	}
	if(SifbIntLed_init(&me->led) != 0) {
		return 1;
	}
	if(CfbBreakerController_init(&me->cb) != 0) {
		return 1;
	}
	if(SifbAmmeter_init(&me->amm) != 0) {
		return 1;
	}
	if(SawmillMessageHandler_init(&me->msgh) != 0) {
		return 1;
	}
	
	

	return 0;
}



/* _CBCore2_syncOutputEvents() synchronises the output events of an
 * instance of _CBCore2 as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore2_syncOutputEvents(_CBCore2_t _SPM  *me) {
	//first, for all cfb children, call this same function
	
	CfbBreakerController_syncOutputEvents(&me->cb);//sync for cb (of type CfbBreakerController) which is a CFB//then, for all connections that are connected to an output on the parent, run their run their copy
	}

/* _CBCore2_syncInputEvents() synchronises the input events of an
 * instance of _CBCore2 as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore2_syncInputEvents(_CBCore2_t _SPM  *me) {
	//first, we explicitly synchronise the children
	
	me->tx.inputEvents.event.DataPresent = me->msgh.outputEvents.event.TxDataPresent; 
	
	me->led.inputEvents.event.i_change = me->cb.outputEvents.event.b_change; 
	
	me->cb.inputEvents.event.tick = me->timer.outputEvents.event.Tick; 
	
	me->cb.inputEvents.event.i_measured = me->amm.outputEvents.event.i_measured; 
	
	me->cb.inputEvents.event.i_set_change = me->hmi.outputEvents.event.i_set_change; 
	
	me->cb.inputEvents.event.brk = me->hmi.outputEvents.event.brk; 
	
	me->cb.inputEvents.event.rst = me->hmi.outputEvents.event.rst; 
	
	me->msgh.inputEvents.event.MessageChanged = me->cb.outputEvents.event.b_change; 
	
	me->msgh.inputEvents.event.TxSuccessChanged = me->tx.outputEvents.event.SuccessChanged; 
	

	//then, call this same function on all cfb children
	
	CfbBreakerController_syncInputEvents(&me->cb);//sync for cb (of type CfbBreakerController) which is a CFB
	
}

/* _CBCore2_syncOutputData() synchronises the output data connections of an
 * instance of _CBCore2 as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore2_syncOutputData(_CBCore2_t _SPM  *me) {
	//for all composite function block children, call this same function
	//sync for cb (of type CfbBreakerController) which is a CFB
	CfbBreakerController_syncOutputData(&me->cb);
	
	//for data that is sent from child to this CFB (me), always copy (event controlled copies will be resolved at the next level up) //TODO: arrays!?
	
	
}

/* _CBCore2_syncInputData() synchronises the input data connections of an
 * instance of _CBCore2 as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void _CBCore2_syncInputData(_CBCore2_t _SPM  *me) {
	//for all basic function block children, perform their synchronisations explicitly
	
	//sync for timer (of type SifbTimer) which is a BFB
	
	
	//sync for tx (of type ArgoTx) which is a BFB
	
	if(me->tx.inputEvents.event.DataPresent == 1) {
		me->tx.Data = me->msgh.TxData;
		
	} 
	
	//sync for hmi (of type SifbManagementControls) which is a BFB
	
	
	//sync for led (of type SifbIntLed) which is a BFB
	
	if(me->led.inputEvents.event.i_change == 1) {
		me->led.i = me->cb.b;
		
	} 
	//sync for cb (of Type CfbBreakerController) which is a CFB
	
	
		me->cb.i = me->amm.i;
	
	
		me->cb.i_set = me->hmi.i_set;
	
	//sync for amm (of type SifbAmmeter) which is a BFB
	
	
	//sync for msgh (of type SawmillMessageHandler) which is a BFB
	
	if(me->msgh.inputEvents.event.MessageChanged == 1) {
		me->msgh.Message = me->cb.b;
		
	} 
	if(me->msgh.inputEvents.event.TxSuccessChanged == 1) {
		me->msgh.TxSuccess = me->tx.Success;
		
	} 
	
	
	//for all composite function block children, call this same function
	//sync for cb (of type CfbBreakerController) which is a CFB
	CfbBreakerController_syncInputData(&me->cb);
	
}


/* _CBCore2_run() executes a single tick of an
 * instance of _CBCore2 according to synchronise semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * is done using the _syncX functions at this (and any higher) level.
 */
void _CBCore2_run(_CBCore2_t  _SPM *me) {
	
	
	SifbTimer_run(&me->timer);
	
	ArgoTx_run(&me->tx);
	
	SifbManagementControls_run(&me->hmi);
	
	SifbIntLed_run(&me->led);
	
	CfbBreakerController_run(&me->cb);
	
	SifbAmmeter_run(&me->amm);
	
	SawmillMessageHandler_run(&me->msgh);
	
}





