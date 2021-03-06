// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB

// This file represents the interface of Function Block Top
#ifndef TOP_H_
#define TOP_H_

#include "fbtypes.h"
#include "util.h"



//This is a CFB, so we need the #includes for the child blocks embedded here
#include "FB_Ticksource.h"
#include "FB_Generator.h"
#include "FB_Generator.h"
#include "FB_IntegralController.h"
#include "FB_IntegralController.h"
#include "FB_Load.h"
#include "FB_Load.h"
#include "FB_Tieline.h"
#include "FB_LfcPrint.h"




//this block had no input events


//this block had no output events


typedef struct {
    //input events
	

    //output events
	

    //input vars
	
    //output vars
	
	//any internal vars (BFBs only)
    
	//any child FBs (CFBs only)
	Ticksource_t ticksource;
	Generator_t gen1;
	Generator_t gen2;
	IntegralController_t ic1;
	IntegralController_t ic2;
	Load_t load1;
	Load_t load2;
	Tieline_t tieline;
	LfcPrint_t print;
	
	//resource vars
	
	//resources (Devices only)
	
	//state and trigger (BFBs only)
	
	

	

} Top_t;

//all FBs get a preinit function
int Top_preinit(Top_t  *me);

//all FBs get an init function
int Top_init(Top_t  *me);

//all FBs get a run function
void Top_run(Top_t  *me);

//composite/resource/device FBs get sync functions
void Top_syncOutputEvents(Top_t  *me);
void Top_syncInputEvents(Top_t  *me);

void Top_syncOutputData(Top_t  *me);
void Top_syncInputData(Top_t  *me);



#endif
