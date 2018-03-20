
--This is an autogenerated file
--Do not modify it by hand
--Generated at 2017-12-08T14:25:09+13:00
library ieee;
use ieee.std_logic_1164.all;
use ieee.numeric_std.all;
use work.enforcement_types_AlphabetEnforcer.all;

entity enforcer_AlphabetEnforcer_P1 is
    port
    (
        clk         : in std_logic;
        reset       : in std_logic;
        t           : in unsigned(63 downto 0); --current time in nanoseconds
        e           : out std_logic;            --if enforcement occured

        --the input signals
        
        --the enforce signals
        q           : in enforced_signals_AlphabetEnforcer;
        q_prime     : out enforced_signals_AlphabetEnforcer
    );
end entity;

architecture behaviour of enforcer_AlphabetEnforcer_P1 is
    signal trigger_tA : std_logic := '0';
    signal trigger_tA_time : unsigned(63 downto 0) := (others => '0');
    signal trigger_tB : std_logic := '0';
    signal trigger_tB_time : unsigned(63 downto 0) := (others => '0');
    
begin
    
    --trigger process
    process(reset, clk, q, t)
    variable q_enf: enforced_signals_AlphabetEnforcer;
    
    begin
        if(rising_edge(clk)) then
            --default values
            
            q_enf := q;
            e <= '0';

            --policies begin
            
            if((trigger_tA = '1') and (trigger_tB = '1') and not((q_enf.C = '1')) and (t > (to_unsigned(60000000, 64) + trigger_tA_time)) ) then
                e <= '1';
                --recover
                q_enf.C := '1';
                
            end if;
            
            if((trigger_tB = '1') and not((q_enf.D = '1')) and (t > (to_unsigned(10000000, 64) + trigger_tB_time)) ) then
                e <= '1';
                --recover
                q_enf.D := '1';
                
            end if;
            

            --Triggers begin (triggers are after policies because a policy might edit a value that a trigger depends on)
            
            if(trigger_tA = '0' and ((q_enf.A = '1'))) then
                trigger_tA <= '1';
                trigger_tA_time <= t;
            end if;
            
            if(trigger_tB = '0' and (((q_enf.B = '1')) and (t > (to_unsigned(30000000, 64) + trigger_tA_time)))) then
                trigger_tB <= '1';
                trigger_tB_time <= t;
            end if;
            
            
            q_prime <= q_enf;
        end if;
    end process;


end architecture;