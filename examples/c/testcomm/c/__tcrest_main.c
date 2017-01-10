//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

#include <machine/patmos.h>
#include "libcorethread/corethread.h"
#include "libmp/mp.h"

//put a copy of the top level block into global memory
struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;

void t1(void* param);

void task0(void* param);
void task1(void* param);

#define LED ( *( ( volatile _IODEV unsigned * )	0xF0090000 ) )

int main() {
	printf("testcomm startup.\n");

	printf("Starting t1 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	printf("Started t1\n");
	_Core0_init(&my_TCREST.rx_core);
	printf("init core0\n");

	if(mp_init_ports() == 0) {
		printf("Failed to initalize all NoC ports\n");
		return 1;
	}
	printf("mp_init_ports\n");

	while(c1init != 1);

	printf("everything's initialised, now running program!\n\n");
	c0init = 1;

	task0(NULL);

	int* res;
	corethread_join(core1, (void**)&res);
	
	return 0;
}

void task0(void* param) {
	//task0 runs core0

	int count = 0;
	unsigned int tickCount = 0;
	do {
		if(count++ > 10000) {
			LED = !(LED);
			count = 0;
		}

		_Core0_syncEvents(&my_TCREST.rx_core);
		_Core0_syncData(&my_TCREST.rx_core);
		_Core0_run(&my_TCREST.rx_core);

	} while(tickCount++ < 100000);
}

void task1(void* param) {
	//task1 runs core1

	int count = 0;
	unsigned int tickCount = 0;

	do {
		if(count++ > 10000) {
			LED = !(LED);
			count = 0;
		}

		_Core1_syncEvents(&my_TCREST.tx_core);
		_Core1_syncData(&my_TCREST.tx_core);
		_Core1_run(&my_TCREST.tx_core);
	} while(tickCount++ < 100000);
}

void t1(void* param) {

	_Core1_init(&my_TCREST.tx_core);
	mp_init_ports();

	c1init = 1;
	while(c0init != 1 && c1init != 1);

	task1(NULL);
	
	int ret = 0;
  	corethread_exit(&ret);
	return;
}