
.PHONY: build
build: clean
	go run cmd/unamica/main.go build

.PHONY: clean
clean:
	rm -rf publish

.PHONY: serve
serve:
	go run cmd/unamica/main.go serve

