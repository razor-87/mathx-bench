ACTION=./cmd/action
BENCH=VecAddition

lint:
	@golangci-lint run

run:
	@go run $(ACTION)

run-measure:
	@go run $(ACTION) -measure=true

bench:
	@go test -count=8 -bench=Benchmark$(BENCH) -benchtime=333ms -run=^$

benchmem:
	@go test -bench=$(BENCH) -benchmem -benchtime=33ms -run=^$

.PHONY: lint run run-measure bench benchmem
