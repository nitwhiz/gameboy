PWD := $(shell pwd)

IMAGE_NAME_GOLANG := golang:1.23.1-alpine
IMAGE_NAME_BENCHSTAT := gameboy-benchstat
IMAGE_NAME_TESTDATA := gameboy-testdata

BENCHMARK_COUNT := 8

GO := docker run --rm -v $(PWD):/source -v gameboy_modules:/go/pkg/mod --workdir=/source $(IMAGE_NAME_GOLANG) go
BENCHSTAT := docker run --rm -v $(PWD):/source -v gameboy_modules:/go/pkg/mod --workdir=/source $(IMAGE_NAME_BENCHSTAT)
BENCHMARK := $(GO) test -count=$(BENCHMARK_COUNT) -run='^$$' -bench=.

.PHONY: clean
clean:
	rm -rf ./testdata/roms/mooneye

.PHONY: clean_images
clean_images:
	-docker image rm $(IMAGE_NAME_BENCHSTAT) $(IMAGE_NAME_TESTDATA) $(IMAGE_NAME_GOLANG)

.PHONY: image_benchstat
image_benchstat:
	docker build --tag $(IMAGE_NAME_BENCHSTAT) ./docker/benchstat

.PHONY: image_testdata_mooneye
image_testdata_mooneye:
	docker build --tag $(IMAGE_NAME_TESTDATA) --target=mooneye-test-suite ./docker/testdata

.PHONY: mooneye_test_roms
mooneye_test_roms: image_testdata_mooneye
	docker run --rm -v $(PWD)/testdata/roms:/roms $(IMAGE_NAME_TESTDATA) /copy_mooneye_tests.sh

.PHONY: go_generate
go_generate:
	$(GO) generate ./...

.PHONY: generate
generate: clean mooneye_test_roms go_generate

.PHONY: benchmark_baseline
benchmark_baseline:
	mkdir -p testdata/benchmarks
	$(BENCHMARK) -cpuprofile testdata/benchmarks/cpu.baseline.new.prof ./test/integration > testdata/benchmarks/baseline.new.txt
	mv -f testdata/benchmarks/cpu.baseline.new.prof testdata/benchmarks/cpu.baseline.prof
	mv -f testdata/benchmarks/baseline.new.txt testdata/benchmarks/baseline.txt

.PHONY: benchmark
benchmark: image_benchstat
	mkdir -p testdata/benchmarks
	$(BENCHMARK) -pgo testdata/benchmarks/cpu.baseline.prof ./test/integration > testdata/benchmarks/latest.txt
	$(BENCHSTAT) testdata/benchmarks/baseline.txt testdata/benchmarks/latest.txt
