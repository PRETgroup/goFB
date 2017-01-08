// This file has been automatically generated by goFB and should not be edited by hand
// Transpiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the implementation of the Basic Function Block for IOManager
#include "IOManager.h"




/* IOManager_init() is required to be called to 
 * initialise an instance of IOManager. 
 * It sets all I/O values to zero.
 */
void IOManager_init(struct IOManager *me) {
	//if there are input events, reset them
	me->inputEvents.events[0] = 0;
	
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//if there are input vars with default values, set them
	
	//if there are output vars with default values, set them
	
	//if there are internal vars with default values, set them (BFBs only)
	
	//if there are resource vars with default values, set them
	
	//if there are resources with set parameters, set them
	
	//if there are fb children (CFBs only), call this same function on them
	
	//if this is a BFB, set _trigger to be true and start state so that the start state is properly executed
	me->_trigger = true;
	me->_state = STATE_IOManager_Start;
	
}



/* IOManager_run() executes a single tick of an
 * instance of IOManager according to synchronous semantics.
 * Notice that it does NOT perform any I/O - synchronisation
 * will need to be done in the parent.
 * Also note that on the first run of this function, trigger will be set
 * to true, meaning that on the very first run no next state logic will occur.
 */
void IOManager_run(struct IOManager *me) {
	//if there are output events, reset them
	me->outputEvents.events[0] = 0;
	
	//next state logic
	if(me->_trigger == false) {
		switch(me->_state) {
		case STATE_IOManager_Start:
			if(true) {
				me->_state = STATE_IOManager_Start;
				me->_trigger = true;
			};
			break;
		
		}
	}

	//state output logic
	if(me->_trigger == true) {
		switch(me->_state) {
		case STATE_IOManager_Start:
			IOManager_IOAlgorithm(me);
			me->outputEvents.event.EmergencyStopChanged = 1;
			break;

		
		}
	}

	me->_trigger = false;
}

//algorithms

void IOManager_IOAlgorithm(struct IOManager *me) {
#define NUM_BOTTLES 4

static int emergencyStopped = 1;

static int conveyorSpeed = 0;

int i;

static int tickNum = 0;

printf("ATTN: Tick number %i\n", tickNum);
tickNum++;


static int bottlePositions[NUM_BOTTLES] = {0};
static int bottlesActive[NUM_BOTTLES] = {0};
static int nextBottle = 0;

//reset all the things
me->EmergencyStop = 0;
me->CanisterPressure = 255;
me->FillContentsAvailable = 255;
me->DoorSiteLaser = 0;
me->InjectSiteLaser = 0;
me->RejectSiteLaser = 0;
me->RejectBinLaser = 0;
me->AcceptBinLaser = 0;


//printf("=====new tick\n");

//continue progress
if(conveyorSpeed) {
	for(i = 0; i < NUM_BOTTLES; i++) {
		if(bottlesActive[i]) {
			bottlePositions[i] += conveyorSpeed;
			printf("IO: Canister %i moves to %i\n", i, bottlePositions[i]);
			
			if(bottlePositions[i] == 5) {
				printf("IO: Canister %i at 5, triggering InjectSiteLaser\n", i);
				me->outputEvents.event.LasersChanged = 1;
				me->InjectSiteLaser = 1;
			}

			if(bottlePositions[i] == 10) {
				printf("IO: Canister %i at 10, triggering RejectSiteLaser\n", i);
				me->outputEvents.event.LasersChanged = 1;
				me->RejectSiteLaser = 1;
			}

			if(bottlePositions[i] == 20) {
				printf("IO: Canister %i at 20, falls off conveyor, triggering AcceptBinLaser\n", i);
				me->outputEvents.event.LasersChanged = 1;
				me->AcceptBinLaser = 1;
				bottlesActive[i] = 0;
				bottlePositions[i] = 0;
			}

			if(me->inputEvents.event.GoRejectArm && (bottlePositions[i] == 10 || bottlePositions[i] == 11 || bottlePositions[i] == 12)) {
				printf("IO: Go Reject Arm. Canister %i knocked from conveyor.\n", i);
				//progress = 0;
				me->outputEvents.event.LasersChanged = 1;
				me->RejectBinLaser = 1;
				bottlesActive[i] = 0;
				bottlePositions[i] = 0;
			}
		}
	}
}

if(tickNum == 25) {
	printf("Progress at 25, halting\n");
	while(1);
}

if(me->inputEvents.event.InjectDone) {
	printf("IO: Inject done\n");
}



if(emergencyStopped == 1) {
	printf("IO: Releasing emergency stop\n");
	me->outputEvents.event.EmergencyStopChanged = 1;
	me->EmergencyStop = 0;
	emergencyStopped++;
} else {
	
	if(me->inputEvents.event.DoorReleaseCanister) {
		printf("IO: Door released. Adding canister %i\n", nextBottle);

		me->outputEvents.event.LasersChanged = 1;
		me->DoorSiteLaser = 1;
			
		bottlesActive[nextBottle] = 1;

		nextBottle++;
		nextBottle = nextBottle % NUM_BOTTLES;
		
	}
	if(me->inputEvents.event.InjectorPositionChanged) {
		printf("IO: Injector position changed. Setting move finished.\n");
		me->outputEvents.event.InjectorArmFinishMovement = 1;
	}

	if(me->inputEvents.event.ConveyorChanged) {
		conveyorSpeed = me->ConveyorSpeed;
		printf("IO: Setting conveyor movement to %i\n", conveyorSpeed);
	}

	

	if(me->inputEvents.event.InjectorControlsChanged) {
		printf("IO: Injector controls changed. Now they are Vac: %1i Val: %1i Pmp: %1i\n", me->InjectorVacuumRun, me->InjectorContentsValveOpen, me->InjectorPressurePumpRun);
		if(me->InjectorVacuumRun) {
			printf("IO: Due to vacuum, changing canister pressure to 5.\n");
			me->CanisterPressure = 5;
			me->outputEvents.event.CanisterPressureChanged = 1;
		}
		if(me->InjectorContentsValveOpen) {
			printf("IO: Contents valve now open. Pressure changes slightly, sucking in contents.\n");
			me->CanisterPressure = 20;
			me->outputEvents.event.CanisterPressureChanged = 1;
		}
		if(me->InjectorPressurePumpRun) {
			printf("IO: Due to pressure pump, changing canister pressure to 250.\n");
			me->CanisterPressure = 250;
			me->outputEvents.event.CanisterPressureChanged = 1;
		}
	}
	
	if(me->inputEvents.event.FillContentsChanged) {
		printf("IO: Fill contents changed.\n");
	}

	if(me->inputEvents.event.StartVacuumTimer) {
		printf("IO: Start vacuum timer.\n");//Elapsing timer.\n");
		//me->outputEvents.event.VacuumTimerElapsed = 1;
	}

	

	if(me->inputEvents.event.CanisterCountChanged) {
		printf("IO: Canister count changed. New value: %i\n", me->CanisterCount);
	}


}



}


