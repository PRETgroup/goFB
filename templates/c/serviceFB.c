{{define "serviceFB"}}// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}{{$basicFB := $block.BasicFB}}{{$tcrestUsingSPM := .TcrestUsingSPM}}{{$tcrestSmartSPM := .TcrestSmartSPM}}
// This file represents the implementation of the Basic Function Block for {{$block.Name}}
#include "FB_{{$block.Name}}.h"

{{template "_fbinit" .}}

/* {{$block.Name}}_run() executes a single tick of an
 * instance of {{$block.Name}}. As it is a SIFB, the execution code is provided
 */
void {{$block.Name}}_run({{$block.Name}}_t {{if or $tcrestUsingSPM $tcrestSmartSPM}}_SPM{{end}} *me) {
	//Code provided in SIFB
	{{if $block.ServiceFB}}{{if $block.ServiceFB.Autogenerate}}{{$block.ServiceFB.Autogenerate.RunText}}{{end}}{{end}}
}


{{end}}