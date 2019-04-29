// This file is generated by FBC.

#ifndef SIFBTIMER_H_
#define SIFBTIMER_H_

#include "fbtypes.h"

typedef union {
    UDINT events;
    struct {
        UDINT Tick : 1; // 
    } event;
} SifbTimerOEvents;

typedef struct {
    BOOL _entered;
    SifbTimerOEvents _output;
} SifbTimer;

/* Function block initialization function */
void SifbTimerinit(SifbTimer* me);

/* Function block execution function */
void SifbTimerrun(SifbTimer* me);

#endif // SIFBTIMER_H_
