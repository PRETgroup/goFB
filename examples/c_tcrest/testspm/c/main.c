#include <machine/spm.h>
#include <machine/patmos.h>
#include <stdio.h>

#include "main.h"

typedef struct {
	int x;
	int y[3];
} data_t;

typedef struct {
	data_t t;
} container_t;

void data_set(int size, int _SPM *arr) {
	for(int i=0; i < size; i++) {
		arr[i] = i;
	}
}

void data_augment(data_t _SPM *d) {
	d->y[2] = 40;
}

int main() {
	container_t _SPM *c;
	c = SPM_BASE;
	
	// printf("Setting *u...");
	// u->t.x = 1;
	// printf("Done\n");

	printf("Setting *c...");
	c->t.x = 1;
	data_set(3, c->t.y);
	data_augment(&c->t);
	printf("Done\n");

	printf("c->t.x = %i\n", c->t.x);
	printf("c->t.y[0] = %i\n", c->t.y[0]);
	printf("c->t.y[1] = %i\n", c->t.y[1]);
	printf("c->t.y[2] = %i\n", c->t.y[2]);
	return 0;
}