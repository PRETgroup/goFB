// This file is generated by FBC.

#ifndef RESOURCE1__CBCORE1_H_
#define RESOURCE1__CBCORE1_H_

#include "fbtypes.h"
#include "CfbBreakerController.h"
#include "SawmillMessageHandler.h"
#include "ArgoTx.h"
#include "SifbIntLed.h"
#include "SifbTimer.h"
#include "SifbManagementControls.h"
#include "SifbAmmeter.h"

typedef struct {
    CfbBreakerController cb;
    SawmillMessageHandler msgh;
    ArgoTx tx;
    SifbIntLed led;
    SifbTimer timer;
    SifbManagementControls hmi;
    SifbAmmeter amm;
} Resource1__CBCore1;

/* Function block initialization function */
void Resource1__CBCore1init(Resource1__CBCore1* me);

/* Function block execution function */
void Resource1__CBCore1run(Resource1__CBCore1* me);

#endif // RESOURCE1__CBCORE1_H_
