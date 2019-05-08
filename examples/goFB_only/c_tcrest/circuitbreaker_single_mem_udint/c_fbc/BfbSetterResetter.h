// This file is generated by FBC.

#ifndef BFBSETTERRESETTER_H_
#define BFBSETTERRESETTER_H_

#include "fbtypes.h"

typedef union {
    UDINT events;
    struct {
        UDINT brk ; // 
        UDINT rst ; // 
        UDINT unsafe ; // 
    } event;
} BfbSetterResetterIEvents;

typedef union {
    UDINT events;
    struct {
        UDINT b_change ; // 
    } event;
} BfbSetterResetterOEvents;

typedef struct {
    UINT _state;
    BOOL _entered;
    BfbSetterResetterIEvents _input;
    BfbSetterResetterOEvents _output;
    INT b; // 
    INT _b;
} BfbSetterResetter;

/* Function block initialization function */
void BfbSetterResetterinit(BfbSetterResetter* me);

/* Function block execution function */
void BfbSetterResetterrun(BfbSetterResetter* me);

/* ECC algorithms */
void BfbSetterResetter_BreakB(BfbSetterResetter* me);
void BfbSetterResetter_CloseB(BfbSetterResetter* me);

#endif // BFBSETTERRESETTER_H_
