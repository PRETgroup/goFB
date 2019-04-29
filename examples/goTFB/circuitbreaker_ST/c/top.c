// This file has been automatically generated by goFB
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

//This is the main file for the iec61499 network with CfbBreakerController as the top level block

#include "FB_CfbBreakerController.h"
#include <sys/time.h>

// #ifndef MAX_TICKS
// #define MAX_TICKS 100
// #endif



//put a copy of the top level block into global memory
CfbBreakerController_t CfbBreakerController;

int main() {
	if(CfbBreakerController_preinit(&CfbBreakerController) != 0) {
		printf("Failed to preinitialize.");
		return 1;
	}
	if(CfbBreakerController_init(&CfbBreakerController) != 0) {
		printf("Failed to initialize.");
		return 1;
	}
	printf("\n");
	
//this is executing with synchronous semantics
	//printf("\n\n\n");
	//printf("\nTop: %20s   Size: %lu\n", "CfbBreakerController", sizeof(myCfbBreakerController));

	#ifdef PRINT_VALS
	printf("Simulation time,");
	#endif
	
	struct timeval tv1, tv2;
	gettimeofday(&tv1, NULL);

	int tickNum = 1;
	do {
		#ifdef PRINT_VALS
			printf("%f,",(double)tickNum*0.01);
		#endif
		
		CfbBreakerController_syncOutputEvents(&CfbBreakerController);
		CfbBreakerController_syncInputEvents(&CfbBreakerController);

		CfbBreakerController_syncOutputData(&CfbBreakerController);
		CfbBreakerController_syncInputData(&CfbBreakerController);
		
		
		CfbBreakerController_run(&CfbBreakerController);
		#ifdef PRINT_VALS
			printf("\n");
		#endif
	} 
	#ifdef MAX_TICKS
		while(tickNum++ < MAX_TICKS);
	#else
		while(1);
	#endif
	gettimeofday(&tv2, NULL);
	#ifdef PRINT_TIME
	printf ("Total time = %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec) / 1000000 +
         (double) (tv2.tv_sec - tv1.tv_sec));
	#endif
	return 0;

}


