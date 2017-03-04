//This is the main file for the iec61499 network with _TCREST as the top level block

#include "_TCREST.h"

#ifndef PROGS_PER_CORE
#define PROGS_PER_CORE 1
#endif

//put a copy of the top level block into global memory
//struct _TCREST my_TCREST;

const int NOC_MASTER = 0;

void t(void* param);

void task(_Core_t * c0);


int main() {
	printf("pid_mem tcrest4 startup.\n");
	printf("sizes: %lu, %lu, %lu, %lu\n", sizeof(_Core_t), sizeof(_Core_t), sizeof(_Core_t), sizeof(_Core_t));
	mp_init();
	printf("Starting t1,t2,t3 and initialising my_TCREST...\n");
	corethread_t core1 = 1;
	corethread_create(&core1, &t, NULL);
	corethread_t core2 = 2;
	corethread_create(&core2, &t, NULL);
	corethread_t core3 = 3;
	corethread_create(&core3, &t, NULL);
	printf("Started t1,t2,t3\n");

	t(NULL);
	int* res;
	//corethread_join(core1, (void**)&res);

	return 0;
}

void __attribute__ ((noinline)) timed_task(_Core_t * c) {
	int i;
	_Core_t _SPM *c_spm;
	c_spm = SPM_BASE;
	for(i=0; i < PROGS_PER_CORE; i++) {
		spm_copy_from_ext(c_spm, &c[i], sizeof(_Core_t));
		_Core_syncOutputEvents(c_spm);
		_Core_syncInputEvents(c_spm);
		_Core_syncOutputData(c_spm);
		_Core_syncInputData(c_spm);
		_Core_run(c_spm);
		spm_copy_to_ext(&c[i], c_spm, sizeof(_Core_t));
	}
}

void task(_Core_t * c) {
	//task0 runs core0
	unsigned int tickCount = 0;

	unsigned long long start_time;
	unsigned long long end_time;

	do {
		start_time = get_cpu_cycles();

		timed_task(c);

		end_time = get_cpu_cycles();
		// printf("%4d\t\t%lld\n", tickCount, end_time-start_time-3);

		tickCount++;
	} while(1);
}

void t(void* param) {
	HEX = 7;

	_Core_t c[PROGS_PER_CORE];
	_Core_t _SPM *c_spm;

	int i;
	for(i=0; i<PROGS_PER_CORE; i++) {

		c_spm = SPM_BASE;

		spm_copy_from_ext(c_spm, &c[i], sizeof(_Core_t));

		if(_Core_preinit(c_spm) != 0 || _Core_init(c_spm) != 0) {
			HEX = 15;
			return;
		}

		spm_copy_to_ext(&c[i], c_spm, sizeof(_Core_t));
	}
	

	HEX = 8;

	// if(mp_init_ports() == 0) {
	// 	HEX = 16;
	// 	return;
	// }
	HEX = 9;

	// c0init = 1;
	// while(c0init == 0 || c1init == 0 || c2init == 0 || c3init == 0);

	HEX = 10;
	
	task(c);
}

