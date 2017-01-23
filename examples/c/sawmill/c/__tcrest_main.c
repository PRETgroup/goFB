//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;
volatile _UNCACHED int c2init = 0;
volatile _UNCACHED int c3init = 0;

void t0(void* param);
void t1(void* param);
void t2(void* param);
void t3(void* param);

void task0(_Core0_t _SPM * c0);
void task1(_Core1_t _SPM * c1);
void task2(_Core2_t _SPM * c2);
void task3(_Core3_t _SPM * c3);


int main() {
	printf("sawmill tcrest4 startup.\n");
	//mp_init();
	printf("Starting t1,t2,t3 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	corethread_t core2 = 2;
	corethread_create(&core2, &t2, NULL);
	corethread_t core3 = 3;
	corethread_create(&core3, &t3, NULL);
	printf("Started t1,t2,t3\n");

	t0(NULL);
	LED = 1;
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core0_t _SPM * c0) {
	_Core0_syncEvents(c0);
	_Core0_syncData(c0);
	_Core0_run(c0);
}

void task0(_Core0_t _SPM * c0) {
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
}

void t0(void* param) {
	HEX = 7;
	_Core0_t _SPM * c0;
	c0 = SPM_BASE;

	if(_Core0_preinit(c0) != 0 || _Core0_init(c0) != 0) {
		HEX = 15;
		return;
	}

	HEX = 8;

	if(mp_init_ports() == 0) {
		HEX = 16;
		return;
	}
	HEX = 9;

	c0init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task0(c0);
}

void __attribute__ ((noinline)) timed_task1(_Core1_t _SPM * c1) {
	_Core1_syncEvents(c1);
	_Core1_syncData(c1);
	_Core1_run(c1);
}

void task1(_Core1_t _SPM * c1) {
	do {
		timed_task1(c1);
	} while(1);
}

void t1(void* param) {
	HEX = 7;
	_Core1_t _SPM * c1;
	c1 = SPM_BASE;

	if(_Core1_preinit(c1) != 0 || _Core1_init(c1) != 0) {
		HEX = 15;
		return;
	}

	HEX = 8;

	if(mp_init_ports() == 0) {
		HEX = 16;
		return;
	}
	HEX = 9;

	c1init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	return;
	task1(c1);
}

void task2(_Core2_t _SPM * c2) {
	//taskn runs coren
	do {
		_Core2_syncEvents(c2);
		_Core2_syncData(c2);
		_Core2_run(c2);
	} while(1);
}

void t2(void* param) {
	HEX = 7;
	_Core2_t _SPM * c2;
	c2 = SPM_BASE;

	if(_Core2_preinit(c2) != 0 || _Core2_init(c2) != 0) {
		HEX = 15;
		return;
	}

	HEX = 8;

	if(mp_init_ports() == 0) {
		HEX = 16;
		return;
	}
	HEX = 9;

	c2init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task2(c2);
}


void task3(_Core3_t _SPM * c3) {
	//taskn runs coren
	do {
		_Core3_syncEvents(c3);
		_Core3_syncData(c3);
		_Core3_run(c3);
	} while(1);
}

void t3(void* param) {
	HEX = 7;
	_Core3_t _SPM * c3;
	c3 = SPM_BASE;

	if(_Core3_preinit(c3) != 0 || _Core3_init(c3) != 0) {
		HEX = 15;
		return;
	}

	HEX = 8;

	if(mp_init_ports() == 0) {
		HEX = 16;
		return;
	}
	HEX = 9;

	c3init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task3(c3);
}