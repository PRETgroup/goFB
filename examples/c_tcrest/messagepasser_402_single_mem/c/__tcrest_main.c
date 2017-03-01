//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_Core402.h"

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t0(void* param);

void task0(_Core402_t * c0, _Core402_t * c1, _Core402_t * c2, _Core402_t * c3);


int main() {
	printf("messagepasser_402_single_mem patmos startup.\n");
	printf("size: %lu\n", sizeof(_Core402_t)*4);

	t0(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task0(_Core402_t * c0, _Core402_t * c1, _Core402_t * c2, _Core402_t * c3) {
	_Core402_syncOutputEvents(c0);
	_Core402_syncInputEvents(c0);
	_Core402_syncOutputData(c0);
	_Core402_syncInputData(c0);
	_Core402_run(c0);

	_Core402_syncOutputEvents(c1);
	_Core402_syncInputEvents(c1);
	_Core402_syncOutputData(c1);
	_Core402_syncInputData(c1);
	_Core402_run(c1);

	_Core402_syncOutputEvents(c2);
	_Core402_syncInputEvents(c2);
	_Core402_syncOutputData(c2);
	_Core402_syncInputData(c2);
	_Core402_run(c2);

	_Core402_syncOutputEvents(c3);
	_Core402_syncInputEvents(c3);
	_Core402_syncOutputData(c3);
	_Core402_syncInputData(c3);
	_Core402_run(c3);
}

void task0(_Core402_t * c0, _Core402_t * c1, _Core402_t * c2, _Core402_t * c3) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task0(c0, c1, c2, c3);

		end_time = get_cpu_cycles();
		printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t0(void* param) {
	HEX = 7;
	_Core402_t c0;
	_Core402_t c1;
	_Core402_t c2;
	_Core402_t c3;
	
	_Core402_preinit(&c0);
	_Core402_init(&c0);

	_Core402_preinit(&c1);
	_Core402_init(&c1);
	
	_Core402_preinit(&c2);
	_Core402_init(&c2);
	
	_Core402_preinit(&c3);
	_Core402_init(&c3);

	HEX = 10;
	
	task0(&c0, &c1, &c2, &c3);
}


