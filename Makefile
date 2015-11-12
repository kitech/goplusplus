all:
	go install gopp
	go build -v

testv:
	go test -v ./...

test:
	go test ./...
