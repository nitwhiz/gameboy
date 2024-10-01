PWD := $(shell pwd)
GO_IMAGE_TAG := 1.23.1-alpine
GO := docker run --rm -v $(PWD):/source --workdir=/source golang:$(GO_IMAGE_TAG) go

BENCHMARK_COUNT ?= 8

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
	$(GO) test -count=$(BENCHMARK_COUNT) -run='^$$' -bench=. ./... > testdata/benchmarks/baseline.new.txt
	mv -f testdata/benchmarks/baseline.new.txt testdata/benchmarks/baseline.txt

.PHONY: benchstat_image
benchstat_image:
	docker build --tag benchstat docker/benchstat

.PHONY: benchmark
benchmark: benchstat_image
	mkdir -p testdata/benchmarks
	$(GO) test -count=$(BENCHMARK_COUNT) -run='^$$' -bench=. ./... > testdata/benchmarks/latest.txt
	docker run --rm -v $(PWD):/source --workdir=/source benchstat testdata/benchmarks/baseline.txt testdata/benchmarks/latest.txt
