// This file is generated by FBC.

#include "Resource__CBCore0.h"
#include <string.h>
#include <stdio.h>

/* Function block initialization function */
void Resource__CBCore0init(Resource__CBCore0* me)
{
    me->cb3rx.ChanId = 3;
    me->cb2rx.ChanId = 2;
    me->cb1rx.ChanId = 1;
    ArgoRxinit(&me->cb3rx);
    ArgoRxinit(&me->cb2rx);
    ArgoRxinit(&me->cb1rx);
    SifbCBPrintStatusinit(&me->Reference);
}

/* Function block execution function */
void Resource__CBCore0run(Resource__CBCore0* me)
{
    me->Reference._input.event.StatusUpdate = me->cb3rx._output.event.DataPresent || me->cb1rx._output.event.DataPresent || me->cb2rx._output.event.DataPresent;
    me->Reference._St3 = me->cb3rx._Data;
    me->Reference._St2 = me->cb2rx._Data;
    me->Reference._St1 = me->cb1rx._Data;
    
    ArgoRxrun(&me->cb3rx);
    ArgoRxrun(&me->cb2rx);
    ArgoRxrun(&me->cb1rx);
    SifbCBPrintStatusrun(&me->Reference);
    
}