#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

typedef struct {
	int x;
} T_t;

typedef struct {
	T_t t;
} U_t;

int main() {
	U_t _SPM *u_spm;
	u_spm = SPM_BASE;
	
	// printf("Setting *u...");
	// u->t.x = 1;
	// printf("Done\n");

	printf("Setting *u_spm...");
	u_spm->t.x = 1;
	printf("Done\n");

	printf("u_spm->t.x = %i\n", u_spm->t.x);

	return 0;
}