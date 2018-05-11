{{define "_policyIn"}}{{$block := .}}
	//input policies
	{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
	{{if not $pfbEnf}}//{{$pol.Name}} is broken!
	{{else}}{{/* this is where the policy comes in */}}//{{$pfbEnf}}
		{{range $vari, $var := $pfbEnf.InputPolicy.InternalVars}}static {{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}} {0};{{end}}
		{{end}}
		{{range $tri, $tr := $pfbEnf.InputPolicy.GetViolationTransitions}}
		if({{$tr.Condition}}) {

		}
		{{end}}
	{{end}}
	{{end}}
{{end}}

{{define "_policyOut"}}
//output policies
{{end}}