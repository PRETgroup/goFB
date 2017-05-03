
#include "util.h"

int LREALEqual(LREAL a, LREAL b) {
	if(a > b) {
		return (a-b) <= REL_TOL;
	}
	return (b-a) <= REL_TOL;
}
