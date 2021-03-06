// This file is generated by FBC.

#include "BfbSetterResetter.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void BfbSetterResetterinit(BfbSetterResetter* me)
{
    me->_state = 0;
    me->_entered = false;
    me->_input.event.brk = 0;
    me->_input.event.rst = 0;
    me->_input.event.unsafe = 0;
    me->_output.event.b_change = 0;
}

/* ECC algorithms */
void BfbSetterResetter_BreakB(BfbSetterResetter* me)
{
me->b = 1;
}

void BfbSetterResetter_CloseB(BfbSetterResetter* me)
{
me->b = 0;
}

/* Function block execution function */
void BfbSetterResetterrun(BfbSetterResetter* me)
{
    me->_output.event.b_change = 0;

    #pragma loopbound min 1 max 2
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
                // State: OPEN_CONTACTS
                if (!me->_entered) {
                    BfbSetterResetter_BreakB(me);
                    me->_output.event.b_change = 1;
                    me->_entered = true;
                }
                else {
                    if (me->_input.event.rst) {
                        me->_state = 2;
                        me->_entered = false;
                        continue;
                    }
                }
                break;
            case 2:
                // State: CLOSE_CONTACTS
                if (!me->_entered) {
                    BfbSetterResetter_CloseB(me);
                    me->_output.event.b_change = 1;
                    me->_entered = true;
                }
                else {
                    if (me->_input.event.brk || me->_input.event.unsafe) {
                        me->_state = 1;
                        me->_entered = false;
                        continue;
                    }
                }
                break;
        }
        break;
    }
    if (me->_output.event.b_change) {
        me->_b = me->b;
    }

    me->_input.event.brk = 0;
    me->_input.event.rst = 0;
    me->_input.event.unsafe = 0;
}

