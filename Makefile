build:
	export GO111MODULE=on
	go get ./...
	go build -o functions/profile ./...