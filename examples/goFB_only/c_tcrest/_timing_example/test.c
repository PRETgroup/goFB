#include <stdio.h>

#include <machine/spm.h>
#include <machine/patmos.h>
#include "libcorethread/corethread.h"
#include "libmp/mp.h"

const int NOC_MASTER = 0;

__attribute__ ((noinline)) void t(void* param);

int main() {
	printf("Hallo\n");

	t(NULL);

	return 0;
}

int x;

__attribute__ ((noinline)) void t(void* param) {
    int i = 1;
    i++;
    if(i > 2) {
        i--;
    }
}
