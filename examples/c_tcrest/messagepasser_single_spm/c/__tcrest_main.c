//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_Core022.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t(void* param);

void task(_Core022_t _SPM * c);

int main() {
	printf("messagepasser_mem tcrest4 startup.\n");
	printf("size: %lu", sizeof(_Core022_t)*PROGS_PER_CORE*4);

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core022_t _SPM * c) {
	for(int i=0; i<PROGS_PER_CORE*2;i++) {
		_Core022_syncOutputEvents(&c[i]);
		_Core022_syncInputEvents(&c[i]);
		_Core022_syncOutputData(&c[i]);
		_Core022_syncInputData(&c[i]);
		_Core022_run(&c[i]);
	}
}

void task(_Core022_t _SPM * c) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task(c);

		end_time = get_cpu_cycles();
		printf("%4d\t\t%lld\n", tickCount, (end_time-start_time-3)*2); //some cheese is done with this analysis

		tickCount++;
	} while(1);
}

void t(void* param) {
	HEX = 7;
	_Core022_t _SPM *c;
	c = SPM_BASE;
	
	//c0 = &c; //SPM_BASE;
	for(int i=0; i < PROGS_PER_CORE*2; i++) {
		if(_Core022_preinit(&c[i]) != 0 || _Core022_init(&c[i]) != 0) {
			HEX = 15;
			return;
		}
	}

	HEX = 10;
	
	task(c);
}
