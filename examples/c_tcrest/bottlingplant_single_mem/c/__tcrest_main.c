//This is the main file for the iec61499 network with _TCREST as the top level block

#include "FlexPRET.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t(void* param);

void task(FlexPRET_t * c);

int main() {
	printf("messagepasser_mem tcrest4 startup.\n");
	printf("sizes: %lu", sizeof(FlexPRET_t)*PROGS_PER_CORE*4);

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(FlexPRET_t * c) {
	for(int i=0; i<PROGS_PER_CORE*4;i++) {
		FlexPRET_syncOutputEvents(&c[i]);
		FlexPRET_syncInputEvents(&c[i]);
		FlexPRET_syncOutputData(&c[i]);
		FlexPRET_syncInputData(&c[i]);
		FlexPRET_run(&c[i]);
	}
}

void task(FlexPRET_t * c) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task(c);

		end_time = get_cpu_cycles();
		printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t(void* param) {
	HEX = 7;
	FlexPRET_t c[PROGS_PER_CORE*4];
	
	//c0 = &c; //SPM_BASE;
	for(int i=0; i < PROGS_PER_CORE*4; i++) {
		if(FlexPRET_preinit(&c[i]) != 0 || FlexPRET_init(&c[i]) != 0) {
			HEX = 15;
			return;
		}
	}

	HEX = 10;
	
	task(c);
}
