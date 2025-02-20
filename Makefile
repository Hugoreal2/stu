build:
	go build ./cmd/stud

clean:
	rm stud

test:
	go test ./...

.PHONY: build clean test
