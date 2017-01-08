//This is the main file for the iec61499 network with FlexPRET as the top level block

#include "FlexPRET.h"
#include <machine/rtc.h>

//put a copy of the top level block into global memory
struct FlexPRET myFlexPRET;

int main() {
	printf("goFB BottlingPlant starting up\n");
	FlexPRET_init(&myFlexPRET);
	printf("Using %d bytes of memory\n", sizeof(myFlexPRET));
	int x = 0;
	unsigned long long start_time = 0;
	unsigned long long end_time = 0;
	printf("Tick #\t\t# Cycles\n");
	do {
		start_time = get_cpu_cycles();

		FlexPRET_syncEvents(&myFlexPRET);
		FlexPRET_syncData(&myFlexPRET);
		FlexPRET_run(&myFlexPRET);

		end_time = get_cpu_cycles();		
		printf("%4d\t\t%lld\n", x, end_time-start_time-3);
	} while(x++<100);
	printf("Execution halted, we are done here\n");
}

