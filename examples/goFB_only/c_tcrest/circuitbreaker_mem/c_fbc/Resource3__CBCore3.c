// This file is generated by FBC.

#include "Resource3__CBCore3.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void Resource3__CBCore3init(Resource3__CBCore3* me)
{
    me->tx.ChanId = 3;
    SifbTimerinit(&me->timer);
    ArgoTxinit(&me->tx);
    SifbManagementControlsinit(&me->hmi);
    SifbIntLedinit(&me->led);
    CfbBreakerControllerinit(&me->cb);
    SifbAmmeterinit(&me->amm);
    SawmillMessageHandlerinit(&me->msgh);
}

/* Function block execution function */
void Resource3__CBCore3run(Resource3__CBCore3* me)
{
    me->tx._input.event.DataPresent = me->msgh._output.event.TxDataPresent;
    me->tx._Data = me->msgh._TxData;
    me->led._i = me->cb._b;
    me->led._input.event.i_change = me->cb._output.event.b_change;
    me->cb._input.event.brk = me->hmi._output.event.brk;
    me->cb._input.event.i_measured = me->amm._output.event.i_measured;
    me->cb._input.event.rst = me->hmi._output.event.rst;
    me->cb._i_set = me->hmi._i_set;
    me->cb._i = me->amm._i;
    me->cb._input.event.i_set_change = me->hmi._output.event.i_set_change;
    me->cb._input.event.tick = me->timer._output.event.Tick;
    me->msgh._input.event.TxSuccessChanged = me->tx._output.event.SuccessChanged;
    me->msgh._input.event.MessageChanged = me->cb._output.event.b_change;
    me->msgh._Message = me->cb._b;
    me->msgh._TxSuccess = me->tx._Success;
    
    SifbTimerrun(&me->timer);
    ArgoTxrun(&me->tx);
    SifbManagementControlsrun(&me->hmi);
    SifbIntLedrun(&me->led);
    CfbBreakerControllerrun(&me->cb);
    SifbAmmeterrun(&me->amm);
    SawmillMessageHandlerrun(&me->msgh);
    
}
