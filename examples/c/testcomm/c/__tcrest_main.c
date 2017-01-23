//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_Core0.h"
#include "_Core1.h"
//put a copy of the top level block into global memory



const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;

void t1(void* param);

void task1(_Core1_t _SPM * c1);

void __attribute__ ((noinline)) timed_task(_Core0_t _SPM * c0);

int main() {
	mp_init();
	printf("testcomm startup.\n");

	printf("Starting t1 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	printf("Started t1\n");

	_Core0_t _SPM * c0;
	c0 = SPM_BASE;
	
	c0->rx.outputEvents.events[0] = 0;

	printf("size is %lu\n", sizeof(*c0));

	int x;
	
	x = _Core0_preinit(c0);
	if(x != 0) {
		printf("preinit returns 1\n");
		//return 1;
	}
	x = _Core0_init(c0);
	if(x != 0) {
		printf("init returns 1\n");
	}
	printf("init core0\n");

	if(mp_init_ports() == 0) {
		printf("Failed to initalize all NoC ports\n");
		return 1;
	}
	printf("mp_init_ports\n");

	while(c1init != 1);

	printf("everything's initialised, now running program!\n\n");
	c0init = 1;

	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task(c0);

		end_time = get_cpu_cycles();
		//printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
	LED = 1;
	// int* res;
	// corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core0_t _SPM * c0) {
	_Core0_syncEvents(c0);
	_Core0_syncData(c0);
	_Core0_run(c0);
}

void task1(_Core1_t _SPM * c1) {
	//task1 runs core1

	unsigned int tickCount = 0;
	do {
		_Core1_syncEvents(c1);
		_Core1_syncData(c1);
		_Core1_run(c1);
	} while(1);
}

void t1(void* param) {
				HEX = HEX + 1;
							HEX = HEX + 1;
										HEX = HEX + 1;
													HEX = HEX + 1;
	mp_init();
	_Core1_t _SPM * c1;
	c1 = SPM_BASE;
	_Core1_preinit(c1);
	_Core1_init(c1);
	mp_init_ports();

	c1init = 1;
	while(c0init != 1 && c1init != 1);

	task1(c1);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}