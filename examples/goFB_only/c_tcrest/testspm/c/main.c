#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

#define MAX_ARRAY 10

typedef struct {
	unsigned int x: 1;
	unsigned int y: 1;
} pt_t;

typedef struct {
	pt_t pts[MAX_ARRAY];
} container_t;

void spm_data_set(container_t _SPM *c) {
	for(int i=0; i < MAX_ARRAY; i++) {
		c->pts[i].x = i % 2 == 1;
		c->pts[i].y = i % 2 == 0;
	}
}

void spm_data_check(container_t _SPM *c) {
	for(int i=0; i < MAX_ARRAY; i++) {
		if(c->pts[i].x != (i % 2 == 1) || c->pts[i].y != (i % 2 == 0)) {
			printf("ERROR at i:%i, stored:%2i,%2i should be %2i,%2i\n", i, c->pts[i].x, c->pts[i].y, (i % 2 == 1), (i % 2 == 0));
		}
	}
}

void mem_data_set(container_t *c) {
	for(int i=0; i < MAX_ARRAY; i++) {
		c->pts[i].x = i % 2 == 1;
		c->pts[i].y = i % 2 == 0;
	}
}

void mem_data_check(container_t *c) {
	for(int i=0; i < MAX_ARRAY; i++) {
		if(c->pts[i].x != (i % 2 == 1) || c->pts[i].y != (i % 2 == 0)) {
			printf("ERROR at i:%i, stored:%2i,%2i should be %2i,%2i\n", i, c->pts[i].x, c->pts[i].y, (i % 2 == 1), (i % 2 == 0));
		}
	}
}

int main() {
	container_t _SPM *c_spm;
	c_spm = SPM_BASE;

	container_t *c_mem;
	container_t c;
	c_mem = &c;
	
	// printf("Setting *u...");
	// u->t.x = 1;
	// printf("Done\n");

	printf("Testing spm...\n");
	spm_data_set(c_spm);
	spm_data_check(c_spm);

	printf("Testing main mem...\n");
	mem_data_set(c_mem);
	mem_data_check(c_mem);
	
	printf("Done\n");
	return 0;
}