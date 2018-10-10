module top_with_timer_4

(
		input wire clk,
		
		//input events
		input wire VPulse_eI_1,
		input wire VPulse_eI_2,
		input wire VPulse_eI_3,
		input wire VPulse_eI_4,
		
				
		//output events
		output wire VPace_eO_1,
		output wire VPace_eO_2,
		output wire VPace_eO_3,
		output wire VPace_eO_4,
		
		output wire VRefractory_eO_1,
		output wire VRefractory_eO_2,
		output wire VRefractory_eO_3,
		output wire VRefractory_eO_4,
		
	
		input reset
);

top_with_timer t1(
		.clk(clk),
		
		//input events
		.VPulse_eI(VPulse_eI_1),
				
		//output events
		.VPace_eO(VPace_eO_1),
		.VRefractory_eO(VRefractory_eO_1),
	
		.reset(reset)
);

top_with_timer t2(
		.clk(clk),
		
		//input events
		.VPulse_eI(VPulse_eI_2),
				
		//output events
		.VPace_eO(VPace_eO_2),
		.VRefractory_eO(VRefractory_eO_2),
	
		.reset(reset)
);

top_with_timer t3(
		.clk(clk),
		
		//input events
		.VPulse_eI(VPulse_eI_3),
				
		//output events
		.VPace_eO(VPace_eO_3),
		.VRefractory_eO(VRefractory_eO_3),
	
		.reset(reset)
);

top_with_timer t4(
		.clk(clk),
		
		//input events
		.VPulse_eI(VPulse_eI_4),
				
		//output events
		.VPace_eO(VPace_eO_4),
		.VRefractory_eO(VRefractory_eO_4),
	
		.reset(reset)
);


endmodule