# Text-FB specification

Initial thoughts on the textual representation of IEC 61499 Basic and Composite Function Blocks. This document is not exhaustive.

## BFB specification:

```
basicFB (blockname) [compilerheader "(filename"];

interface of (blockname) {
	< 
	in|out event|bool|byte|word|dword|lword|sint|usint|int|uint|dint|udint|lint|ulint|real|lreal|time|any[[size]] [initial "default"] (name[, name]) [with (name[, name]) - non-event types only]; //(comments) 
	>
}

architecture of (blockname) {
	< 
	internal bool|byte|word|dword|lword|sint|usint|int|uint|dint|udint|lint|ulint|real|lreal|time [initial "default"] (name[, name]); //(comments) 
	>

	states {
		< 
		(statename) {
			< 
			run|emit (name[, name])|(in "(languagename)" `(algorithmtext)`); //(comments) 
			>

			< 
			-> (statename) [on (condition text)]; 
			>
		} 
		>
	} 
	
	algorithms {
		<
		(algorithmname) in "(languagename)" `
			(algorithm text)
		`
		>
	} 
}

```
Here's an example:
```
basicFB InjectorMotorController;

interface of InjectorMotorController {
	in event EmergencyStopChanged; //comments are valid
	in event ConveyorStoppedForInject;
	in event PumpFinished;

	in bool EmergencyStop with EmergencyStopChanged;

	out event StartPump, InjectDone, InjectorPositionChanged, InjectRunning; //listing things is permissible in all interface locations

	out byte InjectorPosition with InjectorPositionChanged;
}

architecture of InjectorMotorController {
	internals {
		bool exampleVariableNotActuallyUsed;
		bool otherVariable;
	}

	states {
		MoveArmUp {
			run SetArmUpPosition; 
			emit InjectorPositionChanged;
			emit InjectRunning;

			-> AwaitBottle on InjectorArmFinishedMovement;
		}
		
		AwaitBottle {
			emit InjectDone;

			-> MoveArmDown on ConveyorStoppedForInject;
		}

		MoveArmDown {
			run SetArmDownPosition;
			emit InjectorPositionChanged, InjectRunning;

			-> AwaitPumping on InjectorArmFinishedMovement;
		}

		AwaitPumping {
			emit StartPump;

			-> MoveArmUp on PumpFinished;
		}
	} 
	
	algorithms {
		SetArmDownPosition in 'C' `
			me->InjectorPosition = 255;
			printf("Injector: Set Injector Arm to Down Position");
		`;

		SetArmUpPosition in 'C' `
			me->InjectorPosition = 0;
			printf("Injector: Set Injector Arm to Up Position");
		`;
	} 

}

```

## CFB example

```
CompositeFB InjectorController;

interface of InjectorController {
	in event InjectorArmFinishedMovement;
	in event EmergencyStopChanged;
	in event CanisterPressureChanged;
	in event FillContentsAvailableChanged;
	in event VacuumTimerElapsed;

	in bool EmergencyStop;
	in byte CanisterPressure;
	in byte FillContentsAvailable;

	out event InjectDone;
	out event InjectorPositionChanged;
	out event InjectorControlsChanged;
	out event RejectCanister;
	out event FillContentsChanged;
	out event StartVacuumTimer;
	out event InjectRunning;

	out byte InjectorPosition;
	out bool InjectorContentsValveOpen;
	out bool InjectorVacuumRun;
	out bool InjectorPressurePumpRun;
	out bool FillContents;
}

architecture of InjectorController {
	instance InjectorMotorController Arm;
	instance InjectorPumpsController Pumps;

	events {
		Arm.InjectorArmFinishedMovement <- InjectorArmFinishedMovement;
		Arm.EmergencyStopChanged <- EmergencyStopChanged;
		Arm.ConveyorStoppedForInject <- ConveyorStoppedForInject;
		Arm.PumpFinished <- Pumps.PumpFinished;

		Pumps.EmergencyStopChanged <- EmergencyStopChanged;
		Pumps.CanisterPressureChanged <- CanisterPressureChanged;
		Pumps.FillContentsAvailableChanged <- FillContentsAvailableChanged;
		Pumps.VacuumTimerElapsed <- VacuumTimerElapsed;
		Pumps.StartPump <- Arm.StartPump;

		InjectDone <- Arm.InjectDone;
		InjectorPositionChanged <- Arm.InjectorPositionChanged;
		InjectorControlsChanged <- Pumps.InjectorControlsChanged;
		RejectCanister <- Pumps.RejectCanister;
		FillContentsChanged <- Pumps.FillContentsChanged;
		StartVacuumTimer <- Pumps.StartVacuumTimer;
		InjectRunning <- Arm.InjectRunning;
	}

	data {
		Arm.EmergencyStop <- EmergencyStop;

		Pumps.EmergencyStop <- EmergencyStop;
		Pumps.CanisterPressure <- CanisterPressure;
		Pumps.FillContentsAvailable <- FillContentsAvailable;

		InjectorPosition <- Arm.InjectorPosition;
		InjectorContentsValveOpen <- Pumps.InjectorContentsValveOpen;
		InjectorVacuumRun <- Pumps.InjectorVacuumRun;
		InjectorPressurePumpRun <- Pumps.InjectorPressurePumpRun;
		FillContents <- Pumps.FillContents;

		Example.NotActuallyHere <- `0.075`;
	}
}

```

