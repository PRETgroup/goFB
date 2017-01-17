//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

//put a copy of the top level block into global memory
struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;
volatile _UNCACHED int c2init = 0;
volatile _UNCACHED int c3init = 0;

void t1(void* param);
void t2(void* param);
void t3(void* param);

void task0(void* param);
void task1(void* param);
void task2(void* param);
void task3(void* param);

int main() {
	printf("sawmill tcrest4 startup.\n");

	printf("Starting t1,t2,t3 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	corethread_t core2 = 2;
	corethread_create(&core2, &t2, NULL);
	corethread_t core3 = 3;
	corethread_create(&core3, &t3, NULL);
	printf("Started t1,t2,t3\n");

	_Core0_preinit(&my_TCREST.c0);
	_Core0_init(&my_TCREST.c0);
	printf("init t0\n");

	if(mp_init_ports() == 0) {
		printf("Failed to initalize all NoC ports\n");
		return 1;
	}
	printf("mp_init_ports\n");

	while(c1init == 0 || c2init == 0 || c3init == 0);

	printf("everything's initialised, now running program!\n\n");
	c0init = 1;

	task0(NULL);
	LED = 1;
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task() {
	_Core0_syncEvents(&my_TCREST.c0);
	_Core0_syncData(&my_TCREST.c0);
	_Core0_run(&my_TCREST.c0);
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
	//taskn runs coren

	unsigned int tickCount = 0;
	do {
		_Core1_syncEvents(&my_TCREST.c1);
		_Core1_syncData(&my_TCREST.c1);
		_Core1_run(&my_TCREST.c1);
	} while(1);
}

void t1(void* param) {
	_Core1_preinit(&my_TCREST.c1);
	_Core1_init(&my_TCREST.c1);

	mp_init_ports();
	c1init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);
	task1(NULL);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}

void task2(void* param) {
	//taskn runs coren

	unsigned int tickCount = 0;
	do {

		_Core2_syncEvents(&my_TCREST.c2);
		_Core2_syncData(&my_TCREST.c2);
		_Core2_run(&my_TCREST.c2);
	} while(1);
}

void t2(void* param) {
	_Core2_preinit(&my_TCREST.c2);
	_Core2_init(&my_TCREST.c2);
	
	mp_init_ports();
	
	c2init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);
	
	task2(NULL);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}


void task3(void* param) {
	//taskn runs coren

	unsigned int tickCount = 0;
	do {
		_Core3_syncEvents(&my_TCREST.c3);
		_Core3_syncData(&my_TCREST.c3);
		_Core3_run(&my_TCREST.c3);
	} while(1);
}

void t3(void* param) {

	_Core3_preinit(&my_TCREST.c3);
	_Core3_init(&my_TCREST.c3);
	
	mp_init_ports();

	c3init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	task3(NULL);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}