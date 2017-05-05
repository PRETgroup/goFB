# Text-FB specification

## BFB example

```
BasicFB InjectorMotorController;

interface of InjectorMotorController {
	in event EmergencyStopChanged;
	in event ConveyorStoppedForInject;
	in event PumpFinished;

	in bool EmergencyStop on EmergencyStopChanged;

	out event StartPump;
	out event InjectDone;
	out event InjectorPositionChanged;
	out event InjectRunning;

	out byte InjectorPosition on InjectorPositionChanged;
}

architecture of InjectorMotorController {
	internal bool exampleVariableNotActuallyUsed;

	states {
		MoveArmUp {
			run SetArmUpPosition; 
			emit InjectorPositionChanged;
			emit InjectRunning;
		}
		
		AwaitBottle {
			emit InjectDone;
		}

		MoveArmDown {
			run SetArmDown;
			emit InjectorPositionChanged;
			emit InjectRunning;
		}

		AwaitPumping {
			emit StartPump;
		}
	} 

	transitions {
		AwaitBottle -> MoveArmDown on ConveyorStoppedForInject;
		MoveArmDown -> AwaitPumping on InjectorArmFinishedMovement;
		AwaitPumping -> MoveArmUp on PumpFinished;
		MoveArmUp -> AwaitBottle on InjectorArmFinishedMovement;
	}
	
	algorithms {
		SetArmDownPosition in 'C' {
			me->InjectorPosition = 255;
			printf("Injector: Set Injector Arm to Down Position");
		}

		SetArmUpPosition in 'C' {
			me->InjectorPosition = 0;
			printf("Injector: Set Injector Arm to Up Position");
		}
	} 

}

```

## CFB example

```
CompositeFB InjectorController;

interface of myExampleBlock {
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
	internal InjectorMotorController Arm;
	internal InjectorPumpsController Pumps;

	events {
		InjectorArmFinishedMovement -> Arm.InjectorArmFinishedMovement;
		EmergencyStopChanged -> Arm.EmergencyStopChanged;
		EmergencyStopChanged -> Pumps.EmergencyStopChanged;
		CanisterPressureChanged -> Pumps.CanisterPressureChanged;
		FillContentsAvailableChanged -> Pumps.FillContentsAvailableChanged;
		ConveyorStoppedForInject -> Arm.ConveyorStoppedForInject;
		VacuumTimerElapsed -> Pumps.VacuumTimerElapsed;
		Arm.InjectDone -> InjectDone;
		Arm.InjectorPositionChanged -> InjectorPositionChanged;
		Pumps.InjectorControlsChanged -> InjectorControlsChanged;
		Pumps.RejectCanister -> RejectCanister;
		Pumps.FillContentsChanged -> FillContentsChanged;
		Pumps.StartVacuumTimer -> StartVacuumTimer;
		Arm.InjectRunning -> InjectRunning;
		Arm.StartPump -> Pumps.StartPump;
		Pumps.PumpFinished -> Arm.PumpFinished;
	}

	data {
		EmergencyStop -> Arm.EmergencyStop;
		EmergencyStop -> Pumps.EmergencyStop;
		CanisterPressure -> Pumps.CanisterPressure;
		FillContentsAvailable -> Pumps.FillContentsAvailable;
		Arm.InjectorPosition -> InjectorPosition;
		Pumps.InjectorContentsValveOpen -> InjectorContentsValveOpen;
		Pumps.InjectorVacuumRun -> InjectorVacuumRun;
		Pumps.InjectorPressurePumpRun -> InjectorPressurePumpRun;
		Pumps.FillContents -> FillContents;
	}
}

```