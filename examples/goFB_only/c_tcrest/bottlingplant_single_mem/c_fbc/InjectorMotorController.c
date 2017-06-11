// This file is generated by FBC.

#include "InjectorMotorController.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void InjectorMotorControllerinit(InjectorMotorController* me)
{
    me->_state = 0;
    me->_entered = false;
    me->_input.event.InjectorArmFinishedMovement = 0;
    me->_input.event.EmergencyStopChanged = 0;
    me->_input.event.ConveyorStoppedForInject = 0;
    me->_input.event.PumpFinished = 0;

    me->_output.event.StartPump = 0;
    me->_output.event.InjectDone = 0;
    me->_output.event.InjectorPositionChanged = 0;
    me->_output.event.InjectRunning = 0;
}

/* ECC algorithms */
void InjectorMotorController_SetArmDownPosition(InjectorMotorController* me)
{
me->InjectorPosition = 255;
//printf("Injector: Set Injector Arm to Down position\n");
}

void InjectorMotorController_SetArmUpPosition(InjectorMotorController* me)
{
//printf("Injector: Set injector arm to up position\n");
me->InjectorPosition = 0;

}

void InjectorMotorController_Algorithm1(InjectorMotorController* me)
{
//printf("lalalala\n");
}

/* Function block execution function */
void InjectorMotorControllerrun(InjectorMotorController* me)
{
    int _PRET_BOUND_i = 0;
    me->_output.event.StartPump = 0;
    me->_output.event.InjectDone = 0;
    me->_output.event.InjectorPositionChanged = 0;
    me->_output.event.InjectRunning = 0;


        if (me->_input.event.EmergencyStopChanged) {
            me->EmergencyStop = me->_EmergencyStop;
        }
    
	#pragma loopbound min 1 max 2
    for (_PRET_BOUND_i = 0; _PRET_BOUND_i < 2; _PRET_BOUND_i++) {
    asm("#@PRET_Bound 2");
        switch (me->_state) {
        asm("#@PRET_switch_start 1");
            case 0:
                // State: MoveArmUp
                if (!me->_entered) {
                    InjectorMotorController_SetArmUpPosition(me);
                    me->_output.event.InjectorPositionChanged = 1;
                    me->_entered = true;
                    break;
                }
                else {
                    if (me->_input.event.InjectorArmFinishedMovement) {
                        me->_state = 1;
                        me->_entered = false;
                        continue;
                    }
                    break;
                }
                break;
            case 1:
                // State: Await_Bottle
                if (!me->_entered) {
                    me->_output.event.InjectDone = 1;
                    me->_entered = true;
                    break;
                }
                else {
                    if (me->_input.event.ConveyorStoppedForInject) {
                        me->_state = 2;
                        me->_entered = false;
                        continue;
                    }
                    break;
                }
                break;
            case 2:
                // State: MoveArmDown
                if (!me->_entered) {
                    InjectorMotorController_SetArmDownPosition(me);
                    me->_output.event.InjectorPositionChanged = 1;
                    me->_output.event.InjectRunning = 1;
                    me->_entered = true;
                    break;
                }
                else {
                    if (me->_input.event.InjectorArmFinishedMovement) {
                        me->_state = 3;
                        me->_entered = false;
                        continue;
                    }
                    break;
                }
                break;
            case 3:
                // State: Await_Pumping
                if (!me->_entered) {
                    me->_output.event.StartPump = 1;
                    me->_entered = true;
                    break;
                }
                else {
                    if (me->_input.event.PumpFinished) {
                        me->_state = 0;
                        me->_entered = false;
                        continue;
                    }
                    break;
                }
                break;
        }
        break;
    asm("#@PRET_switch_end 1");
    }
    if (me->_output.event.InjectorPositionChanged) {
        me->_InjectorPosition = me->InjectorPosition;
    }

    me->_input.event.InjectorArmFinishedMovement = 0;
    me->_input.event.EmergencyStopChanged = 0;
    me->_input.event.ConveyorStoppedForInject = 0;
    me->_input.event.PumpFinished = 0;
}
