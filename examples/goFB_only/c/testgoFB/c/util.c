
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


int REAL_EQ(REAL a, REAL b) {
	if(a > b) {
		return (a-b) <= REL_TOL;
	}
	return (b-a) <= REL_TOL;
}

int REAL_GTE(REAL a, REAL b) {
	return (a + REL_TOL) >= b;
}

int REAL_GT(REAL a, REAL b) {
	return (a + REL_TOL) > b;
}

int REAL_LTE(REAL a, REAL b) {
	return (a - REL_TOL) <= b;
}

int REAL_LT(REAL a, REAL b) {
	return (a - REL_TOL) < b;
}
