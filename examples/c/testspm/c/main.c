#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

struct T {
	int x;
};

struct U {
	struct T t;
};

int main() {
	struct U u_place;

	struct U *u;
	u = &u_place;

	struct U _SPM *u_spm;
	u_spm = SPM_BASE;
	
	printf("Setting *u...");
	u->t.x = 1;
	printf("Done\n");

	printf("Setting *u_spm...");
	u_spm->t.x = 1;
	printf("Done\n");

	return 0;
}