-- This file has been automatically generated by go-iec61499-vhdl and should not be edited by hand
-- Converter written by Hammond Pearce and available at github.com/kiwih/go-iec61499-vhdl

-- This file represents the Basic Function Block for BoxDropper_SIFB

library ieee;
use ieee.std_logic_1164.all;
use ieee.numeric_std.all;



entity BoxDropper_SIFB is

	port(
		--for clock and reset signal
		clk		: in	std_logic;
		reset	: in	std_logic;
		enable	: in	std_logic;
		sync	: in	std_logic;
		
		--input events
		box_dropper_run_changed_eI : in std_logic := '0';
		
		
		
		--input variables
		box_dropper_run_I : in std_logic := '0'; --type was BOOL
		
		
		
		--special emitted internal vars for I/O
		tx_box_dropper_run : out std_logic; --type was BOOL
		
		--for done signal
		done : out std_logic
	);

end entity;


architecture rtl of BoxDropper_SIFB is
	-- Build an enumerated type for the state machine
	type state_type is (STATE_Start);

	-- Register to hold the current state
	signal state   : state_type := STATE_Start;

	-- signals to store variable sampled on enable 
	signal box_dropper_run : std_logic := '0'; --register for input
	
	

	

	-- signals for enabling algorithms	
	signal boxdropper_alg_alg_en : std_logic := '0'; 
	signal boxdropper_alg_alg_done : std_logic := '1';
	

	-- signal for algorithm completion
	signal AlgorithmsStart : std_logic := '0';
	signal AlgorithmsDone : std_logic;

	--internal variables 
begin
	-- Registers for data variables (only updated on relevant events)
	process (clk)
	begin
		if rising_edge(clk) then
			if sync = '1' then
				
				if box_dropper_run_changed_eI = '1' then
					box_dropper_run <= box_dropper_run_I;
				end if;
				
			end if;
		end if;
	end process;
	
			
	
	-- Logic to advance to the next state
	process (clk, reset)
	begin
		if reset = '1' then
			state <= STATE_Start;
			AlgorithmsStart <= '1';
		elsif (rising_edge(clk)) then
			if AlgorithmsStart = '1' then --algorithms should be triggered only once via this pulse signal
				AlgorithmsStart <= '0';
			elsif enable = '1' then 
				--default values
				state <= state;
				AlgorithmsStart <= '0';

				--next state logic
				case state is
					when STATE_Start =>
						if true then
							state <= STATE_Start;
							AlgorithmsStart <= '1';
						end if;
					
				end case;

			end if;
		end if;
	end process;

	-- Event outputs and internal algorithm triggers depend solely on the current state
	process (state)
	begin
		--default values
		
		--algorithms
		boxdropper_alg_alg_en <= '0'; 

		case state is
			when STATE_Start =>
				boxdropper_alg_alg_en <= '1';
				
			
		end case;
	end process;

	-- Algorithms process
	process(clk)
	begin
		if rising_edge(clk) then
			if AlgorithmsStart = '1' then			
				
				if boxdropper_alg_alg_en = '1' then -- Algorithm boxdropper_alg
					boxdropper_alg_alg_done <= '0';
					
				end if;
				
			end if;

			
			if boxdropper_alg_alg_done = '0' then -- Algorithm boxdropper_alg

--begin algorithm raw text
tx_box_dropper_run <= box_dropper_run;
boxdropper_alg_alg_done <= '1';
--end algorithm raw text

			end if;
			
		end if;
	end process;

	--Done signal
	AlgorithmsDone <= (not AlgorithmsStart) and (not enable) and boxdropper_alg_alg_done;
	Done <= AlgorithmsDone;

	

end rtl;
