// This file is generated by FBC.

#ifndef BFBIDMTCURVE_H_
#define BFBIDMTCURVE_H_

#include "fbtypes.h"

typedef union {
    UDINT events;
    struct {
        UDINT tick ; // 
        UDINT i_measured ; // 
        UDINT i_set_change ; // 
    } event;
} BfbIDMTCurveIEvents;

typedef union {
    UDINT events;
    struct {
        UDINT unsafe ; // 
    } event;
} BfbIDMTCurveOEvents;

typedef struct {
    UINT _state;
    BOOL _entered;
    BfbIDMTCurveIEvents _input;
    UDINT i; // 
    UDINT _i;
    UDINT i_set; // 
    UDINT _i_set;
    BfbIDMTCurveOEvents _output;
    ULINT cnt; // 
    UDINT max; // 
    UDINT k; // 
    UDINT b; // 
    UDINT a; // 
} BfbIDMTCurve;

/* Function block initialization function */
void BfbIDMTCurveinit(BfbIDMTCurve* me);

/* Function block execution function */
void BfbIDMTCurverun(BfbIDMTCurve* me);

/* ECC algorithms */
void BfbIDMTCurve_ResetCnt(BfbIDMTCurve* me);
void BfbIDMTCurve_UpdateCnt(BfbIDMTCurve* me);
void BfbIDMTCurve_UpdateMax(BfbIDMTCurve* me);

#endif // BFBIDMTCURVE_H_
