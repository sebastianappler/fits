build :
	go build .

build-debug:
	go build -gcflags=all="-N -l" .

docker:
	docker build . -t fits
