{{define "_policyIn"}}{{$block := .}}
	//input policies
	{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
	{{if not $pfbEnf}}//{{$pol.Name}} is broken!
	{{else}}{{/* this is where the policy comes in */}}//INPUT POLICY {{$pol.Name}} BEGIN 
		switch(me->_policy_{{$pol.Name}}_state) {
			{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
				{{range $tri, $tr := $pfbEnf.InputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
				*/}}
				if({{$cond := getCECCTransitionCondition $block $tr.Condition}}{{$cond.IfCond}}) {
					//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
					//select a transition to solve the problem
					{{$solution := $pfbEnf.SolveViolationTransition $tr true}}
					{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
					{{if $solution.Expression}}{{$sol := getCECCTransitionCondition $block $solution.Expression}}{{$sol.IfCond}};{{end}}
				} {{end}}{{end}}
				
				break;

			{{end}}
		}
	{{end}}
	//INPUT POLICY {{$pol.Name}} END
	{{end}}
{{end}}

{{define "_policyOut"}}{{$block := .}}
	//output policies
	{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
	{{if not $pfbEnf}}//{{$pol.Name}} is broken!
	{{else}}{{/* this is where the policy comes in */}}//OUTPUT POLICY {{$pol.Name}} BEGIN 
		switch(me->_policy_{{$pol.Name}}_state) {
			{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
				{{range $tri, $tr := $pfbEnf.OutputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
				*/}}
				if({{$cond := getCECCTransitionCondition $block $tr.Condition}}{{$cond.IfCond}}) {
					//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
					//select a transition to solve the problem
					{{$solution := $pfbEnf.SolveViolationTransition $tr false}}
					{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
					{{if $solution.Expression}}{{$sol := getCECCTransitionCondition $block $solution.Expression}}{{$sol.IfCond}};{{end}}
				} {{end}}{{end}}

				break;

			{{end}}
		}

		//advance timers
		{{range $varI, $var := $pfbEnf.OutputPolicy.GetDTimers}}
		me->{{$var.Name}}++;{{end}}

		//select transition to advance state
		switch(me->_policy_{{$pol.Name}}_state) {
			{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
				{{range $tri, $tr := $pfbEnf.OutputPolicy.GetNonViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
				*/}}
				if({{$cond := getCECCTransitionCondition $block $tr.Condition}}{{$cond.IfCond}}) {
					//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
					me->_policy_{{$pol.Name}}_state = POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$tr.Destination}};
				} {{end}}{{end}}
				
				break;

			{{end}}
		}
	{{end}}
	//OUTPUT POLICY {{$pol.Name}} END
	{{end}}
{{end}}