// This file is generated by FBC.

#ifndef PLANTPROCESS_H_
#define PLANTPROCESS_H_

#include "fbtypes.h"

typedef union {
    struct {
        UDINT Zero; // 
        UDINT ControlChange ; // 
        UDINT RandomChange ; // 
    } event;
} PlantProcessIEvents;

typedef union {
    struct {
        UDINT ValueChange ; // 
    } event;
} PlantProcessOEvents;

typedef struct {
    UINT _state;
    BOOL _entered;
    PlantProcessIEvents _input;
    REAL Control; // 
    REAL _Control;
    REAL Random; // 
    REAL _Random;
    PlantProcessOEvents _output;
    REAL Value; // 
    REAL _Value;
} PlantProcess;

/* Function block initialization function */
void PlantProcessinit(PlantProcess* me);

/* Function block execution function */
void PlantProcessrun(PlantProcess* me);

/* ECC algorithms */
void PlantProcess_PlantZero(PlantProcess* me);
void PlantProcess_PlantTick(PlantProcess* me);

#endif // PLANTPROCESS_H_