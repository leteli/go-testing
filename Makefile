test:
	- go test -v ./...
test_coverage:
	- go test ./... -cover
test_coverage_report:
	- go test ./... -coverprofile=coverage.out
test_cover_analyze:
	- go tool cover -func=coverage.out