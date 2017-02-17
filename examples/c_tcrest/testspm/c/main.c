#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

typedef struct {
	int y[500];
} data_t;

typedef struct {
	data_t t;
} container_t;

void data_set(int _SPM *arr) {
	for(int i=0; i < 500; i++) {
		arr[i] = i;
	}
}

void data_check(int _SPM *arr) {
	for(int i=0; i < 500; i++) {
		if(arr[i] != i) {
			printf("ERROR at i:%i, value is %i\n", i, arr[i]);
		}
	}
}





int main() {
	container_t _SPM *c;
	c = SPM_BASE;
	
	// printf("Setting *u...");
	// u->t.x = 1;
	// printf("Done\n");

	printf("Testing spm...");
	data_set(c->t.y);
	data_check(c->t.y);
	
	printf("Done\n");
	return 0;
}