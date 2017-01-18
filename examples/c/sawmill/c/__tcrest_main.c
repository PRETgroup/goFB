//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

volatile _UNCACHED int c0init = 0;
volatile _UNCACHED int c1init = 0;
volatile _UNCACHED int c2init = 0;
volatile _UNCACHED int c3init = 0;

void t1(void* param);
void t2(void* param);
void t3(void* param);

void task0(struct _Core0 _SPM * c0);
void task1(struct _Core1 _SPM * c1);
void task2(struct _Core2 _SPM * c2);
void task3(struct _Core3 _SPM * c3);

////void _SPM * mp_alloc(const size_t size)

// int main() {
// 	printf("lets test.\n");

// 	struct _Core1 _SPM * c1;
// 	c1 = SPM_BASE+8;
// 	// if(c1 == NULL) {
// 	// 	printf("c1 is NULL, quitting\n");
// 	// 	return 1;
// 	// }
// 	//printf("c1 not null, size:%lu\n", sizeof(struct _Core1));

// 	if(_Core1_preinit((struct _Core1 *)c1) != 0) {
// 		printf("preinit failed\n");
// 		return 1;
// 	}
// 	printf("preinit success\n");

// 	if(_Core1_init((struct _Core1 *)c1) != 0) {
// 		printf("init failed\n");
// 		return 1;
// 	}
// 	printf("init success\n");


// 	// _Core1_preinit(&my_TCREST.c1);
// 	// _Core1_init(&my_TCREST.c1);

// 	task1(c1);


// 	// if(mp_init_ports() == 0) {
// 	// 	printf("Failed to initalize all NoC ports\n");
// 	// 	return 1;
// 	// }
// 	// printf("mp_init_ports");

// 	// while(c1init == 0 || c2init == 0 || c3init == 0);

// 	// printf("everything's initialised, now running program!\n\n");
// 	// c0init = 1;

// 	// task0(c0);
// 	// LED = 1;
// 	// int* res;
// 	// //corethread_join(core1, (void**)&res);

// 	return 0;
// }

int main() {
	printf("sawmill tcrest4 startup.\n");
	mp_init();
	printf("Starting t1,t2,t3 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t1, NULL);
	corethread_t core2 = 2;
	corethread_create(&core2, &t2, NULL);
	corethread_t core3 = 3;
	corethread_create(&core3, &t3, NULL);
	printf("Started t1,t2,t3\n");

	struct _Core0 _SPM * c0;
	c0 = SPM_BASE+8;

	_Core0_preinit(c0);
	if(_Core0_init(c0) != 0) {
		printf("Core0 init failed\n");
	}

	printf("SPM_BASE: 0x%.8x\n", SPM_BASE);
	printf("NOC_SPM_BASE: 0x%.8x\n", NOC_SPM_BASE);

	unsigned int _SPM * testx;
	testx = mp_alloc(sizeof(unsigned int));
	printf("testx at 0x%.8x\n", testx);
	printf("c0 at : 0x%.8x\n", c0);
	printf("c0 saw1rx at : 0x%.8x\n", &c0->saw1rx);
	printf("c0 saw1rx chan at : 0x%.8x\n", &c0->saw1rx.chan);
	printf("c0 saw1rx chan contains : 0x%.8x\n", c0->saw1rx.chan);
	printf("c0 saw1rx chan recv addr at : 0x%.8x\n", c0->saw1rx.chan->recv_addr);

	// _Core1_preinit(&my_TCREST.c1);
	// _Core1_init(&my_TCREST.c1);
	printf("init t0\n");

	if(mp_init_ports() == 0) {
		printf("Failed to initalize all NoC ports\n");
		return 1;
	}
	printf("mp_init_ports\n");

	while(c1init == 0 || c2init == 0 || c3init == 0);

	printf("everything's initialised, now running program!\n\n");
	c0init = 1;

	task0(c0);
	LED = 1;
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(struct _Core0 _SPM * c0) {
	_Core0_syncEvents(c0);
	_Core0_syncData(c0);
	_Core0_run(c0);
}

void task0(struct _Core0 _SPM * c0) {
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

void __attribute__ ((noinline)) timed_task1(struct _Core1 _SPM * c1) {
	_Core1_syncEvents(c1);
	_Core1_syncData(c1);
	_Core1_run(c1);
}

void task1(struct _Core1 _SPM * c1) {
	do {
		timed_task1(c1);
	} while(1);
}

void t1(void* param) {
	mp_init();
	HEX = 7;
	struct _Core1 _SPM * c1;
	c1 = SPM_BASE+8;

	// _Core0_preinit(c0);
	// _Core0_init(c0);

	_Core1_preinit(c1);
	_Core1_init(c1);
	HEX = 8;

	mp_init_ports();
	c1init = 1;
	HEX = 9;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);
	HEX = 10;
	task1(c1);

	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}

void task2(struct _Core2 _SPM * c2) {
	//taskn runs coren
	do {
		_Core2_syncEvents(c2);
		_Core2_syncData(c2);
		_Core2_run(c2);
	} while(1);
}

void t2(void* param) {
	mp_init();
	HEX = 7;
	struct _Core2 _SPM * c2;
	c2 = SPM_BASE+8;

	_Core2_preinit(c2);
	_Core2_init(c2);
	HEX = 8;
	mp_init_ports();
	
	c2init = 1;
	HEX = 9;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);
	HEX = 10;
	task2(c2);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}


void task3(struct _Core3 _SPM * c3) {
	//taskn runs coren
	do {
		_Core3_syncEvents(c3);
		_Core3_syncData(c3);
		_Core3_run(c3);
	} while(1);
}

void t3(void* param) {
	mp_init();
	HEX = 7;
	struct _Core3 _SPM * c3;
	c3 = SPM_BASE+8;

	_Core3_preinit(c3);
	_Core3_init(c3);
	HEX = 8;
	mp_init_ports();

	c3init = 1;
	HEX = 9;
	while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);
	HEX = 10;
	task3(c3);
	
	int ret = 0;
	LED = 1;
  	corethread_exit(&ret);
	return;
}