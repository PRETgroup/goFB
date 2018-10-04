{{define "basicFB"}}// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
// Verilog support is EXPERIMENTAL ONLY
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}{{$basicFB := $block.BasicFB}}
// This file represents the Basic Function Block for {{$block.Name}}

//defines for state names used internally
{{range $index, $state := $basicFB.States}}`define STATE_{{$state.Name}} {{$index}}
{{end}}

module FB_{{$block.Name}} {{template "_moduleDeclr" .}}

////BEGIN algorithm functions
{{if $basicFB.Algorithms}}{{range $algIndex, $alg := $basicFB.Algorithms}}
function {{$alg.Name}}

begin
{{compileAlgorithm $block $alg -}}
endfunction{{end}}{{end}}
////END algorithm functions

////BEGIN internal copies of I/O
{{if $block.EventInputs}}//input events
{{range $index, $event := $block.EventInputs}}wire {{$event.Name}};
assign {{$event.Name}} = {{$event.Name}}_eI;
{{end}}{{end}}
{{if $block.EventOutputs}}//output events
{{range $index, $event := $block.EventOutputs}}reg {{$event.Name}};
assign {{$event.Name}}_eO = {{$event.Name}};
{{end}}{{end}}
{{if $block.InputVars}}//input variables
{{range $index, $var := $block.InputVars}}reg {{getVerilogSize $var.Type}} {{$var.Name}} {{if $var.InitialValue}} = {{$var.InitialValue}}{{end}};
{{end}}{{end}}
{{if $block.OutputVars}}//output variables
{{range $index, $var := $block.OutputVars}}reg {{getVerilogSize $var.Type}} {{$var.Name}} {{if $var.InitialValue}} = {{$var.InitialValue}}{{end}};
{{end}}{{end}}
////END internal copies of I/O

////BEGIN internal vars
{{if $basicFB.InternalVars}}{{range $varIndex, $var := $basicFB.InternalVars}}
reg  {{getVerilogSize $var.Type}} {{$var.Name}} {{if $var.InitialValue}} = {{$var.InitialValue}}{{end}}; {{end}}{{end}}
////END internal vars

//STATE variables
reg integer state = `STATE_{{(index $basicFB.States 0).Name}};
reg entered = 1'b0;

always@(posedge clk) begin

	if(reset) begin
		//reset state 
		state = `STATE_{{(index $basicFB.States 0).Name}};

		//reset I/O registers
		{{- range $varIndex, $var := $block.InputVars}}
		{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
		{{- end}}
		{{range $index, $var := $block.OutputVars}}
		{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
		{{- end}}
		//reset internal vars
		{{- range $varIndex, $var := $basicFB.InternalVars}}
		{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
		{{- end}}
	end else begin

		//BEGIN update internal inputs on relevant events
		{{if $block.EventInputs}}{{if $block.InputVars}}{{range $eventIndex, $event := $block.EventInputs}}{{if $event.With}}
		if({{$event.Name}}) begin 
			{{range $varIndex, $var := $block.InputVars}}{{if $event.IsLoadFor $var}}{{$var.Name}} = {{$var.Name}}_I;
			{{end}}{{end}}
		end
		{{end}}{{end}}{{end}}{{end}}
		//END update internal inputs

		//BEGIN ecc 
		case(state) 
			{{range $curStateIndex, $curState := $basicFB.States}}`STATE_{{$curState.Name}}: begin
				{{range $transIndex, $trans := $basicFB.GetTransitionsForState $curState.Name}}{{if $transIndex}}els{{end}}if({{compileTransition $block $trans.Condition}}) begin
					state = `STATE_{{$trans.Destination}};
					entered = 1'b1;
				{{end}}end;
			end {{end}}
		endcase
		//END ecc

		//BEGIN update external outputs on relevant events
		{{if $block.EventOutputs}}{{if $block.OutputVars}}{{range $eventIndex, $event := $block.EventOutputs}}{{if $event.With}}
		if({{$event.Name}}) begin 
			{{range $varIndex, $var := $block.OutputVars}}{{if $event.IsLoadFor $var}}{{$var.Name}}_O = {{$var.Name}};
			{{end}}{{end}}
		end
		{{end}}{{end}}{{end}}{{end}}
		//END update external outputs
	end
end
endmodule{{end}}