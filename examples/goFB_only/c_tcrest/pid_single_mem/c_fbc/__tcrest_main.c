//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_Core.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t(void* param);

void task(_Core * c0);


int main() {
	printf("fbc pid_mem patmos startup.\n");
	printf("sizes: %lu", sizeof(_Core)*PROGS_PER_CORE*4);


	printf("Started\n");

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core * c) {
	int i;

	for(i=0; i < PROGS_PER_CORE*4; i++) {

		_Corerun(&c[i]);
	}
}

void task(_Core * c) {
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

	_Core c[PROGS_PER_CORE*4];

	int i;
	for(i=0; i<PROGS_PER_CORE*4; i++) {

		_Coreinit(&c[i]);
	}
	

	HEX = 8;

	// if(mp_init_ports() == 0) {
	// 	HEX = 16;
	// 	return;
	// }
	HEX = 9;

	// c0init = 1;
	// while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task(c);
}

