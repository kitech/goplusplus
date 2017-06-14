all:
	go generate
	go install gopp
	go build -v main.go

test:
	go test -v gopp

testallv:
	go test -v ./...

testall:
	go test ./...
