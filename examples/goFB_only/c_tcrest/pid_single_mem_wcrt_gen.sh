#!/bin/bash

echo "Patmos Single core PID experiment"
echo -e "PPC(/4)\tWCRT"

for i in {1..9} {10..45..5} {50..90..10} {100..450..50} {500..1000..100}
do
	make wcet PROJECT=pid_single_mem FUNCTION=timed_task PROGS_PER_CORE=$i T_DELAY_CYCLES=21 -B >/dev/null 2>/dev/null
	WCRT=$(cat pid_single_mem_wcet_report.txt | grep '^  cycles:' | egrep -o '[0-9]+') 
	echo -e "$i\t$WCRT"
done