
.PHONY: build
build: clean unamica
	./unamica build

.PHONY: clean
clean:
	rm -rf publish
	rm unamica

.PHONY: serve
serve: unamica
	./unamica serve -v


unamica:
	go build -o unamica cmd/unamica/main.go