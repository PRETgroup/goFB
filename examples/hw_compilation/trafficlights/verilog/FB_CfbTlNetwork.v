// This file has been automatically generated by goFB and should not be edited by hand
// Compiler written by Hammond Pearce and available at github.com/kiwih/goFB
// Verilog support is EXPERIMENTAL ONLY

// This file represents the Composite Function Block for CfbTlNetwork

module FB_CfbTlNetwork 

(
		input wire clk,
		
		//input events
		input wire Tick_eI,
		input wire Start_eI,
		input wire SpecialInstr_eI,
		input wire N_S_PedWaiting_eI,
		input wire E_W_PedWaiting_eI,
		
		//output events
		output wire N_S_PedLightsChange_eO,
		output wire N_S_TrafLightsChange_eO,
		output wire E_W_PedLightsChange_eO,
		output wire E_W_TrafLightsChange_eO,
		
		//input variables
		input wire  N_S_HoldGreen_I,
		input wire  E_W_HoldGreen_I,
		
		//output variables
		output wire  N_S_PedRed_O ,
		output wire  N_S_PedFlashRed_O ,
		output wire  N_S_PedGreen_O ,
		output wire  N_S_TrafRed_O ,
		output wire  N_S_TrafYellow_O ,
		output wire  N_S_TrafGreen_O ,
		output wire  E_W_PedRed_O ,
		output wire  E_W_PedFlashRed_O ,
		output wire  E_W_PedGreen_O ,
		output wire  E_W_TrafRed_O ,
		output wire  E_W_TrafYellow_O ,
		output wire  E_W_TrafGreen_O ,
		

		input reset
);


//Wires needed for event connections 
wire Tick_conn;
wire SpecialInstr_conn;
wire Start_conn;
wire mt_N_S_Start_conn;
wire N_S_DoneSeq_conn;
wire mt_E_W_GoSeq_conn;
wire E_W_DoneSeq_conn;
wire N_S_PedWaiting_conn;
wire E_W_PedWaiting_conn;
wire N_S_PedLightsChange_conn;
wire N_S_TrafLightsChange_conn;
wire E_W_PedLightsChange_conn;
wire E_W_TrafLightsChange_conn;


//Wires needed for data connections 
wire  N_S_HoldGreen_conn;
wire  E_W_HoldGreen_conn;
wire  N_S_PedRed_conn;
wire  N_S_PedFlashRed_conn;
wire  N_S_PedGreen_conn;
wire  N_S_TrafRed_conn;
wire  N_S_TrafYellow_conn;
wire  N_S_TrafGreen_conn;
wire  E_W_PedRed_conn;
wire  E_W_PedFlashRed_conn;
wire  E_W_PedGreen_conn;
wire  E_W_TrafRed_conn;
wire  E_W_TrafYellow_conn;
wire  E_W_TrafGreen_conn;


//top level I/O to signals
//input events
assign Tick_conn = Tick_eI;
assign Tick_conn = Tick_eI;
assign Start_conn = Start_eI;
assign SpecialInstr_conn = SpecialInstr_eI;
assign SpecialInstr_conn = SpecialInstr_eI;
assign N_S_PedWaiting_conn = N_S_PedWaiting_eI;
assign E_W_PedWaiting_conn = E_W_PedWaiting_eI;

//output events
assign N_S_PedLightsChange_eO = N_S_PedLightsChange_conn;
assign N_S_TrafLightsChange_eO = N_S_TrafLightsChange_conn;
assign E_W_PedLightsChange_eO = E_W_PedLightsChange_conn;
assign E_W_TrafLightsChange_eO = E_W_TrafLightsChange_conn;

//input variables
assign N_S_HoldGreen_conn = N_S_HoldGreen_I;
assign E_W_HoldGreen_conn = E_W_HoldGreen_I;

//output events
assign N_S_PedRed_O = N_S_PedRed_conn;
assign N_S_PedFlashRed_O = N_S_PedFlashRed_conn;
assign N_S_PedGreen_O = N_S_PedGreen_conn;
assign N_S_TrafRed_O = N_S_TrafRed_conn;
assign N_S_TrafYellow_O = N_S_TrafYellow_conn;
assign N_S_TrafGreen_O = N_S_TrafGreen_conn;
assign E_W_PedRed_O = E_W_PedRed_conn;
assign E_W_PedFlashRed_O = E_W_PedFlashRed_conn;
assign E_W_PedGreen_O = E_W_PedGreen_conn;
assign E_W_TrafRed_O = E_W_TrafRed_conn;
assign E_W_TrafYellow_O = E_W_TrafYellow_conn;
assign E_W_TrafGreen_O = E_W_TrafGreen_conn;



// child I/O to signals

FB_CfbOneLink N_S (
	.clk(clk),

	//event outputs 
	.DoneSeq_eO(N_S_DoneSeq_conn),
	.PedLightsChange_eO(N_S_PedLightsChange_conn),
	.TrafLightsChange_eO(N_S_TrafLightsChange_conn),
	
	//event inputs
	.Tick_eI(Tick_conn), 
	.SpecialInstr_eI(SpecialInstr_conn), 
	.GoSeq_eI(mt_N_S_Start_conn), 
	.PedWaiting_eI(N_S_PedWaiting_conn), 
	
	//data outputs
	.PedRed_O(N_S_PedRed_conn), 
	.PedFlashRed_O(N_S_PedFlashRed_conn), 
	.PedGreen_O(N_S_PedGreen_conn), 
	.TrafRed_O(N_S_TrafRed_conn), 
	.TrafYellow_O(N_S_TrafYellow_conn), 
	.TrafGreen_O(N_S_TrafGreen_conn), 
	
	//data inputs
	.HoldGreen_I(N_S_HoldGreen_conn),
	
	
	.reset(reset)
);

FB_CfbOneLink E_W (
	.clk(clk),

	//event outputs 
	.DoneSeq_eO(E_W_DoneSeq_conn),
	.PedLightsChange_eO(E_W_PedLightsChange_conn),
	.TrafLightsChange_eO(E_W_TrafLightsChange_conn),
	
	//event inputs
	.Tick_eI(Tick_conn), 
	.SpecialInstr_eI(SpecialInstr_conn), 
	.GoSeq_eI(mt_E_W_GoSeq_conn), 
	.PedWaiting_eI(E_W_PedWaiting_conn), 
	
	//data outputs
	.PedRed_O(E_W_PedRed_conn), 
	.PedFlashRed_O(E_W_PedFlashRed_conn), 
	.PedGreen_O(E_W_PedGreen_conn), 
	.TrafRed_O(E_W_TrafRed_conn), 
	.TrafYellow_O(E_W_TrafYellow_conn), 
	.TrafGreen_O(E_W_TrafGreen_conn), 
	
	//data inputs
	.HoldGreen_I(E_W_HoldGreen_conn),
	
	
	.reset(reset)
);

FB_BfbIntersectionMutex mt (
	.clk(clk),

	//event outputs 
	.N_S_Start_eO(mt_N_S_Start_conn),
	.E_W_GoSeq_eO(mt_E_W_GoSeq_conn),
	
	//event inputs
	.Start_eI(Start_conn), 
	.N_S_Done_eI(N_S_DoneSeq_conn), 
	.E_W_Done_eI(E_W_DoneSeq_conn), 
	
	//data outputs
	
	//data inputs
	
	
	.reset(reset)
);



endmodule