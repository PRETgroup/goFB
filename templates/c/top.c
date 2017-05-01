{{define "top"}}// This file has been automatically generated by goFB
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
{{$block := index .Blocks .BlockIndex}}{{$blocks := .Blocks}}
//This is the main file for the iec61499 network with {{$block.Name}} as the top level block

#include "FB_{{$block.Name}}.h"

//put a copy of the top level block into global memory
{{$block.Name}}_t my{{$block.Name}};

int main() {
	printf("\n\n\n");

	if({{$block.Name}}_preinit(&my{{$block.Name}}) != 0) {
		printf("Failed to preinitialize.");
		return 1;
	}
	if({{$block.Name}}_init(&my{{$block.Name}}) != 0) {
		printf("Failed to initialize.");
		return 1;
	}

	printf("Top: %20s   Size: %lu\n", "{{$block.Name}}", sizeof(my{{$block.Name}}));

	int tickNum = 0;
	do {
		printf("%i,",tickNum);
		{{$block.Name}}_syncOutputEvents(&my{{$block.Name}});
		{{$block.Name}}_syncInputEvents(&my{{$block.Name}});

		{{$block.Name}}_syncOutputData(&my{{$block.Name}});
		{{$block.Name}}_syncInputData(&my{{$block.Name}});
		
		{{$block.Name}}_run(&my{{$block.Name}});
		printf("\n");
	} while(tickNum++ < 200);

	return 0;
}

{{end}}