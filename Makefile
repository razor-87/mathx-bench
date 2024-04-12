lint:
	@golangci-lint run

bench:
	@go test -count=8 -bench=. -benchtime=333ms -run=^$

benchmem:
	@go test -bench=. -benchmem -benchtime=33ms -run=^$

.PHONY: lint bench benchmem