## EnforcerFB example

```
//P1: AP and VP cannot happen simultaneously.
//P2: VS or VP must be true within AVI after an atrial event AS or AP.
//P3: AS or AP must be true within AEI after a ventricular event VS or VP.
//P4: After a ventricular event, another ventricular event can happen only after URI.
//P5: After a ventricular event, another ventricular event should happen within LRI.
//P6: (fabricated) VP can only be asserted for VPHigh time units.
//P7: (fabricated) AP can only be asserted for APHigh time units.

EnforcerFB XYEnforcer;

interface of XYEnforcer {
	//enforcement parameters
	in int AVI_time initial 1000;
	in int AEI_time initial 1000;
	in int URI_time initial 1000;
	in int LRI_time initial 1000;
	in int VPHigh_time initial 10;
	in int APHigh_time initial 10;

	//enforcement data lines
	enforce bool VP;
	enforce bool AP;
	enforce bool VS;
	enforce bool AS;
}

architecture of XYEnforcer {
	//policies are executed downwards

	//P1: AP and VP cannot happen simultaneously.
	policy {
		observe (AP && VP == 1);
		recovery {
			VP = 0; //set both to be zero
			AP = 0; //could be either really, but as it's a bad order, we'll cancel both, and rely on subsequent enforcers to correct
		}
	}

	//P2: VS or VP must be true within AVI after an atrial event AS or AP.
	policy {
		observe (AS || AP == 1)
		require (VS || VP == 1)
		before (AVI_time);
		recovery {
			VP = 1; //pulse the heart VP
		}
	}

	//P3: AS or AP must be true within AEI after a ventricular event VS or VP.
	policy {
		observe (VS || VP == 1)
		require (AS || AP == 1)
		before (AEI_time);
		recovery {
			AP = 1; //pulse the heart AP
		}
	}

	//P4: After a ventricular event, another ventricular event can happen only after URI.
	policy {
		observe (VS || VP == 1)
		exclude (VS || VP == 1)
		before (URI_time);
		recovery {
			VP = 0; //cancel pulsing the heart VP
		}
	}

	//P5: After a ventricular event, another ventricular event should happen within LRI.
	policy {
		observe (VS || VP == 1)
		require (VS || VP == 1)
		before (LRI_time)
		recovery {
			VP = 1; //pulse the heart VP
		}
	}

	//P6: (fabricated) VP can only be asserted for VPHigh time units.
	policy {
		observe (VP == 1)
		require (VP == 0)
		before (VPHigh_time);
		recovery {
			VP = 0; //deassert the heart VP
		}
	}
	
	//P7: (fabricated) AP can only be asserted for APHigh time units.
	policy {
		observe (AP == 1)
		require (AP == 0)
		before (APHigh_time);
		recovery {
			AP = 0; //deassert the heart AP
		}
	}
}
```