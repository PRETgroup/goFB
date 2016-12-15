{{define "FBheader"}}// This file has been automatically generated by goFB and should not be edited by hand
// Transpiler written by Hammond Pearce and available at github.com/kiwih/goFB
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}
// This file represents the interface of Function Block {{$block.Name}}
#ifndef {{strToUpper $block.Name}}_H_
#define {{strToUpper $block.Name}}_H_

#include "fbtypes.h"

{{if $block.CompositeFB}}{{range $currChildIndex, $child := $block.CompositeFB.FBs}}//This is a CFB, so we need the #includes for the child blocks embedded here
#include "{{$child.Type}}.h"
{{end}}{{end}}{{if $block.BasicFB}}//This is a BFB, so we need an enum type for the state machine
enum {{$block.Name}}_states { {{range $index, $state := $block.BasicFB.States}}{{if $index}}, {{end}}STATE_{{$block.Name}}_{{$state.Name}}{{end}} };
{{end}}

{{if $block.EventInputs}}union {{$block.Name}}InputEvents {
	struct {
	{{if $block.EventInputs}}{{range $index, $event := $block.EventInputs.Events}}	UDINT {{$event.Name}} : 1;
	{{end}}{{end}}} event;
	UDINT events[{{if $block.EventInputs}}{{add (div (len $block.EventInputs.Events) 32) 1}}{{else}}1{{end}}];
};
{{else}}//this block had no input events
{{end}}

{{if $block.EventOutputs}}union {{$block.Name}}OutputEvents {
	struct {
	{{if $block.EventOutputs}}{{range $index, $event := $block.EventOutputs.Events}}	UDINT {{$event.Name}} : 1;
	{{end}}{{end}}} event;
	UDINT events[{{if $block.EventOutputs}}{{add (div (len $block.EventOutputs.Events) 32) 1}}{{else}}1{{end}}];
};
{{else}}//this block had no output events
{{end}}

struct {{$block.Name}} {
    //input events
	{{if $block.EventInputs}}union {{$block.Name}}InputEvents inputEvents;{{end}}

    //output events
	{{if $block.EventOutputs}}union {{$block.Name}}OutputEvents outputEvents;{{end}}

    //input vars
	{{if $block.InputVars}}{{range $index, $var := $block.InputVars.Variables}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
    {{end}}{{end}}
    //output vars
	{{if $block.OutputVars}}{{range $index, $var := $block.OutputVars.Variables}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
    {{end}}{{end}}
	//any internal vars (BFBs only)
    {{if $block.BasicFB}}{{if $block.BasicFB.InternalVars}}{{range $varIndex, $var := $block.BasicFB.InternalVars.Variables}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
    {{end}}{{end}}{{end}}
	//any child FBs (CFBs only)
	{{if $block.CompositeFB}}{{range $currChildIndex, $child := $block.CompositeFB.FBs}}struct {{$child.Type}} {{$child.Name}};
	{{end}}{{end}}
	//resource vars
	{{if $block.ResourceVars}}{{range $index, $var := $block.ResourceVars}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
	{{end}}{{end}}
	//state and trigger (BFBs only)
	{{if $block.BasicFB}}enum {{$block.Name}}_states _state; //stores current state
	BOOL _trigger; //indicates if a state transition has occured this tick
	{{end}}
};

//all FBs get an init function
void {{$block.Name}}_init(struct {{$block.Name}} *me);

//all FBs get a run function
void {{$block.Name}}_run(struct {{$block.Name}} *me);

{{if $block.CompositeFB}}//composite FBs get sync functions
void {{$block.Name}}_syncEvents(struct {{$block.Name}} *me);
void {{$block.Name}}_syncData(struct {{$block.Name}} *me);{{end}}{{if $block.BasicFB}}{{$basicFB := $block.BasicFB}}
{{if $basicFB.Algorithms}}//basic FBs have a number of algorithm functions
{{range $algIndex, $alg := $basicFB.Algorithms}}
void {{$block.Name}}_{{$alg.Name}}(struct {{$block.Name}} *me);
{{end}}{{end}}{{end}}

#endif
{{end}}