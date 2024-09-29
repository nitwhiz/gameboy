PWD := $(shell pwd)
GO_IMAGE_TAG := 1.23.1-alpine
GO := docker run --rm -v $(PWD):/source --workdir=/source golang:$(GO_IMAGE_TAG) go

.PHONY: clean
clean:
	rm -rf ./roms/mooneye

.PHONY: mooneye_test_roms
mooneye_test_roms:
	docker build --tag mooneye-test-suite ./docker/testdata
	docker run --rm -v $(PWD)/testdata/roms:/roms mooneye-test-suite /copy_mooneye_tests.sh

.PHONY: go_generate
go_generate:
	$(GO) generate ./...

.PHONY: generate
generate: clean mooneye_test_roms go_generate

.PHONY: benchmark_baseline
benchmark_baseline:
	mkdir -p testdata/benchmarks
	$(GO) test -count=10 -run='^$$' -bench=. ./... > testdata/benchmarks/baseline.txt

.PHONY: benchmark
benchmark:
	mkdir -p testdata/benchmarks
	$(GO) test -count=10 -run='^$$' -bench=. ./... > testdata/benchmarks/latest.txt
	docker build --tag benchstat docker/benchstat
	docker run --rm -v $(PWD):/source --workdir=/source benchstat testdata/benchmarks/baseline.txt testdata/benchmarks/latest.txt
