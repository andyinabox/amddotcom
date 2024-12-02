
.PHONY: build
build: clean
	go run cmd/build/main.go

.PHONY: clean
clean:
	rm -rf publish
