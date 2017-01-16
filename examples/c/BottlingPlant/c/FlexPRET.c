// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Composite Function Block for FlexPRET
#include "FlexPRET.h"

//When running a composite block, note that you would call the functions in this order
//_init(); 
//do {
//_syncEvents();
//_syncData();
//_run();
//} loop;


/* FlexPRET_preinit() is required to be called to 
 * initialise an instance of FlexPRET. 
 * It sets all I/O values to zero.
 */
int FlexPRET_preinit(struct FlexPRET *me) {
	//if there are input events, reset them
	
	//if there are output events, reset them
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(IOManager_preinit(&me->IO) != 0) {
		return 1;
	}
	if(CanisterCounter_preinit(&me->CCounter) != 0) {
		return 1;
	}
	if(DoorController_preinit(&me->Door) != 0) {
		return 1;
	}
	if(ConveyorController_preinit(&me->Conveyor) != 0) {
		return 1;
	}
	if(RejectArmController_preinit(&me->RejectArm) != 0) {
		return 1;
	}
	if(InjectorPumpsController_preinit(&me->Pumps) != 0) {
		return 1;
	}
	if(InjectorMotorController_preinit(&me->Motor) != 0) {
		return 1;
	}
	
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	

	return 0;
}

/* FlexPRET_init() is required to be called to 
 * set up an instance of FlexPRET. 
 * It passes around configuration data.
 */
int FlexPRET_init(struct FlexPRET *me) {
	//pass in any parameters on this level
	
	
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	me->Door.EmergencyStop = me->IO.EmergencyStop;
	me->Conveyor.EmergencyStop = me->IO.EmergencyStop;
	me->Motor.EmergencyStop = me->IO.EmergencyStop;
	me->Pumps.EmergencyStop = me->IO.EmergencyStop;
	me->Pumps.CanisterPressure = me->IO.CanisterPressure;
	me->Pumps.FillContentsAvailable = me->IO.FillContentsAvailable;
	me->CCounter.DoorSiteLaser = me->IO.DoorSiteLaser;
	me->Conveyor.InjectSiteLaser = me->IO.InjectSiteLaser;
	me->RejectArm.RejectSiteLaser = me->IO.RejectSiteLaser;
	me->CCounter.RejectBinLaser = me->IO.RejectBinLaser;
	me->CCounter.AcceptBinLaser = me->IO.AcceptBinLaser;
	me->IO.CanisterCount = me->CCounter.CanisterCount;
	me->IO.ConveyorSpeed = me->Conveyor.ConveyorSpeed;
	me->IO.InjectorContentsValveOpen = me->Pumps.InjectorContentsValveOpen;
	me->IO.InjectorVacuumRun = me->Pumps.InjectorVacuumRun;
	me->IO.InjectorPressurePumpRun = me->Pumps.InjectorPressurePumpRun;
	me->IO.FillContents = me->Pumps.FillContents;
	me->IO.InjectorPosition = me->Motor.InjectorPosition;
	

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	if(IOManager_init(&me->IO) != 0) {
		return 1;
	}
	if(CanisterCounter_init(&me->CCounter) != 0) {
		return 1;
	}
	if(DoorController_init(&me->Door) != 0) {
		return 1;
	}
	if(ConveyorController_init(&me->Conveyor) != 0) {
		return 1;
	}
	if(RejectArmController_init(&me->RejectArm) != 0) {
		return 1;
	}
	if(InjectorPumpsController_init(&me->Pumps) != 0) {
		return 1;
	}
	if(InjectorMotorController_init(&me->Motor) != 0) {
		return 1;
	}
	
	

	return 0;
}



