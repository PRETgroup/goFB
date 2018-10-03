{{define "_entityFB"}}
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}{{$specialIO := getSpecialIO $block .Blocks}}
entity {{$block.Name}} is

	port(
		--for clock and reset signal
		clk		: in	std_logic;
		reset	: in	std_logic;
		enable	: in	std_logic;
		sync	: in	std_logic;
		{{if $block.EventInputs}}
		--input events
		{{range $index, $event := $block.EventInputs}}{{$event.Name}}_eI : in std_logic := '0';
		{{end}}{{end}}
		{{if $block.EventOutputs}}
		--output events
		{{range $index, $event := $block.EventOutputs}}{{$event.Name}}_eO : out std_logic;
		{{end}}{{end}}
		{{if $block.InputVars}}
		--input variables
		{{range $index, $var := $block.InputVars}}{{$var.Name}}_I : in {{getVhdlType $var.Type}} := {{if eq (getVhdlType $var.Type) "std_logic"}}'0'{{else}}(others => '0'){{end}}; --type was {{$var.Type}}
		{{end}}{{end}}
		{{if $block.OutputVars}}
		--output variables
		{{range $index, $var := $block.OutputVars}}{{$var.Name}}_O : out {{getVhdlType $var.Type}}; --type was {{$var.Type}}
		{{end}}{{end}}
		{{if $block.BasicFB}}{{if $specialIO.InternalVars}}
		--special emitted internal vars for I/O
		{{range $index, $var := $specialIO.InternalVars}}{{$var.Name}} : {{if variableIsTOPIO_IN $var}}in{{else}}out{{end}} {{getVhdlType $var.Type}}; --type was {{$var.Type}}
		{{end}}{{end}}{{else if $block.CompositeFB}}{{if $specialIO.InternalVars}}
		--special emitted internal variables for child I/O
		{{range $index, $var := $specialIO.InternalVars}}{{$var.Name}} : {{if variableIsTOPIO_IN $var}}in{{else}}out{{end}} {{getVhdlType $var.Type}}; --type was {{$var.Type}} 
		{{end}}{{end}}{{end}}
		--for done signal
		done : out std_logic
	);

end entity;

{{end}}