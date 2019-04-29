// This file is generated by FBC.

#ifndef _CBCORESINGLE_H_
#define _CBCORESINGLE_H_

#include "fbtypes.h"
#include "SifbAmmeter.h"
#include "SifbTimer.h"
#include "CfbBreakerController.h"
#include "SifbManagementControls.h"
#include "SifbIntLed.h"
#include "SifbCBPrintStatus.h"

typedef struct {
    SifbAmmeter amm1;
    SifbTimer timer1;
    CfbBreakerController cb1;
    SifbManagementControls hm1;
    SifbIntLed led1;
    SifbManagementControls hm3;
    SifbIntLed led3;
    SifbAmmeter amm3;
    CfbBreakerController cb3;
    SifbTimer timer3;
    SifbManagementControls hm2;
    SifbIntLed led2;
    SifbAmmeter amm2;
    CfbBreakerController cb2;
    SifbTimer timer2;
    SifbCBPrintStatus print;
} _CBCoreSingle;

/* Function block initialization function */
void _CBCoreSingleinit(_CBCoreSingle* me);

/* Function block execution function */
void _CBCoreSinglerun(_CBCoreSingle* me);

#endif // _CBCORESINGLE_H_
