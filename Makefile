format:
	go mod tidy
	gofmt -w .
test:
	go test -v ./...
coverage:
	go test ./... -cover
coverage-report:
	rm -r test_report || true
	mkdir test_report
	go test -coverprofile=test_report/coverage.out ./...
	go tool cover -html=test_report/coverage.out -o=test_report/coverage.html
clean:
	go clean
	rm -r build || true
compile:
	mkdir -p build
	go build -o build
build: clean test compile
release:
	git tag $(version)
	git push origin $(version)