/* FlexPRET_syncEvents() synchronises the events of an
 * instance of FlexPRET as required by synchronous semantics.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void FlexPRET_syncEvents(struct FlexPRET *me) {
	//for all composite function block children, call this same function
	
	//for all basic function block children, perform their synchronisations explicitly
	//events are always copied
	me->Motor.inputEvents.event.InjectorArmFinishedMovement = me->IO.outputEvents.event.InjectorArmFinishMovement;
	me->Door.inputEvents.event.EmergencyStopChanged = me->IO.outputEvents.event.EmergencyStopChanged;
	me->Conveyor.inputEvents.event.EmergencyStopChanged = me->IO.outputEvents.event.EmergencyStopChanged;
	me->Motor.inputEvents.event.EmergencyStopChanged = me->IO.outputEvents.event.EmergencyStopChanged;
	me->Pumps.inputEvents.event.EmergencyStopChanged = me->IO.outputEvents.event.EmergencyStopChanged;
	me->Pumps.inputEvents.event.CanisterPressureChanged = me->IO.outputEvents.event.CanisterPressureChanged;
	me->Pumps.inputEvents.event.FillContentsAvailableChanged = me->IO.outputEvents.event.FillContentsAvailableChanged;
	me->CCounter.inputEvents.event.LasersChanged = me->IO.outputEvents.event.LasersChanged;
	me->RejectArm.inputEvents.event.LasersChanged = me->IO.outputEvents.event.LasersChanged;
	me->Conveyor.inputEvents.event.LasersChanged = me->IO.outputEvents.event.LasersChanged;
	me->Door.inputEvents.event.ReleaseDoorOverride = me->IO.outputEvents.event.DoorOverride;
	me->Pumps.inputEvents.event.VacuumTimerElapsed = me->IO.outputEvents.event.VacuumTimerElapsed;
	me->IO.inputEvents.event.CanisterCountChanged = me->CCounter.outputEvents.event.CanisterCountChanged;
	me->IO.inputEvents.event.DoorReleaseCanister = me->Door.outputEvents.event.DoorReleaseCanister;
	me->IO.inputEvents.event.ConveyorChanged = me->Conveyor.outputEvents.event.ConveyorChanged;
	me->Motor.inputEvents.event.ConveyorStoppedForInject = me->Conveyor.outputEvents.event.ConveyorStoppedForInject;
	me->IO.inputEvents.event.GoRejectArm = me->RejectArm.outputEvents.event.GoRejectArm;
	me->Motor.inputEvents.event.PumpFinished = me->Pumps.outputEvents.event.PumpFinished;
	me->RejectArm.inputEvents.event.RejectCanister = me->Pumps.outputEvents.event.RejectCanister;
	me->IO.inputEvents.event.InjectorControlsChanged = me->Pumps.outputEvents.event.InjectorControlsChanged;
	me->IO.inputEvents.event.FillContentsChanged = me->Pumps.outputEvents.event.FillContentsChanged;
	me->IO.inputEvents.event.StartVacuumTimer = me->Pumps.outputEvents.event.StartVacuumTimer;
	me->Pumps.inputEvents.event.StartPump = me->Motor.outputEvents.event.StartPump;
	me->Door.inputEvents.event.BottlingDone = me->Motor.outputEvents.event.InjectDone;
	me->Conveyor.inputEvents.event.InjectDone = me->Motor.outputEvents.event.InjectDone;
	me->IO.inputEvents.event.InjectDone = me->Motor.outputEvents.event.InjectDone;
	me->IO.inputEvents.event.InjectorPositionChanged = me->Motor.outputEvents.event.InjectorPositionChanged;
	
}

/* FlexPRET_syncData() synchronises the data connections of an
 * instance of FlexPRET as required by synchronous semantics.
 * It does the checking to ensure that only connections which have had their
 * associated event fire are updated.
 * Notice that it does NOT perform any computation - this occurs in the
 * _run function.
 */
