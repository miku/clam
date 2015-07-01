test:
	go test -v .

imports:
	goimports -w .

cover:
	go test -coverprofile coverage.out
	go tool cover -html=coverage.out

clean:
	rm -f coverage.out
