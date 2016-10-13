{{define "_entityFB"}}

entity {{.Name}} is

	port(
		--for clock and reset signal
		clk		: in	std_logic;
		reset	: in	std_logic;
		enable	: in	std_logic;
		{{if .EventInputs}}
		--input events
		{{range $index, $event := .EventInputs.Events}}{{$event.Name}} : in std_logic;
		{{end}}{{end}}
		{{if .EventOutputs}}
		--output events
		{{range $index, $event := .EventOutputs.Events}}{{$event.Name}} : out std_logic;
		{{end}}{{end}}
		{{if .InputVars}}
		--input variables
		{{range $index, $var := .InputVars.Variables}}{{$var.Name}} : in {{getVhdlType $var.Type}}; --type was {{$var.Type}}
		{{end}}{{end}}
		{{if .OutputVars}}
		--output variables
		{{range $index, $var := .OutputVars.Variables}}{{$var.Name}} : out {{getVhdlType $var.Type}}; --type was {{$var.Type}}
		{{end}}{{end}}
		--for done signal
		done : out std_logic
	);

end entity;

{{end}}