bench-matrix:
	mkdir -p results/matrix
	go test -race -v -bench="Matrix" ./... -benchtime=1x -benchmem -json | tee results/matrix/$(shell date +%Y-%m-%d-%T).json

bench-utils:
	go test -race -v -bench="Utils" ./... -benchtime=100000x -benchmem

test:
	go test -race -v  ./...
