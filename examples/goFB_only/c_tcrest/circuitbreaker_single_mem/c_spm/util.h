
#ifndef _UTIL_H_ //prevent recursive defines
#define _UTIL_H_
#include "fbtypes.h"

#define REL_TOL 1e-6
#define ABS_TOL 1e-8

//these are useful as terminal colour modifiers
#define T_NRM  "\x1B[0m"

//TF = text foreground
#define TF_RED  "\x1B[31m"
#define TF_GRN  "\x1B[32m"
#define TF_YEL  "\x1B[33m"
#define TF_BLU  "\x1B[34m"
#define TF_MAG  "\x1B[35m"
#define TF_CYN  "\x1B[36m"
#define TF_WHT  "\x1B[37m"

//TB = text background
#define TB_BLK "\x1B[40m"
#define TB_RED "\x1B[41m"
#define TB_GRN "\x1B[42m"
#define TB_YEL "\x1B[43m"
#define TB_BLU "\x1B[44m"
#define TB_MAG "\x1B[45m"
#define TB_CYN "\x1B[46m"
#define TB_WHT "\x1B[47m"

//#define PRINT_VALS
#define PRINT_TIME
//#define MAX_TICKS 1000000

int LREAL_EQ(LREAL a, LREAL b);

int LREAL_GTE(LREAL a, LREAL b);

int LREAL_GT(LREAL a, LREAL b);

int LREAL_LTE(LREAL a, LREAL b);

int LREAL_LT(LREAL a, LREAL b);

#endif //_UTIL_H_
