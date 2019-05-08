//This is the main file for the iec61499 network with _TCREST as the top level block

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

#include "FB__CBCoreSingle.h"


const int NOC_MASTER = 0;

void t0(void* param);

//void task0(_CBCoreSingle_t _SPM * c0);

int main() {
	printf("circuitbreaker_mem single startup.\n");
	printf("sizes: %lu\n", sizeof(_CBCoreSingle_t));

	
	//mp_init();


	t0(NULL);
	//int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task0(_CBCoreSingle_t _SPM * c0) {
	_CBCoreSingle_syncOutputEvents(c0);
	_CBCoreSingle_syncInputEvents(c0);
	_CBCoreSingle_syncOutputData(c0);
	_CBCoreSingle_syncInputData(c0);
	_CBCoreSingle_run(c0);
}

void task0(_CBCoreSingle_t _SPM * c0) {
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
	_CBCoreSingle_t _SPM * c0;
	//_CBCoreSingle_t c;
	c0 = SPM_BASE;

	// printf("01 sizes: %lu\n", sizeof( c.amm1 ));
	// printf("02 sizes: %lu\n", sizeof( c.timer1 ));
	// printf("03 sizes: %lu\n", sizeof( c.cb1 ));
	// printf("04 sizes: %lu\n", sizeof( c.hm1 ));
	// printf("05 sizes: %lu\n", sizeof( c.led1 ));
	// printf("06 sizes: %lu\n", sizeof( c.hm3 ));
	// printf("07 sizes: %lu\n", sizeof( c.led3 ));
	// printf("08 sizes: %lu\n", sizeof( c.amm3 ));
	// printf("09 sizes: %lu\n", sizeof( c.cb3 ));
	// printf("10 sizes: %lu\n", sizeof( c.timer3 ));
	// printf("11 sizes: %lu\n", sizeof( c.hm2 ));
	// printf("12 sizes: %lu\n", sizeof( c.led2 ));
	// printf("13 sizes: %lu\n", sizeof( c.amm2 ));
	// printf("14 sizes: %lu\n", sizeof( c.cb2 ));
	// printf("15 sizes: %lu\n", sizeof( c.timer2 ));
	// printf("16 sizes: %lu\n", sizeof( c.print ));

	if(_CBCoreSingle_preinit(c0) != 0 || _CBCoreSingle_init(c0) != 0) {
		HEX = 15;
		return;
	}

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


