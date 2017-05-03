{{define "utilheader"}}
#include "fbtypes.h"

#define REL_TOL 1e-6
#define ABS_TOL 1e-8

int LREALEqual(LREAL a, LREAL b);
{{end}}

{{define "util"}}
#include "util.h"

int LREALEqual(LREAL a, LREAL b) {
	if(a > b) {
		return (a-b) <= REL_TOL;
	}
	return (b-a) <= REL_TOL;
}
{{end}}