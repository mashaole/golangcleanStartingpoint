setup:
	go mod vendor
	go mod tidy
	go install
	
run:
	nodemon --exec "go run" *.go

start:
	go run *.go
build:
	go build