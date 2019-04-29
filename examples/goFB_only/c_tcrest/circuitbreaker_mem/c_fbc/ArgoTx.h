// This file is generated by FBC.

#ifndef ARGOTX_H_
#define ARGOTX_H_

#include "fbtypes.h"

typedef union {
    UDINT events;
    struct {
        UDINT DataPresent : 1; // 
    } event;
} ArgoTxIEvents;

typedef union {
    UDINT events;
    struct {
        UDINT SuccessChanged : 1; // 
    } event;
} ArgoTxOEvents;

/* goFB_ignore */
typedef struct {
    BOOL _entered;
    ArgoTxIEvents _input;
    INT Data; // 
    INT _Data;
    UDINT ChanId; // 
    UDINT _ChanId;
    ArgoTxOEvents _output;
    BOOL Success; // 
    BOOL _Success;
} ArgoTx;

/* Function block initialization function */
void ArgoTxinit(ArgoTx* me);

/* Function block execution function */
void ArgoTxrun(ArgoTx* me);

#endif // ARGOTX_H_
