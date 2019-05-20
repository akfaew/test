TEST_ARGS = -failfast

fmt:
	go fmt ./...

lint:
	golangci-lint run

test: fmt lint
	go test $(TEST_ARGS) ./...

test-cover: fmt
	go test $(TEST_ARGS) -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

push: test
	git push
	git push --tags

clean:
	rm coverage.out
