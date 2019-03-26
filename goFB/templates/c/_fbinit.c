{{define "_fbinit"}}{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}{{$tcrestUsingSPM := .TcrestUsingSPM}}{{$tcrestSmartSPM := .TcrestSmartSPM}}{{$runOnECC := .RunOnECC}}{{$eventQueue := .EventQueue}}
/* {{$block.Name}}_preinit() is required to be called to 
 * initialise an instance of {{$block.Name}}. 
 * It sets all I/O values to zero.
 */
int {{$block.Name}}_preinit({{$block.Name}}_t {{if .TcrestUsingSPM}}_SPM{{end}} *me{{if $eventQueue}}, short myInstanceID{{end}}) {
	{{if $eventQueue}}//as we're using event queue, each FB has a unique instance ID to be associated with emitted events
	me->myInstanceID = myInstanceID;
	printf("I just set a {{$block.Name}} to have instance ID %i\n", myInstanceID);
	short instanceOffset = 1;//this is what is used to create the instance numbers of any child blocks
	{{end}}

	{{range $index, $event := $block.EventInputs}}//reset the input events
	me->inputEvents.event.{{$event.Name}} = 0;
	{{end}}
	{{range $index, $event := $block.EventOutputs}}//reset the output events
	me->outputEvents.event.{{$event.Name}} = 0;
	{{end}}
	//set any input vars with default values
	{{range $index, $var := $block.InputVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}
	//set any output vars with default values
	{{range $index, $var := $block.OutputVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}
	{{if $block.BasicFB}}//set any internal vars with default values
	{{if $block.BasicFB.InternalVars}}{{range $varIndex, $var := $block.BasicFB.InternalVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}{{end}}
	{{if $block.ResourceVars}}//set any resource vars with default values
	{{range $index, $var := $block.ResourceVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}
	{{if $block.Resources}}//set any resource params
	{{range $index, $res := $block.Resources}}{{if $res.Parameter}}{{range $paramIndex, $param := $res.Parameter}}me->{{$res.Name}}.{{$param.Name}} = {{$param.Value}};
	{{end}}{{end}}{{end}}{{end}}
	{{if $block.CompositeFB}}//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	{{range $currChildIndex, $child := $block.CompositeFB.FBs}}{{$childType := findBlockDefinitionForType $blocks $child.Type}}if({{$child.Type}}_preinit(&me->{{$child.Name}}{{if $eventQueue}}, myInstanceID + instanceOffset{{end}}) != 0) {
		return 1;
	}
	{{if $eventQueue}}instanceOffset += ({{$childType.NumChildren}} + 1);{{end}}
	{{end}}{{end}}
	{{if $block.Resources}}{{range $index, $res := $block.Resources}}{{$childType := findBlockDefinitionForType $blocks $res.Type}}if({{$res.Type}}_preinit(&me->{{$res.Name}}{{if $eventQueue}}, myInstanceID + instanceOffset{{end}}) != 0) {
		return 1;
	}
	{{if $eventQueue}}instanceOffset += ({{$childType.NumChildren}} + 1);{{end}}
	{{end}}{{end}}

	{{if $block.ServiceFB}}{{if $block.ServiceFB.Autogenerate}}//Code provided in SIFB
	{{$block.ServiceFB.Autogenerate.PreInitText}}{{end}}{{end}}
	
	{{if $block.BasicFB}}//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	{{if len $block.BasicFB.States}}me->_state = STATE_{{$block.Name}}_{{(index $block.BasicFB.States 0).Name}};{{else}}me->_state = STATE_{{$block.Name}}_unknown;{{end}}
	me->_trigger = true;
	{{if and .CvodeEnabled (blockNeedsCvode $block)}}
	me->cvode_mem = NULL;
	me->Tcurr = 0;
	me->Tnext = 0;
	me->T0 = 0;
	me->solveInProgress = 0;
	
	#ifdef PRINT_VALS
	{{if $block.OutputVars}}{{range $ind, $outputVar := $block.OutputVars.Variables}}	printf("{{$block.Name}}-{{$outputVar.Name}},");
	{{end}}{{end}}
	#endif{{else}}
	{{end}}
	{{end}}

	{{if $block.Policies}}//this block has policies
	{{range $polI, $pol := $block.Policies}}
	me->_policy_{{$pol.Name}}_state = {{if $pol.States}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{(index $pol.States 0).Name}}{{else}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_unknown{{end}};
	{{$pfbEnf := getPolicyEnfInfo $block $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//input policy internal vars
	{{/*{{range $vari, $var := $pfbEnf.InputPolicy.InternalVars}}
	{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
	{{end}}
	{{end}}*/}}//output policy internal vars
	{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}
	{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
	{{end}}
	{{end}}{{end}}
	{{end}}
	{{end}}

	return 0;
}

/* {{$block.Name}}_init() is required to be called to 
 * set up an instance of {{$block.Name}}. 
 * It passes around configuration data.
 */
int {{$block.Name}}_init({{$block.Name}}_t {{if .TcrestUsingSPM}}_SPM{{end}} *me) {
	//pass in any parameters on this level
	{{if $block.CompositeFB}}{{range $currChildIndex, $child := $block.CompositeFB.FBs}}{{range $currParamIndex, $param := $child.Parameter}}me->{{$child.Name}}.{{$param.Name}} = {{$param.Value}};
	{{end}}{{end}}{{end}}
	{{if $block.Resources}}{{range $currChildIndex, $child := $block.Resources}}{{range $currParamIndex, $param := $child.Parameter}}me->{{$child.Name}}.{{$param.Name}} = {{$param.Value}};
	{{end}}{{end}}{{end}}
	

	//perform a data copy to all children (if any present) (can move config data around, doesn't do anything otherwise)
	{{/*{{if $block.CompositeFB}}{{range $currLinkIndex, $link := $block.CompositeFB.DataConnections}}me->{{$link.Destination}} = me->{{$link.Source}};
	{{end}}{{end}}*/}}
	
	{{if $block.CompositeFB -}}
		{{$compositeFB := $block.CompositeFB -}}
		{{range $currChildIndex, $child := $compositeFB.FBs -}}
			{{$childType := findBlockDefinitionForType $blocks $child.Type}}//sync config for {{$child.Name}} (of Type {{$childType.Name}}) 
			{{if $childType.InputVars -}}
				{{range $inputVarIndex, $inputVar := $childType.InputVars -}}
				{{$source := findSourceDataName $compositeFB.DataConnections $child.Name $inputVar.Name -}}
					{{if $source -}}
						{{if $inputVar.GetArraySize -}}
							{{range $index, $count := count $inputVar.GetArraySize -}}
	me->{{$child.Name}}.{{$inputVar.Name}}[{{$count}}] = me->{{$source}}[{{$count}}];{{end}}
						{{else -}}
							{{if isNumeric $source -}}
	me->{{$child.Name}}.{{$inputVar.Name}} = {{$source}};
							{{else -}}
	me->{{$child.Name}}.{{$inputVar.Name}} = me->{{$source}};
							{{- end -}}
						{{- end -}}
					{{- end -}}
				{{- end -}}
			{{- end -}}
		{{- end -}}
	{{- end -}}

	{{if $block.ServiceFB}}{{if $block.ServiceFB.Autogenerate}}//Code provided in SIFB
	{{$block.ServiceFB.Autogenerate.InitText}}{{end}}{{end}}

	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	{{if $block.CompositeFB}}{{range $currChildIndex, $child := $block.CompositeFB.FBs}}if({{$child.Type}}_init(&me->{{$child.Name}}) != 0) {
		return 1;
	}
	{{end}}{{end}}
	{{if $block.Resources}}{{range $index, $res := $block.Resources}}if({{$res.Type}}_init(&me->{{$res.Name}}) != 0) {
		return 1;
	}
	{{end}}{{end}}

	return 0;
}

{{end}}