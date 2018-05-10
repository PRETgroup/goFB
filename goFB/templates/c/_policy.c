{{define "_policyIn"}}{{$block := .}}
	//input policies
	{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
	{{if not $pfbEnf}}//{{$pol.Name}} is broken!
	{{else}}//{{$pfbEnf}}
	{{end}}
	{{end}}
{{end}}

{{define "_policyOut"}}
//output policies
{{end}}