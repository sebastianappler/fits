build :
	go build .

docker:
	docker build . -t fits
