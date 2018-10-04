{{define "_moduleDeclr"}}
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}
(
		input wire clk,
		
		{{if $block.EventInputs}}
		//input events
		{{range $index, $event := $block.EventInputs}}input wire {{$event.Name}}_eI,
		{{end}}{{end}}
		{{if $block.EventOutputs}}
		//output events
		{{range $index, $event := $block.EventOutputs}}output wire {{$event.Name}}_eO,
		{{end}}{{end}}
		{{if $block.InputVars}}
		//input variables
		{{range $index, $var := $block.InputVars}}input wire {{getVerilogSize $var.Type}} {{$var.Name}}_I,
		{{end}}{{end}}
		{{if $block.OutputVars}}
		//output variables
		{{range $index, $var := $block.OutputVars}}output reg {{getVerilogSize $var.Type}} {{$var.Name}}_O {{if $var.InitialValue}} = {{$var.InitialValue}}{{end}},
		{{end}}{{end}}

		input reset
);
{{end}}