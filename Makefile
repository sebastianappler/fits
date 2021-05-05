build :
	go build ./cmd/fits

build-debug:
	go build -gcflags=all="-N -l" ./cmd/fits

docker:
	docker build . -t fits
