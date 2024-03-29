TEST_ARGS = -failfast

fmt:
	go fmt ./...

lint:
	golangci-lint run

test: fmt lint
	go test $(TEST_ARGS) ./...

test-regen: fmt lint
	rm -rf testdata/output
	mkdir -p testdata/output
	go test -regen $(TEST_ARGS) ./...

test-cover: fmt
	go test $(TEST_ARGS) -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

push: test
	git push
	git push --tags

clean:
	rm coverage.out

update:
	go get -u
	go mod tidy
	go mod verify
