//This is the main file for the iec61499 network with _TCREST as the top level block

#include "Resource__CBCore0.h"
#include "Resource1__CBCore1.h"
#include "Resource2__CBCore2.h"
#include "Resource3__CBCore3.h"

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

void task0(Resource__CBCore0 * c0);
void task1(Resource1__CBCore1 * c1);
void task2(Resource2__CBCore2 * c2);
void task3(Resource3__CBCore3 * c3);


int main() {
	printf("circuitbreaker_mem fbc tcrest4 startup.\n");
	printf("sizes: %lu, %lu, %lu, %lu\n", sizeof(Resource__CBCore0), sizeof(Resource1__CBCore1), sizeof(Resource2__CBCore2), sizeof(Resource3__CBCore3));
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

void __attribute__ ((noinline)) timed_task0(Resource__CBCore0 * c0) {
	Resource__CBCore0run(c0);
}

void __attribute__ ((noinline)) timed_task1(Resource1__CBCore1 * c1) {
	Resource1__CBCore1run(c1);
}

void __attribute__ ((noinline)) timed_task2(Resource2__CBCore2 * c2) {
	Resource2__CBCore2run(c2);
}

void __attribute__ ((noinline)) timed_task3(Resource3__CBCore3 * c3) {
	Resource3__CBCore3run(c3);
}

void task0(Resource__CBCore0 * c0) {
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
	Resource__CBCore0 * c0;
	Resource__CBCore0 c;
	c0 = &c; //SPM_BASE;

	Resource__CBCore0init(c0);

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



void task1(Resource1__CBCore1 * c1) {
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
	Resource1__CBCore1 * c1;
	Resource1__CBCore1 c;
	c1 = &c; //SPM_BASE;

	Resource1__CBCore1init(c1);

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

void task2(Resource2__CBCore2 * c2) {
	//taskn runs coren
	do {
		timed_task2(c2);
	} while(1);
}

void t2(void* param) {
	HEX = 0x21;
	Resource2__CBCore2 * c2;
	Resource2__CBCore2 c;
	c2 = &c; //SPM_BASE;

	
	Resource2__CBCore2init(c2);

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


void task3(Resource3__CBCore3 * c3) {
	//taskn runs coren
	do {
		timed_task3(c3);
	} while(1);
}

void t3(void* param) {
	HEX = 0x31;
	Resource3__CBCore3 * c3;
	Resource3__CBCore3 c;
	c3 = &c; //SPM_BASE;

	
	Resource3__CBCore3init(c3);

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