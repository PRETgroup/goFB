#!/bin/bash

echo "T-CREST 4 Multicore SPM PID experiment"
echo -e "PPC\tWCRT"

for i in {1..8}
do
	make wcet PROJECT=pid_spm FUNCTION=timed_task PROGS_PER_CORE=$i -B >/dev/null 2>/dev/null
	WCRT=$(cat pid_spm_wcet_report.txt | grep '^  cycles:' | egrep -o '[0-9]+') 
	echo -e "$i\t$WCRT"
done