void FlexPRET_syncData(struct FlexPRET *me) {
	//for all composite function block children, call this same function
	
	//for all basic function block children, perform their synchronisations explicitly
	//Data is sometimes copied
	
	//sync for IO (of type IOManager) which is a BFB
	
	if(me->IO.inputEvents.event.ConveyorChanged == 1) { 
		me->IO.ConveyorSpeed = me->Conveyor.ConveyorSpeed;
	} 
	if(me->IO.inputEvents.event.InjectorPositionChanged == 1) { 
		me->IO.InjectorPosition = me->Motor.InjectorPosition;
	} 
	if(me->IO.inputEvents.event.InjectorControlsChanged == 1) { 
		me->IO.InjectorContentsValveOpen = me->Pumps.InjectorContentsValveOpen;
		me->IO.InjectorVacuumRun = me->Pumps.InjectorVacuumRun;
		me->IO.InjectorPressurePumpRun = me->Pumps.InjectorPressurePumpRun;
	} 
	if(me->IO.inputEvents.event.FillContentsChanged == 1) { 
		me->IO.FillContents = me->Pumps.FillContents;
	} 
	if(me->IO.inputEvents.event.CanisterCountChanged == 1) { 
		me->IO.CanisterCount = me->CCounter.CanisterCount;
	} 
	
	//sync for CCounter (of type CanisterCounter) which is a BFB
	
	if(me->CCounter.inputEvents.event.LasersChanged == 1) { 
		me->CCounter.DoorSiteLaser = me->IO.DoorSiteLaser;
		me->CCounter.RejectBinLaser = me->IO.RejectBinLaser;
		me->CCounter.AcceptBinLaser = me->IO.AcceptBinLaser;
	} 
	
	//sync for Door (of type DoorController) which is a BFB
	
	if(me->Door.inputEvents.event.EmergencyStopChanged == 1) { 
		me->Door.EmergencyStop = me->IO.EmergencyStop;
	} 
	
	//sync for Conveyor (of type ConveyorController) which is a BFB
	
	if(me->Conveyor.inputEvents.event.EmergencyStopChanged == 1) { 
		me->Conveyor.EmergencyStop = me->IO.EmergencyStop;
	} 
	if(me->Conveyor.inputEvents.event.LasersChanged == 1) { 
		me->Conveyor.InjectSiteLaser = me->IO.InjectSiteLaser;
	} 
	
	//sync for RejectArm (of type RejectArmController) which is a BFB
	
	if(me->RejectArm.inputEvents.event.LasersChanged == 1) { 
		me->RejectArm.RejectSiteLaser = me->IO.RejectSiteLaser;
	} 
	
	//sync for Pumps (of type InjectorPumpsController) which is a BFB
	
	if(me->Pumps.inputEvents.event.EmergencyStopChanged == 1) { 
		me->Pumps.EmergencyStop = me->IO.EmergencyStop;
	} 
	if(me->Pumps.inputEvents.event.CanisterPressureChanged == 1) { 
		me->Pumps.CanisterPressure = me->IO.CanisterPressure;
	} 
	if(me->Pumps.inputEvents.event.FillContentsAvailableChanged == 1) { 
		me->Pumps.FillContentsAvailable = me->IO.FillContentsAvailable;
	} 
	
	//sync for Motor (of type InjectorMotorController) which is a BFB
	
	if(me->Motor.inputEvents.event.EmergencyStopChanged == 1) { 
		me->Motor.EmergencyStop = me->IO.EmergencyStop;
	} 
	

}


/* FlexPRET_run() executes a single tick of an
 * instance of FlexPRET according to synchronise semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * is done using the _syncX functions at this (and any higher) level.
 */
void FlexPRET_run(struct FlexPRET *me) {
	IOManager_run(&me->IO);
	CanisterCounter_run(&me->CCounter);
	DoorController_run(&me->Door);
	ConveyorController_run(&me->Conveyor);
	RejectArmController_run(&me->RejectArm);
	InjectorPumpsController_run(&me->Pumps);
	InjectorMotorController_run(&me->Motor);
	
}

