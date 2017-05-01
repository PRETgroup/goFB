
#include "util.h"

int LREALEqual(LREAL a, LREAL b) {
	if(a > b) {
		return (a-b) < ABS_TOL;
	}
	return (b-a) < ABS_TOL;
}
