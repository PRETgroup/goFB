//This is the main file for the iec61499 network with _TCREST as the top level block

#include "FB__CB_TCREST.h"

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

void task0(_CBCore0_t _SPM * c0);
void task1(_CBCore1_t _SPM * c1);
void task2(_CBCore2_t _SPM * c2);
void task3(_CBCore3_t _SPM * c3);


int main() {
	printf("circuitbreaker_mem fbc tcrest4 startup.\n");
	printf("sizes: %lu, %lu, %lu, %lu\n", sizeof(_CBCore0_t), sizeof(_CBCore1_t), sizeof(_CBCore2_t), sizeof(_CBCore3_t));
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
	//int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task0(_CBCore0_t _SPM * c0) {
	_CBCore0_syncOutputEvents(c0);
	_CBCore0_syncInputEvents(c0);
	_CBCore0_syncOutputData(c0);
	_CBCore0_syncInputData(c0);
	_CBCore0_run(c0);
}

void __attribute__ ((noinline)) timed_task1(_CBCore1_t _SPM * c1) {
	_CBCore1_syncOutputEvents(c1);
	_CBCore1_syncInputEvents(c1);
	_CBCore1_syncOutputData(c1);
	_CBCore1_syncInputData(c1);
	_CBCore1_run(c1);
}

void __attribute__ ((noinline)) timed_task2(_CBCore2_t _SPM * c2) {
	_CBCore2_syncOutputEvents(c2);
	_CBCore2_syncInputEvents(c2);
	_CBCore2_syncOutputData(c2);
	_CBCore2_syncInputData(c2);
	_CBCore2_run(c2);
}

void __attribute__ ((noinline)) timed_task3(_CBCore3_t _SPM * c3) {
	_CBCore3_syncOutputEvents(c3);
	_CBCore3_syncInputEvents(c3);
	_CBCore3_syncOutputData(c3);
	_CBCore3_syncInputData(c3);
	_CBCore3_run(c3);
}

void task0(_CBCore0_t  _SPM * c0) {
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
	HEX = 0x01;
	_CBCore0_t _SPM * c0;
	c0 = mp_alloc(sizeof(_CBCore0_t));

	if(_CBCore0_preinit(c0) != 0 || _CBCore0_init(c0) != 0) {
		HEX = 0x0E;
		return;
	}

	HEX = 0x02;

	if(mp_init_ports() == 0) {
		HEX = 0x0F;
		return;
	}
	HEX = 0x03;

	c0init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 0x04;
	
	task0(c0);
}



void task1(_CBCore1_t  _SPM * c1) {
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
	HEX = 0x11;
	_CBCore1_t _SPM * c1;
	c1 = mp_alloc(sizeof(_CBCore1_t));
	

	if(_CBCore1_preinit(c1) != 0 || _CBCore1_init(c1) != 0) {
		HEX = 0x1E;
		return;
	}

	HEX = 0x12;

	if(mp_init_ports() == 0) {
		HEX = 0x1F;
		return;
	}
	HEX = 0x13;

	c1init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 0x14;
	task1(c1);
}

void task2(_CBCore2_t  _SPM * c2) {
	//taskn runs coren
	do {
		timed_task2(c2);
	} while(1);
}

void t2(void* param) {
	HEX = 0x21;
	_CBCore2_t _SPM * c2;
	c2 = mp_alloc(sizeof(_CBCore2_t));
	
	if(_CBCore2_preinit(c2) != 0 || _CBCore2_init(c2) != 0) {
		HEX = 0x2E;
		return;
	}

	HEX = 0x22;

	if(mp_init_ports() == 0) {
		HEX = 0x2F;
		return;
	}
	HEX = 0x23;

	c2init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 0x24;
	
	task2(c2);
}


void task3(_CBCore3_t  _SPM * c3) {
	//taskn runs coren
	do {
		timed_task3(c3);
	} while(1);
}

void t3(void* param) {
	HEX = 0x31;
	_CBCore3_t _SPM * c3;
	c3 = mp_alloc(sizeof(_CBCore3_t));
	//c3 = &c; //SPM_BASE;

	if(_CBCore3_preinit(c3) != 0 || _CBCore3_init(c3) != 0) {
		HEX = 0x3E;
		return;
	}

	HEX = 0x32;

	if(mp_init_ports() == 0) {
		HEX = 0x3E;
		return;
	}
	HEX = 0x33;

	c3init = 1;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 0x34;
	
	task3(c3);
}