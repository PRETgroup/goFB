//This is the main file for the iec61499 network with _TCREST as the top level block

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

#include "_CBCoreSingle.h"


const int NOC_MASTER = 0;

void t0(void* param);

//void task0(_CBCoreSingle_t * c0);

int main() {
	printf("circuitbreaker_mem fbc single startup.\n");
	printf("sizes: %lu\n", sizeof(_CBCoreSingle));
	//mp_init();


	t0(NULL);
	//int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task0(_CBCoreSingle * c0) {
	_CBCoreSinglerun(c0);
}

void task0(_CBCoreSingle * c0) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task0(c0);

		end_time = get_cpu_cycles();
		printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t0(void* param) {
	HEX = 7;
	_CBCoreSingle * c0;
	_CBCoreSingle c;
	c0 = &c; //SPM_BASE;

	_CBCoreSingleinit(c0);

	HEX = 8;

	// if(mp_init_ports() == 0) {
	// 	HEX = 16;
	// 	return;
	// }
	HEX = 9;

	//c0init = 1;
	//while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task0(c0);
}


