// This file is generated by FBC.

#ifndef PASS_H_
#define PASS_H_

#include "fbtypes.h"

typedef union {
    struct {
        UDINT CountChanged ; // 
    } event;
} PassIEvents;

typedef union {
    struct {
        UDINT OutCountChanged ; // 
    } event;
} PassOEvents;

typedef struct {
    UINT _state;
    BOOL _entered;
    PassIEvents _input;
    LINT Count; // 
    LINT _Count;
    PassOEvents _output;
    LINT OutCount; // 
    LINT _OutCount;
} Pass;

/* Function block initialization function */
void Passinit(Pass* me);

/* Function block execution function */
void Passrun(Pass* me);

/* ECC algorithms */
void Pass_UpdateCountOut(Pass* me);

#endif // PASS_H_