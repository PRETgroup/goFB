{{define "_fbinit"}}{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}
/* {{$block.Name}}_preinit() is required to be called to 
 * initialise an instance of {{$block.Name}}. 
 * It sets all I/O values to zero.
 */
int {{$block.Name}}_preinit({{$block.Name}}_t {{if .TcrestUsingSPM}}_SPM{{end}} *me) {
	//if there are input events, reset them
	{{if $block.EventInputs}}{{range $index, $event := $block.EventInputs.Events}}me->inputEvents.event.{{$event.Name}} = 0;
	{{end}}{{end}}
	//if there are output events, reset them
	{{if $block.EventOutputs}}{{range $index, $event := $block.EventOutputs.Events}}me->outputEvents.event.{{$event.Name}} = 0;
	{{end}}{{end}}
	//if there are input vars with default values, set them
	{{if $block.InputVars}}{{range $index, $var := $block.InputVars.Variables}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}
	//if there are output vars with default values, set them
	{{if $block.OutputVars}}{{range $index, $var := $block.OutputVars.Variables}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}
	//if there are internal vars with default values, set them (BFBs only)
	{{if $block.BasicFB}}{{if $block.BasicFB.InternalVars}}{{range $varIndex, $var := $block.BasicFB.InternalVars.Variables}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}{{end}}
	//if there are resource vars with default values, set them
	{{if $block.ResourceVars}}{{range $index, $var := $block.ResourceVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}{{end}}
	//if there are resources with set parameters, set them
	{{if $block.Resources}}{{range $index, $res := $block.Resources}}{{if $res.Parameter}}{{range $paramIndex, $param := $res.Parameter}}me->{{$res.Name}}.{{$param.Name}} = {{$param.Value}};
	{{end}}{{end}}{{end}}{{end}}
	//if there are fb children (CFBs/Devices/Resources only), call this same function on them
	{{if $block.CompositeFB}}{{range $currChildIndex, $child := $block.CompositeFB.FBs}}if({{$child.Type}}_preinit(&me->{{$child.Name}}) != 0) {
		return 1;
	}
	{{end}}{{end}}
	{{if $block.Resources}}{{range $index, $res := $block.Resources}}if({{$res.Type}}_preinit(&me->{{$res.Name}}) != 0) {
		return 1;
	}
	{{end}}{{end}}
	//if this is a BFB/odeFB, set start state so that the start state is properly executed and _trigger if necessary
	{{if $block.BasicFB}}me->_state = STATE_{{$block.Name}}_{{(index $block.BasicFB.States 0).Name}};
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
	{{if $block.CompositeFB}}{{range $currLinkIndex, $link := $block.CompositeFB.DataConnections}}me->{{$link.Destination}} = me->{{$link.Source}};
	{{end}}{{end}}

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