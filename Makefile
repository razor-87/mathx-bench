lint:
	@golangci-lint run

bench:
	@go test -count=8 -bench=. -benchtime=333ms -run=^$

benchmem:
	@go test -bench=. -benchmem -benchtime=33ms -run=^$

benchcpu:
	@go test -bench=. -cpuprofile cpu.prof -benchtime=33ms -run=^$

pprof: mathx.test.exe cpu.prof
	@go tool pprof -http=:8080 -no_browser mathx.test.exe cpu.prof

clean:
	@rm -rf mathx.test.exe cpu.prof

.PHONY: lint bench benchmem benchcpu pprof clean
