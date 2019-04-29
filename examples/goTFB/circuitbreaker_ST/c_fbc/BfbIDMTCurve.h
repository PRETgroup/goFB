// This file is generated by FBC.

#ifndef BFBIDMTCURVE_H_
#define BFBIDMTCURVE_H_

#include "fbtypes.h"
#include "STFunctions.h"

typedef union {
    UDINT events;
    struct {
        UDINT tick : 1; // 
        UDINT i_measured : 1; // 
        UDINT iSet_change : 1; // 
    } event;
} BfbIDMTCurveIEvents;

typedef union {
    UDINT events;
    struct {
        UDINT unsafe : 1; // 
    } event;
} BfbIDMTCurveOEvents;

typedef struct {
    UINT _state;
    BOOL _entered;
    BfbIDMTCurveIEvents _input;
    USINT i; // 
    USINT _i;
    USINT iSet; // 
    USINT _iSet;
    BfbIDMTCurveOEvents _output;
    ULINT cnt; // 
    LREAL max; // 
    REAL k; // 
    REAL b; // 
    REAL a; // 
} BfbIDMTCurve;

/* Function block initialization function */
void BfbIDMTCurveinit(BfbIDMTCurve* me);

/* Function block execution function */
void BfbIDMTCurverun(BfbIDMTCurve* me);

/* ECC algorithms */
void BfbIDMTCurve_s_safe_alg0(BfbIDMTCurve* me);
void BfbIDMTCurve_s_count_alg0(BfbIDMTCurve* me);
void BfbIDMTCurve_updateMax(BfbIDMTCurve* me);

#endif // BFBIDMTCURVE_H_