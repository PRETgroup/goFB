// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block topCFB1
#ifndef TOPCFB1_H_
#define TOPCFB1_H_

#include "fbtypes.h"

//This is a CFB, so we need the #includes for the child blocks embedded here
#include "container.h"
#include "container.h"


//this block had no input events


//this block had no output events


typedef struct {
    //input events
	

    //output events
	

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	container_t cf1;
	container_t cf2;
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	
} topCFB1_t;

//all FBs get a preinit function
int topCFB1_preinit(topCFB1_t  *me);

//all FBs get an init function
int topCFB1_init(topCFB1_t  *me);

//all FBs get a run function
void topCFB1_run(topCFB1_t  *me);

//composite/resource/device FBs get sync functions
void topCFB1_syncOutputEvents(topCFB1_t  *me);
void topCFB1_syncInputEvents(topCFB1_t  *me);

void topCFB1_syncOutputData(topCFB1_t  *me);
void topCFB1_syncInputData(topCFB1_t  *me);



#endif
