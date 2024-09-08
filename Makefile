PWD := $(shell pwd)
GO_IMAGE_TAG := 1.23.0-alpine
GO := docker run --rm -v $(PWD):/source --workdir=/source golang:$(GO_IMAGE_TAG) go

.PHONY: clean
clean:
	rm -rf /roms/mooneye

.PHONY: mooneye_test_roms
mooneye_test_roms:
	docker build --tag mooneye-test-suite ./docker/testdata
	docker run --rm -v $(PWD)/testdata/roms:/roms mooneye-test-suite /copy_mooneye_tests.sh

.PHONY: go_generate
go_generate:
	$(GO) generate ./...

.PHONY: generate
generate: clean mooneye_test_roms go_generate
