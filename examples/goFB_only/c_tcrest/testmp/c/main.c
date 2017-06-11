#include <machine/spm.h>
#include <machine/patmos.h>
#include "libcorethread/corethread.h"
#include "libmp/mp.h"
#include "main.h"
#include "libmp/mp_internal.h"

extern volatile _UNCACHED chan_info_t chan_info[MAX_CHANNELS];

typedef struct {
	qpd_t * chan;
} chan_cont_t;

typedef struct {
	chan_cont_t c;
} consumer_t;

typedef struct {
	chan_cont_t c;
} producer_t;

const int NOC_MASTER = 0;

#define HEX 		( *( ( volatile _IODEV unsigned * )	0xF0070000 ) )
#define SWITCHES 	( *( ( volatile _IODEV unsigned * )	0xF0060000 ) )

void rx_tick(chan_cont_t _SPM *c) {
	int data;
	int success;

	success = mp_nbrecv(c->chan);
		
	if(success) {
		data = *((volatile unsigned int _SPM*)c->chan->read_buf);

		printf("Received %d...", data);

		do {
			success = mp_nback(c->chan);
		} while(success == 0);

		printf("acknowledged\n");
	}
}

void consumer_tick(consumer_t _SPM *ct) {
	

	do {
		rx_tick(&ct->c);
	} while(1);
}

void consumer() {

	consumer_t _SPM *ct;
	ct = SPM_BASE;

	ct->c.chan = mp_create_qport(1, SINK, sizeof(unsigned int), 1);
	if(ct->c.chan == NULL) {
		printf("mp_create_qport failed\n");
		return;
	}

	printf("trying to init ports\n");

	if(mp_init_ports() == 0) {
		printf("mp_init_ports failed\n");
		return;
	}
	printf("awaiting data reception\n");
	consumer_tick(ct);

}

void producer(void* param) {

	producer_t _SPM *pt;
	pt = SPM_BASE;

	pt->c.chan = mp_create_qport(1, SOURCE, sizeof(unsigned int), 1);
	HEX = 1;
	if(pt->c.chan == NULL) {
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
		*((volatile unsigned int _SPM *)pt->c.chan->write_buf) = count;
		if(SWITCHES) {
			if(mp_nbsend(pt->c.chan)) {
				HEX = ++count;
			}
		}
	} while(1);
}


int main() {
	mp_init();

	corethread_t core1 = 1;
	corethread_create(&core1, &producer, NULL);

	consumer();
}