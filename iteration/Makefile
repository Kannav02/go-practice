
.PHONY: fmt vet build


fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build 
test:
	go test
benchmark:
	go test -bench=.

.DEFAULT_GOAL = build