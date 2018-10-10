// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
// Verilog support is EXPERIMENTAL ONLY

// This file represents the Basic Function Block for CanisterCounter

//defines for state names used internally
`define STATE_Start 0


module FB_CanisterCounter 

(
		input wire clk,
		
		//input events
		input wire LasersChanged_eI,
		
		//output events
		output wire CanisterCountChanged_eO,
		
		//input variables
		input wire  DoorSiteLaser_I,
		input wire  RejectBinLaser_I,
		input wire  AcceptBinLaser_I,
		
		//output variables
		output reg [7:0] CanisterCount_O ,
		

		input reset
);


////BEGIN internal copies of I/O
//input events
wire LasersChanged;
assign LasersChanged = LasersChanged_eI;

//output events
reg CanisterCountChanged;
assign CanisterCountChanged_eO = CanisterCountChanged;

//input variables
reg  DoorSiteLaser ;
reg  RejectBinLaser ;
reg  AcceptBinLaser ;

//output variables
reg [7:0] CanisterCount ;

////END internal copies of I/O

////BEGIN internal vars

////END internal vars

//BEGIN STATE variables
reg  state = `STATE_Start;
reg entered = 1'b0;
//END STATE variables

//BEGIN algorithm triggers
reg ChangeCount_alg_en = 1'b0; 

//END algorithm triggers


always@(posedge clk) begin

	if(reset) begin
		//reset state 
		state = `STATE_Start;

		//reset I/O registers
		CanisterCountChanged = 1'b0;
		
		DoorSiteLaser = 0;
		RejectBinLaser = 0;
		AcceptBinLaser = 0;
		
		CanisterCount = 0;
		//reset internal vars
	end else begin

		//BEGIN clear output events
		CanisterCountChanged = 1'b0;
		
		//END clear output events

		//BEGIN update internal inputs on relevant events
		
		if(LasersChanged) begin 
			DoorSiteLaser = DoorSiteLaser_I;
			RejectBinLaser = RejectBinLaser_I;
			AcceptBinLaser = AcceptBinLaser_I;
			
		end
		
		//END update internal inputs

		//BEGIN ecc 
		entered = 1'b0;
		case(state) 
			`STATE_Start: begin
				if(LasersChanged) begin
					state = `STATE_Start;
					entered = 1'b1;
				end
			end 
			default: begin
				state = 0;
			end
		endcase
		//END ecc

		//BEGIN triggers
		ChangeCount_alg_en = 1'b0; 
		
		if(entered) begin
			case(state)
				`STATE_Start: begin
					ChangeCount_alg_en = 1'b1;
					CanisterCountChanged = 1'b1;
					
				end 
				default: begin

				end
			endcase
		end
		//END triggers
		
		//BEGIN algorithms
		if(ChangeCount_alg_en) begin
			if (DoorSiteLaser) begin
			CanisterCount = CanisterCount + 1;

		end
		if (RejectBinLaser) begin
			CanisterCount = CanisterCount - 1;

		end
		if (AcceptBinLaser) begin
			CanisterCount = CanisterCount - 1;

		end
		
		end 
		
		//END algorithms

		//BEGIN update external output variables on relevant events
		
		if(CanisterCountChanged) begin 
			CanisterCount_O = CanisterCount;
			
		end
		
		//END update external output variables 
	end
end
endmodule