
run-interpreter: build-interpreter
	@./monkey

build-interpreter:
	@go build -o monkey ./cmd/interpreter
