// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
// Verilog support is EXPERIMENTAL ONLY

// This file represents the Basic Function Block for BfbIDMTCurve

//defines for state names used internally
`define STATE_s_start 0
`define STATE_s_wait 1
`define STATE_s_count 2
`define STATE_s_over 3


module FB_BfbIDMTCurve 

(
		input wire clk,
		
		
		//input events
		input wire tick_eI,
		input wire i_measured_eI,
		input wire iSet_change_eI,
		
		
		//output events
		output wire unsafe_eO,
		
		
		//input variables
		input wire unsigned [63:0] i_I,
		input wire unsigned [63:0] iSet_I,
		
		

		input reset
);


////BEGIN algorithm functions

function s_wait_alg0

begin
v = 0;
endfunction
function s_count_alg0

begin
v = v + 1;
endfunction
function updateThresh

begin
thresh = K * B / (I_mA / Iset_mA - 1);
endfunction
////END algorithm functions

////BEGIN internal copies of I/O
//input events
wire tick;
assign tick = tick_eI;
wire i_measured;
assign i_measured = i_measured_eI;
wire iSet_change;
assign iSet_change = iSet_change_eI;

//output events
reg unsafe;
assign unsafe_eO = unsafe;

//input variables
reg unsigned [63:0] i ;
reg unsigned [63:0] iSet ;


////END internal copies of I/O

////BEGIN internal vars

reg  unsigned [63:0] v  = 0; 
reg  unsigned [63:0] thresh  = 0; 
reg  unsigned [63:0] K  = 10000; 
reg  unsigned [63:0] B  = 135; 
////END internal vars

//STATE variable
reg integer state = `STATE_s_start;

always@(posedge clk) begin
	//BEGIN update internal inputs on relevant events
	
	if(i_measured) begin 
		i = i_I;
		
	end
	
	if(iSet_change) begin 
		iSet = iSet_I;
		
	end
	
	//END update internal inputs

	//BEGIN ecc 
	


	//END ecc

	//BEGIN update external outputs on relevant events
	
	//END update external outputs

end
endmodule