all:
	6g fraction.go
	6l -o fraction fraction.6

clean:
	rm -f *.6 fraction
