# Text-FB specification

## BFB example

```
BasicFB myExampleController;

interface of myExampleBlock {
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

architecture of myExampleBlock {
	//internal type Name

	ecc {
		state MoveArmUp {
			run SetArmUpPosition; 
			emit InjectorPositionChanged;
			emit InjectRunning;
		}

		state AwaitBottle {
			emit InjectDone;
		}

		state MoveArmDown {
			run SetArmDown;
			emit InjectorPositionChanged;
			emit InjectRunning;
		}

		state AwaitPumping {
			emit StartPump;
		}

		transition from AwaitBottle to MoveArmDown needs {
			ConveyorStoppedForInject;
		}

		transition from MoveArmDown to AwaitPumping needs {
			InjectorArmFinishedMovement;
		}

		transition from AwaitPumping to MoveArmUp needs {
			PumpFinished;
		}

		transition from MoveArmUp to AwaitBottle needs {
			InjectorArmFinishedMovement;
		}
	}

	algorithm SetArmDownPosition in 'C' {
		me->InjectorPosition = 255;
		printf("Injector: Set Injector Arm to Down Position");
	}

	algorithm SetArmUpPosition in 'C' {
		me->InjectorPosition = 0;
		printf("Injector: Set Injector Arm to Up Position");
	}

}

```

## CFB example