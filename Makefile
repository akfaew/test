TEST_ARGS = -failfast

fmt:
	go fmt ./...

test: fmt
	go test $(TEST_ARGS) ./...

test-cover: fmt
	go test $(TEST_ARGS) -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out |\
		grep -v 100.0% |\
		grep -v total: |\
		perl -nae 'printf("%7s %s %s\n", $$F[2], $$F[0], $$F[1])' | sort -nr
	go tool cover -html=coverage.out
