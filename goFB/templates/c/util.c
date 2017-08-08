{{define "utilheader"}}
#ifndef _UTIL_H_ //prevent recursive defines
#define _UTIL_H_
#include "fbtypes.h"

#define REL_TOL 1e-6
#define ABS_TOL 1e-8

//these are useful as terminal colour modifiers
#define T_NRM  "\x1B[0m"
#define T_RED  "\x1B[31m"
#define T_GRN  "\x1B[32m"
#define T_YEL  "\x1B[33m"
#define T_BLU  "\x1B[34m"
#define T_MAG  "\x1B[35m"
#define T_CYN  "\x1B[36m"
#define T_WHT  "\x1B[37m"

//#define PRINT_VALS
#define PRINT_TIME
//#define MAX_TICKS 1000000

int LREAL_EQ(LREAL a, LREAL b);

int LREAL_GTE(LREAL a, LREAL b);

int LREAL_GT(LREAL a, LREAL b);

int LREAL_LTE(LREAL a, LREAL b);

int LREAL_LT(LREAL a, LREAL b);

#endif //_UTIL_H_
{{end}}

{{define "util"}}
#include "util.h"

int LREAL_EQ(LREAL a, LREAL b) {
	if(a > b) {
		return (a-b) <= REL_TOL;
	}
	return (b-a) <= REL_TOL;
}

int LREAL_GTE(LREAL a, LREAL b) {
	return (a + REL_TOL) >= b;
}

int LREAL_GT(LREAL a, LREAL b) {
	return (a + REL_TOL) > b;
}

int LREAL_LTE(LREAL a, LREAL b) {
	return (a - REL_TOL) <= b;
}

int LREAL_LT(LREAL a, LREAL b) {
	return (a - REL_TOL) < b;
}
{{end}}