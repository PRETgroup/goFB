// This file is generated by FBC.

#include "BfbIDMTCurve.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void BfbIDMTCurveinit(BfbIDMTCurve* me)
{
    me->_state = 0;
    me->_entered = false;
    me->_input.events = 0;
    me->_output.events = 0;
    me->cnt = 0;
    me->max = 0.0;
    me->k = 100.0;
    me->b = 0.135;
    me->a = 1.0;
}

/* ECC algorithms */
void BfbIDMTCurve_ResetCnt(BfbIDMTCurve* me)
{
me->cnt = 0;
}

void BfbIDMTCurve_UpdateCnt(BfbIDMTCurve* me)
{
me->cnt = me->cnt + 1;
}

void BfbIDMTCurve_UpdateMax(BfbIDMTCurve* me)
{
me->max = (me->k * me->b) / ((me->i / me->i_set) - 1);
}

/* Function block execution function */
void BfbIDMTCurverun(BfbIDMTCurve* me)
{
    me->_output.events = 0;

    if (me->_input.events) {
        if (me->_input.event.i_measured) {
            me->i = me->_i;
        }
        if (me->_input.event.i_set_change) {
            me->i_set = me->_i_set;
        }
    }
    for (;;) {
        switch (me->_state) {
            case 0:
                // State: INIT
                if (!me->_entered) {
                    me->_entered = true;
                }
                else {
                    me->_state = 1;
                    me->_entered = false;
                    continue;
                }
                break;
            case 1:
                // State: SAFE
                if (!me->_entered) {
                    BfbIDMTCurve_ResetCnt(me);
                    me->_entered = true;
                }
                else {
                    if (me->_input.event.tick && (me->i > me->i_set)) {
                        me->_state = 2;
                        me->_entered = false;
                        continue;
                    }
                }
                break;
            case 2:
                // State: COUNT
                if (!me->_entered) {
                    BfbIDMTCurve_UpdateCnt(me);
                    BfbIDMTCurve_UpdateMax(me);
                    me->_entered = true;
                }
                else {
                    if (me->_input.event.tick && (me->cnt >= me->max)) {
                        me->_state = 3;
                        me->_entered = false;
                        continue;
                    }
                    else if (me->_input.event.tick && (me->i < me->i_set)) {
                        me->_state = 1;
                        me->_entered = false;
                        continue;
                    }
                }
                break;
            case 3:
                // State: UNSAFE
                if (!me->_entered) {
                    me->_output.event.unsafe = 1;
                    me->_entered = true;
                }
                else {
                    if (me->_input.event.tick && (me->i < me->i_set)) {
                        me->_state = 1;
                        me->_entered = false;
                        continue;
                    }
                }
                break;
        }
        break;
    }

    me->_input.events = 0;
}
