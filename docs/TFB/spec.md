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

## EnforcerFB semantics
Enforcers are designed to work with GALS semantics with other FBs (at least at the moment). They don't need to use events on their I/O,
and instead use valued signals. 

EnforcerFBs are designed to execute a number of policies on valued signals. 

A ```require (condition) [before (duration)] recovery `code` ``` command is used to specify requirements as well as the recovery plans.

For instance, ```require(A && B == 1) recovery `A = 0;` ``` says that _If 'A' and 'B' are equal to '1', set 'A' to be '0'_. 

`observe(condition) [before (duration)]` are used for sequences, and pre-conditions.
When a given condition in an `observe` block becomes `true`, the contents of the observe block come into effect.
`observe` blocks also set the _start time_ for timers internal to their block. 

For instance,
```
observe(A) {
	require (B) before (30ms) recovery `
		B = 1;
	`
}
```
This says that _If 'A' is true, 'B' must be true before 30ms, else set 'B' to be true._

When considering `[before (duration)]` for `observe` and `require`, it is important to note that the default _start time_ is the currently active `observe`. That is, top-level `observe` and `require` cannot have a `before` clause.

However, in deep-level nesting we can also label our `observe` blocks, and use the `from` keyword in the `before` clause to specify which `observe` we are measuring from. This is present in Policy P1 of AlphabetEnforcer in the examples below.


## EnforcerFB examples

```
//P1: AP and VP cannot happen simultaneously.
//P2: VS or VP must be true within AVI after an atrial event AS or AP.
//P3: AS or AP must be true within AEI after a ventricular event VS or VP.
//P4: After a ventricular event, another ventricular event can happen only after URI.
//P5: After a ventricular event, another ventricular event should happen within LRI.
//P6: (fabricated) VP can only be asserted for VPHigh time units.
//P7: (fabricated) AP can only be asserted for APHigh time units.

EnforcerFB PaceEnforcer;

interface of PaceEnforcer {
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

architecture of PaceEnforcer {
	//policies are executed downwards

	//P1: AP and VP cannot happen simultaneously.
	policy {
		require (AP && VP != 1)
		recover `
			VP = 0; //set both to be zero
			AP = 0; //could be either really, but as it's a bad order, we'll cancel both, and rely on subsequent enforcers to correct
		`;
	}

	//P2: VS or VP must be true within AVI after an atrial event AS or AP.
	policy {
		observe (AS || AP == 1) {
			require (VS || VP == 1)
			before (AVI_time)
			recover `
				VP = 1; //pulse the heart VP
			`;
		}
	}

	//P3: AS or AP must be true within AEI after a ventricular event VS or VP.
	policy {
		observe (VS || VP == 1) {
			require (AS || AP == 1)
			before (AEI_time)
			recover `
				AP = 1; //pulse the heart AP
			`;
		}
	}

	//P4: After a ventricular event, another ventricular event can happen only after URI.
	policy {
		observe (VS || VP == 1) {
			require (VS || VP != 1)
			before (URI_time)
			recover `
				VP = 0; //cancel pulsing the heart VP
			`;
		}
	}

	//P5: After a ventricular event, another ventricular event should happen within LRI.
	policy {
		observe (VS || VP == 1) {
			require (VS || VP == 1)
			before (LRI_time)
			recover `
				VP = 1; //pulse the heart VP
			`;
		}
	}

	//P6: (fabricated) VP can only be asserted for VPHigh time units.
	policy {
		observe (VP == 1) {
			require (VP == 0)
			before (VPHigh_time)
			recover `
				VP = 0; //deassert the heart VP
			`;
		}
	}
	
	//P7: (fabricated) AP can only be asserted for APHigh time units.
	policy {
		observe (AP == 1) {
			require (AP == 0)
			before (APHigh_time)
			recover `
				AP = 0; //deassert the heart AP
			`;
		}
	}
}
```

```
//P1: If B happens within 30ms of A, C must happen within 60ms of A, and D must happen within 10ms of B.
//P2: Flashing light L. L needs to be on for between 500ms and 600ms, and then off for between 500ms and 600ms, and then on again etc.
EnforcerFB AlphabetEnforcer;

interface of AlphabetEnforcer {
	//enforcement data lines
	enforce bool A;
	enforce bool B;
	enforce bool C;
	enforce bool D;
	enforce bool L;
}

architecture of AlphabetEnforcer {
	//policies are executed downwards

	//P1: If B happens within 30ms of A, C must happen within 60ms of A, and D must happen within 10ms of B.
	policy {
		AStart: observe(A) {
			BStart: observe(B) before (30ms from AStart) {
				require (C) before (60ms from AStart) recover `
					C = 1;
				`;
				require (D) before (10ms from BStart) recover `
					D = 1;
				`;
			}
		}
	}

	//equivalent P1: If B happens within 30ms of A, C must happen within 60ms of A, and D must happen within 10ms of B.
	policy {
		AStart: observe(A) {
			observe(B) before (30ms) {
				require (C) before (60ms from AStart) recover `
					C = 1;
				`;
				require (D) before (10ms) recover `
					D = 1;
				`;
			}
		}
	}

	//P2: Flashing light L. L needs to be on for between 500ms and 600ms, and then off for between 500ms and 600ms, and then on again etc.
	policy {
		observe(L==1) {
			require (L==1) before (500ms) recover `
				L = 1;
			`;
			require (L==0) before (600ms) recover `
				L = 0;
			`;
		}
		observe(L==0) {
			require (L==0) before (500ms) recover `
				L = 0;
			`;
			require (L==1) before (600ms) recover `
				L = 1;
			`;
		}
	}
}
```

## EFB v2 example

Here's an example:
```
enforcerFB AEIEnforcer;

interface of AEIEnforcer {
	enforce in event AS, VS; //in here means that they're going from PLANT to CONTROLLER
	enforce out event AP, VP;//out here means that they're going from CONTROLLER to PLANT

	in ulint AEI_ns default 900000000;

	in event current_ns_change;
	in ulint current_ns with current_ns_change;
}

architecture of AEIEnforcer {
	internals {
		ulint tAEI;
	}

	//P3: AS or AP must be true within AEI after a ventricular event VS or VP.

	states {
		s0 {
			//-> <destination> [on guard] [: output expression][, output expression...] ;
			-> s1 on (VS or VP): tAEI := current_ns;
		}

		s1 {
			-> s1 on (AS or AP);
			-> violation on (tAEI + AEI_ns > current_ns);
		}
	} 
}

```
