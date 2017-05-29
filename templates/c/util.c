{{define "utilheader"}}
#include "fbtypes.h"

#define REL_TOL 1e-6
#define ABS_TOL 1e-8

//#define PRINT_VALS
#define PRINT_TIME
//#define MAX_TICKS 1000000

int LREAL_EQ(LREAL a, LREAL b);

int LREAL_GTE(LREAL a, LREAL b);

int LREAL_GT(LREAL a, LREAL b);

int LREAL_LTE(LREAL a, LREAL b);

int LREAL_LT(LREAL a, LREAL b);

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