// This file has been automatically generated by goFB
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

//This is the main file for the iec61499 network with top as the top level block

#include "top.h"

//put a copy of the top level block into global memory
top_t mytop;

int main() {
	if(top_preinit(&mytop) != 0) {
		printf("Failed to preinitialize.");
		return 1;
	}
	if(top_init(&mytop) != 0) {
		printf("Failed to initialize.");
		return 1;
	}

	printf("Top: %20s   Size: %lu\n", "top", sizeof(mytop));

	int tickNum = 0;
	do {
		printf("=====TICK %i=====\n",tickNum);
		top_syncOutputEvents(&mytop);
		top_syncInputEvents(&mytop);

		top_syncOutputData(&mytop);
		top_syncInputData(&mytop);
		
		top_run(&mytop);
	} while(tickNum++ < 10);

	return 0;
}

