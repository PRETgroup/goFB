//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

//put a copy of the top level block into global memory
_TCREST_t my_TCREST;

const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;

void t0(void* param);
void t1(void* param);

void task0(void* param);
void task1(void* param);

int main() {
	printf("testcomm startup.\n");

	printf("Starting t1 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	printf("Started t1\n");
	t0(NULL);
}

void __attribute__ ((noinline)) timed_task() {
	_Core0_syncEvents(&my_TCREST.rx_core);
	_Core0_syncData(&my_TCREST.rx_core);
	_Core0_run(&my_TCREST.rx_core);
}

void task0(void* param) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task();

		end_time = get_cpu_cycles();
		//printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}


void task1(void* param) {
	//task1 runs core1

	unsigned int tickCount = 0;
	do {
		_Core1_syncEvents(&my_TCREST.tx_core);
		_Core1_syncData(&my_TCREST.tx_core);
		_Core1_run(&my_TCREST.tx_core);
	} while(1);
}

void t1(void* param) {

	_Core1_init(&my_TCREST.tx_core);
	mp_init_ports();

	c1init = 1;
	while(c0init != 1 && c1init != 1);

	task1(NULL);

}


void t0(void* param) {

	_Core0_init(&my_TCREST.rx_core);
	mp_init_ports();

	c0init = 1;
	while(c0init != 1 && c1init != 1);

	task0(NULL);
}