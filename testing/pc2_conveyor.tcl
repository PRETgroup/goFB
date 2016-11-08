add wave -position insertpoint  \
sim:/iec61499_network_top/clk \
sim:/iec61499_network_top/reset \
sim:/iec61499_network_top/rx_conveyor_moving \
sim:/iec61499_network_top/rx_conveyor_jammed \
sim:/iec61499_network_top/tx_conveyor_run \
sim:/iec61499_network_top/rx_global_run \
sim:/iec61499_network_top/rx_global_run_infinite \
sim:/iec61499_network_top/rx_global_fault \
sim:/iec61499_network_top/tx_box_dropper_run \
sim:/iec61499_network_top/debug_enable \
sim:/iec61499_network_top/debug_sync \
sim:/iec61499_network_top/debug_done \
sim:/iec61499_network_top/state \
sim:/iec61499_network_top/enable \
sim:/iec61499_network_top/sync \
sim:/iec61499_network_top/done
force -freeze sim:/iec61499_network_top/clk 1 0, 0 {50 ps} -r 100
force -freeze sim:/iec61499_network_top/reset 0 0
force -freeze sim:/iec61499_network_top/rx_conveyor_moving 0 0
force -freeze sim:/iec61499_network_top/rx_conveyor_jammed 0 0
force -freeze sim:/iec61499_network_top/rx_global_run 0 0
force -freeze sim:/iec61499_network_top/rx_global_run_infinite 0 0
force -freeze sim:/iec61499_network_top/rx_global_fault 0 0
sim:/iec61499_network_top/top_block/globals/clk \
sim:/iec61499_network_top/top_block/globals/reset \
sim:/iec61499_network_top/top_block/globals/enable \
sim:/iec61499_network_top/top_block/globals/sync \
sim:/iec61499_network_top/top_block/globals/global_run_changed_eO \
sim:/iec61499_network_top/top_block/globals/global_fault_changed_eO \
sim:/iec61499_network_top/top_block/globals/global_run_O \
sim:/iec61499_network_top/top_block/globals/global_run_infinite_O \
sim:/iec61499_network_top/top_block/globals/global_fault_O \
sim:/iec61499_network_top/top_block/globals/rx_global_run \
sim:/iec61499_network_top/top_block/globals/rx_global_run_infinite \
sim:/iec61499_network_top/top_block/globals/rx_global_fault \
sim:/iec61499_network_top/top_block/globals/done \
sim:/iec61499_network_top/top_block/globals/state \
sim:/iec61499_network_top/top_block/globals/global_run \
sim:/iec61499_network_top/top_block/globals/global_run_infinite \
sim:/iec61499_network_top/top_block/globals/global_fault \
sim:/iec61499_network_top/top_block/globals/global_run_changed_eO_ecc_out \
sim:/iec61499_network_top/top_block/globals/global_run_changed_eO_alg_out \
sim:/iec61499_network_top/top_block/globals/global_fault_changed_eO_ecc_out \
sim:/iec61499_network_top/top_block/globals/global_fault_changed_eO_alg_out \
sim:/iec61499_network_top/top_block/globals/globals_alg_alg_en \
sim:/iec61499_network_top/top_block/globals/globals_alg_alg_done \
sim:/iec61499_network_top/top_block/globals/AlgorithmsStart \
sim:/iec61499_network_top/top_block/globals/AlgorithmsDone \
sim:/iec61499_network_top/top_block/globals/rx_global_run_prev \
sim:/iec61499_network_top/top_block/globals/rx_global_run_infinite_prev \
sim:/iec61499_network_top/top_block/globals/rx_global_fault_prev