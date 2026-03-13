.PHONY: test test-unit test-e2e test-all

test: test-unit

test-unit:
	go test ./... -v -count=1

test-e2e:
	go test ./test/e2e -v -count=1

test-all:
	go test ./... -v -count=1
	go test ./test/e2e -v -count=1
