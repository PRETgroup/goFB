//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t(void* param);

void task(_Core_t * c0);


int main() {
	printf("pid_single_mem patmos startup.\n");
	printf("sizes: %lu\n", sizeof(_Core_t)*4);
	printf("starting\n");

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core_t * c) {
	int i;
	for(i=0; i < PROGS_PER_CORE*4; i++) {
		_Core_syncOutputEvents(&c[i]);
		_Core_syncInputEvents(&c[i]);
		_Core_syncOutputData(&c[i]);
		_Core_syncInputData(&c[i]);
		_Core_run(&c[i]);
	}
}

void task(_Core_t * c) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task(c);

		end_time = get_cpu_cycles();
		// printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t(void* param) {
	HEX = 7;

	_Core_t c[PROGS_PER_CORE*4];

	int i;
	for(i=0; i<PROGS_PER_CORE*4; i++) {
		
		if(_Core_preinit(&c[i]) != 0 || _Core_init(&c[i]) != 0) {
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

