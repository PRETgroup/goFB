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
		input wire unsigned [31:0] i_I,
		input wire unsigned [31:0] iSet_I,
		
		

		input reset
);


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
reg unsigned [31:0] i ;
reg unsigned [31:0] iSet ;


////END internal copies of I/O

////BEGIN internal vars

reg  unsigned [63:0] v  = 0; 
reg  unsigned [63:0] thresh  = 0; 
reg  unsigned [31:0] K  = 10000; 
reg  unsigned [31:0] B  = 135; 
////END internal vars

//BEGIN STATE variables
reg [1:0] state = `STATE_s_start;
reg entered = 1'b0;
//END STATE variables

//BEGIN algorithm triggers
reg s_wait_alg0_alg_en = 1'b0; 
reg s_count_alg0_alg_en = 1'b0; 
reg updateThresh_alg_en = 1'b0; 

//END algorithm triggers


always@(posedge clk) begin

	if(reset) begin
		//reset state 
		state = `STATE_s_start;

		//reset I/O registers
		unsafe = 1'b0;
		
		i = 0;
		iSet = 0;
		
		//reset internal vars
		v = 0;
		thresh = 0;
		K = 10000;
		B = 135;
	end else begin

		//BEGIN clear output events
		unsafe = 1'b0;
		
		//END clear output events

		//BEGIN update internal inputs on relevant events
		
		if(i_measured) begin 
			i = i_I;
			
		end
		
		if(iSet_change) begin 
			iSet = iSet_I;
			
		end
		
		//END update internal inputs

		//BEGIN ecc 
		entered = 1'b0;
		case(state) 
			default: begin
				if(1) begin
					state = `STATE_s_wait;
					entered = 1'b1;
				end
			end 
			`STATE_s_wait: begin
				if(i > iSet) begin
					state = `STATE_s_count;
					entered = 1'b1;
				end
			end 
			`STATE_s_count: begin
				if(i <= iSet) begin
					state = `STATE_s_wait;
					entered = 1'b1;
				end else if(v > thresh) begin
					state = `STATE_s_over;
					entered = 1'b1;
				end else if(tick) begin
					state = `STATE_s_count;
					entered = 1'b1;
				end
			end 
			`STATE_s_over: begin
				if(i <= iSet) begin
					state = `STATE_s_wait;
					entered = 1'b1;
				end else if(1) begin
					state = `STATE_s_over;
					entered = 1'b1;
				end
			end 
			
		endcase
		//END ecc

		//BEGIN triggers
		s_wait_alg0_alg_en = 1'b0; 
		s_count_alg0_alg_en = 1'b0; 
		updateThresh_alg_en = 1'b0; 
		
		if(entered) begin
			case(state)
				default: begin
					
				end 
				`STATE_s_wait: begin
					s_wait_alg0_alg_en = 1'b1;
					
				end 
				`STATE_s_count: begin
					updateThresh_alg_en = 1'b1;
					s_count_alg0_alg_en = 1'b1;
					
				end 
				`STATE_s_over: begin
					unsafe = 1'b1;
					
				end 
				
			endcase
		end
		//END triggers
		
		//BEGIN algorithms
		if(s_wait_alg0_alg_en) begin
			v = 0;

		end 
		if(s_count_alg0_alg_en) begin
			v = v + 1;

		end 
		if(updateThresh_alg_en) begin
			thresh = K * B / (i / iSet - 1);

		end 
		
		//END algorithms

		//BEGIN update external output variables on relevant events
		
		//END update external output variables 
	end
end
endmodule