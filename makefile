.PHONY: default

default: goFB goTFB

goFB: 
	cd goFB && go build -o ../bin/goFB

goTFB: 
	cd goTFB && go build -o ../bin/goTFB

clean:
	rm -f ./bin/goFB
	rm -f ./bin/goTFB