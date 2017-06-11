//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_Core402.h"

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

void task0(_Core402_t * c0);
void task1(_Core402_t * c1);
void task2(_Core402_t * c2);
void task3(_Core402_t * c3);


int main() {
	printf("messagepasser_402_mem tcrest4 startup.\n");
	printf("sizes: %lu, %lu, %lu, %lu\n", sizeof(_Core402_t), sizeof(_Core402_t), sizeof(_Core402_t), sizeof(_Core402_t));
	mp_init();
	printf("Starting t1,t2,t3 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t0, NULL);
	corethread_t core2 = 2;
	corethread_create(&core2, &t2, NULL);
	corethread_t core3 = 3;
	corethread_create(&core3, &t3, NULL);
	printf("Started t1,t2,t3\n");

	t1(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task0(_Core402_t * c0) {
	_Core402_syncOutputEvents(c0);
	_Core402_syncInputEvents(c0);
	_Core402_syncOutputData(c0);
	_Core402_syncInputData(c0);
	_Core402_run(c0);
}

void __attribute__ ((noinline)) timed_task1(_Core402_t * c1) {
	_Core402_syncOutputEvents(c1);
	_Core402_syncInputEvents(c1);
	_Core402_syncOutputData(c1);
	_Core402_syncInputData(c1);
	_Core402_run(c1);
}

void __attribute__ ((noinline)) timed_task2(_Core402_t * c2) {
	_Core402_syncOutputEvents(c2);
	_Core402_syncInputEvents(c2);
	_Core402_syncOutputData(c2);
	_Core402_syncInputData(c2);
	_Core402_run(c2);
}

void __attribute__ ((noinline)) timed_task3(_Core402_t * c3) {
	_Core402_syncOutputEvents(c3);
	_Core402_syncInputEvents(c3);
	_Core402_syncOutputData(c3);
	_Core402_syncInputData(c3);
	_Core402_run(c3);
}

void task0(_Core402_t * c0) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task0(c0);

		end_time = get_cpu_cycles();
		//printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t0(void* param) {
	HEX = 7;
	_Core402_t * c0;
	_Core402_t c;
	c0 = &c; //SPM_BASE;

	if(_Core402_preinit(c0) != 0 || _Core402_init(c0) != 0) {
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



void task1(_Core402_t * c1) {
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task1(c1);

		end_time = get_cpu_cycles();
		printf("%5d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t1(void* param) {
	HEX = 7;
	_Core402_t * c1;
	_Core402_t c;
	c1 = &c; //SPM_BASE;

	if(_Core402_preinit(c1) != 0 || _Core402_init(c1) != 0) {
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
	task1(c1);
}

void task2(_Core402_t * c2) {
	//taskn runs coren
	do {
		timed_task2(c2);
	} while(1);
}

void t2(void* param) {
	HEX = 7;
	_Core402_t * c2;
	_Core402_t c;
	c2 = &c; //SPM_BASE;

	if(_Core402_preinit(c2) != 0 || _Core402_init(c2) != 0) {
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


void task3(_Core402_t * c3) {
	//taskn runs coren
	do {
		timed_task3(c3);
	} while(1);
}

void t3(void* param) {
	HEX = 7;
	_Core402_t * c3;
	_Core402_t c;
	c3 = &c; //SPM_BASE;

	if(_Core402_preinit(c3) != 0 || _Core402_init(c3) != 0) {
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