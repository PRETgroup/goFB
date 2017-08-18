default: fbt/%.fbt c/%.c bin

fbt/%.fbt: tfb/%.tfb
	goTFB -i=tfb -o=fbt

c/%.c: fbt/%.fbt
	goFB -i=./fbt -o=./c -l=c -t=Top -eventMoC=true

bin: c/%.c
	gcc c/*.c -o railway_station_v2.out

clean:
	rm -f railway_station_v2.out
	rm -f fbt/*
	rm -f c/*