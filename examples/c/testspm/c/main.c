#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

int main() {

	unsigned int _SPM * x;
	x = SPM_BASE;
	*x = 10;

	printf("x = %d\n", *x);
	printf("x = %d\n", *x);
	printf("x = %d\n", *x);

}