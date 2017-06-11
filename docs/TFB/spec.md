# Text-FB specification

Initial thoughts on the textual representation of IEC 61499 Basic and Composite Function Blocks

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