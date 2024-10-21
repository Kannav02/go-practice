# this is used to avoid confusion if there is a file that exists in the directory as the same name of the commands
.PHONY : fmt vet build

# formats the code , is inbuilt in go
fmt:
	go fmt ./...
# vet is like linter but for go
vet: fmt
	go vet ./...
clean:
	go clean
	rm -f go-practice
# target: dependency if any
# build is used to make executable
build: vet
	go build

# default target that will run when calling `make`
.DEFAULT_GOAL = build