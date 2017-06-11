//This is the main file for the iec61499 network with _TCREST as the top level block

#include "FakePacemakerTop.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory

const int NOC_MASTER = 0;

void t(void* param);

void task(FakePacemakerTop_t * c0);


int main() {
	printf("vvi mode pacemaker single-mem startup.\n");
	printf("size: %lu", sizeof(FakePacemakerTop_t)*4);

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(FakePacemakerTop_t * c) {
	int i;

	for(i=0; i < PROGS_PER_CORE; i++) {
		FakePacemakerTop_syncOutputEvents(&c[i]);
		FakePacemakerTop_syncInputEvents(&c[i]);
		FakePacemakerTop_syncOutputData(&c[i]);
		FakePacemakerTop_syncInputData(&c[i]);
		FakePacemakerTop_run(&c[i]);
	}
}

void task(FakePacemakerTop_t * c) {
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

	FakePacemakerTop_t c[PROGS_PER_CORE];

	int i;
	for(i=0; i<PROGS_PER_CORE; i++) {



		if(FakePacemakerTop_preinit(&c[i]) != 0 || FakePacemakerTop_init(&c[i]) != 0) {
			HEX = 15;
			return;
		}

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

