{{define "basicFB"}}// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
// Verilog support is EXPERIMENTAL ONLY
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}{{$basicFB := $block.BasicFB}}
// This file represents the Basic Function Block for {{$block.Name}}

module FB_{{$block.Name}} {{template "_moduleDeclr" .}}

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
{{range $index, $var := $block.OutputVars}}reg {{getVerilogSize $var.Type}} {{$var.Name}};
{{end}}{{end}}
////END internal copies of I/O

////BEGIN internal vars
{{if $basicFB.InternalVars}}{{range $varIndex, $var := $basicFB.InternalVars}}
reg  {{getVerilogSize $var.Type}} {{$var.Name}} {{if $var.InitialValue}} = {{$var.InitialValue}}{{end}}; {{end}}{{end}}
////END internal vars

always@(posedge clk) begin
	//BEGIN update internal inputs

	//END update internal inputs



	//BEGIN update 

end
endmodule{{end}}