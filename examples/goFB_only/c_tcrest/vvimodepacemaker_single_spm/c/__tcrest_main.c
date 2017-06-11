//This is the main file for the iec61499 network with _TCREST as the top level block

#include "FakePacemakerTop.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory

const int NOC_MASTER = 0;

void t(void* param);

void task(FakePacemakerTop_t _SPM * c0);

void __attribute__ ((noinline)) timed_task(FakePacemakerTop_t _SPM * c) {
	int i;

	for(i=0; i < PROGS_PER_CORE*4; i++) {
		FakePacemakerTop_syncOutputEvents(&c[i]);
		FakePacemakerTop_syncInputEvents(&c[i]);
		FakePacemakerTop_syncOutputData(&c[i]);
		FakePacemakerTop_syncInputData(&c[i]);
		FakePacemakerTop_run(&c[i]);
	}
}


int main() {
	printf("vvi mode pacemaker single spm startup.\n");
	printf("size: %lu", sizeof(FakePacemakerTop_t)*PROGS_PER_CORE*4);

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void task(FakePacemakerTop_t _SPM * c) {
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

	FakePacemakerTop_t _SPM *c;
	c = SPM_BASE;

	int i;
	for(i=0; i<PROGS_PER_CORE*4; i++) {



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

