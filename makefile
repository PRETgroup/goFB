default: goFB goTFB

goFB: goFB/*
	go build -o ./bin/goFB -i ./goFB

goTFB: goTFB/*
	go build -o ./bin/goTFB -i ./goTFB

clean:
	rm ./bin/goFB
	rm ./bin/goTFB