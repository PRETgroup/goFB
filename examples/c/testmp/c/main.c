#include <machine/spm.h>
#include <machine/patmos.h>
#include "libcorethread/corethread.h"
#include "libmp/mp.h"
#include "main.h"
#include "libmp/mp_internal.h"

extern volatile _UNCACHED chan_info_t chan_info[MAX_CHANNELS];

struct consumer {
	qpd_t * chan;
};

const int NOC_MASTER = 0;

#define HEX 		( *( ( volatile _IODEV unsigned * )	0xF0070000 ) )

void consumer() {
	struct consumer _SPM * cps;
	cps = SPM_BASE;

	cps->chan = mp_create_qport(1, SINK, sizeof(unsigned int), 1);

	printf("trying to init ports\n");

	if(mp_init_ports() == 0) {
		printf("mp_init_ports failed\n");
		return;
	}

	int data;
	int success;

	do {
		success = mp_nbrecv(cps->chan);
		
		if(success) {
			data = *((volatile unsigned int _SPM*)cps->chan->read_buf);

			printf("Received %d...", data);

			do {
				success = mp_nback(cps->chan);
			} while(success == 0);

			printf("acknowledged\n");
		}
	} while(1);

}

void producer(void* param) {
	qpd_t * chan = mp_create_qport(1, SOURCE, sizeof(unsigned int), 1);
	HEX = 1;
	if(chan == NULL) {
		 HEX = 2;
		 return;
	} 
	HEX = 3;
	if(mp_init_ports() == 0) {
		HEX = 4;
		return;
	}
	HEX = 5;
	

	unsigned int count = 0;

	do {
		*((volatile unsigned int _SPM *)chan->write_buf) = count;
		if(mp_nbsend(chan)) {
			HEX = ++count;
		}
	} while(1);
}


int main() {
	mp_init();

	corethread_t core1 = 1;
	corethread_create(&core1, &producer, NULL);

	consumer();
}