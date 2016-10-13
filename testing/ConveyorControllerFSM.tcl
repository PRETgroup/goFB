vsim work.conveyorcontroller

add wave -position insertpoint  \
sim:/conveyorcontroller/clk \
sim:/conveyorcontroller/reset \
sim:/conveyorcontroller/enable \
sim:/conveyorcontroller/InjectDone \
sim:/conveyorcontroller/EmergencyStopChanged \
sim:/conveyorcontroller/LasersChanged \
sim:/conveyorcontroller/ConveyorChanged \
sim:/conveyorcontroller/ConveyorStoppedForInject \
sim:/conveyorcontroller/EmergencyStop \
sim:/conveyorcontroller/InjectSiteLaser \
sim:/conveyorcontroller/ConveyorSpeed \
sim:/conveyorcontroller/done \
sim:/conveyorcontroller/state \
sim:/conveyorcontroller/ConveyorStart_alg_en \
sim:/conveyorcontroller/ConveyorStart_alg_done \
sim:/conveyorcontroller/ConveyorStop_alg_en \
sim:/conveyorcontroller/ConveyorStop_alg_done \
sim:/conveyorcontroller/ConveyorRunning_alg_en \
sim:/conveyorcontroller/ConveyorRunning_alg_done \
sim:/conveyorcontroller/ConveyorEStop_alg_en \
sim:/conveyorcontroller/ConveyorEStop_alg_done \
sim:/conveyorcontroller/AlgorithmsStart \
sim:/conveyorcontroller/AlgorithmsDone \
sim:/conveyorcontroller/Variable1

force -freeze sim:/conveyorcontroller/clk 1 0, 0 {50 ps} -r 100
force -freeze sim:/conveyorcontroller/reset 0 0
force -freeze sim:/conveyorcontroller/enable 1 0
force -freeze sim:/conveyorcontroller/InjectDone 0 0
force -freeze sim:/conveyorcontroller/EmergencyStopChanged 0 0
force -freeze sim:/conveyorcontroller/LasersChanged 0 0
force -freeze sim:/conveyorcontroller/EmergencyStop 1 0
force -freeze sim:/conveyorcontroller/InjectSiteLaser 0 0

run
run

force -freeze sim:/conveyorcontroller/EmergencyStopChanged 1 0
force -freeze sim:/conveyorcontroller/EmergencyStop 0 0

run
run

force -freeze sim:/conveyorcontroller/EmergencyStop 1 0
force -freeze sim:/conveyorcontroller/EmergencyStopChanged 0 0

run
run

force -freeze sim:/conveyorcontroller/EmergencyStopChanged 1 0

run
run
