BENCH=VecAddScaled

lint:
	@golangci-lint run

bench:
	@go test -count=8 -bench=Benchmark$(BENCH) -benchtime=333ms -run=^$

benchmem:
	@go test -bench=$(BENCH) -benchmem -benchtime=33ms -run=^$

.PHONY: lint bench benchmem
