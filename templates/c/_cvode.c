{{define "_cvode_fs"}}{{$basicFB := .BasicFB}}{{$block := .}}

{{range $algIndex, $alg := $basicFB.Algorithms}}
{{if algorithmNeedsCvode $alg}}
void {{$block.Name}}_{{$alg.Name}}_f(realtype t, N_Vector x, N_Vector x_dot, void *f_data) {
	NV_Ith_S(x_dot, 0) = 1; //TODO ????
}
{{else if algorithmNeedsCvodeInit $alg}}
void {{$block.Name}}_{{$alg.Name}}_cvode_init({{$block.Name}}_t *me) {
	//AUTOGENERATED CODE: this algorithm specifies the initialization of a CVODE SUNDIALS solver and was parsed from the algorithm's text field
	int flag;
	{{$odeInit := parseOdeInitAlgo .Other.Text}}
	//specify tolerances
	realtype reltol = 1e-6;
    realtype abstol = 1e-8;

	//free solver if it is initialized
	if(me->cvode_mem != NULL) {
		CVodeFree(&me->cvode_mem);
		N_VDestroy_Serial(me->ode_solution);  /* Free y vector */
	}

	//create solver
	me->ode_solution = N_VNewSerial({{len $odeInit.GetInitialValues}}); //length of initial values
	me->cvode_mem = CVodeCreate(CV_ADAMS, CV_FUNCTIONAL);
	if (me->cvode_mem == 0) {
		fprintf(stderr, "Error in CVodeMalloc: could not allocate\n");
		while(1);
	}

	//specify initial values
	{{range $initVarIndex, $initVar := $odeInit.GetInitialValues}}
	NV_Ith_S(me->ode_solution, {{$initVarIndex}}) = {{$initVar.VarValue}};
	{{end}}
		
	me->T0 = 0; //???? should this always be 0 ????

	//initialize solver with pointer to values
	flag = CVodeInit(me->cvode_mem, {{$block.Name}}_{{$odeInit.OdeFName}}_f, me->T0, me->ode_solution);
    if (flag < 0) {
		fprintf(stderr, "Error in CVodeMalloc: %d\n", flag);
		while(1);
    }

	//set solver tolerances
	flag = CVodeSStolerances(me->cvode_mem, reltol, abstol);
	if (flag < 0) {
		fprintf(stderr, "Error in CVodeSStolerances: %d\n", flag);
		while(1);
	}
}
{{end}}

{{end}}
{{end}}