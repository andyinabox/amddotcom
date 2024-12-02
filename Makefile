
.PHONY: build
build: clean
	go run cmd/build/main.go

.PHONY: clean
clean:
	rm -rf publish

.PHONY: watch
watch:
	reflex -G 'publish/*' make

.PHONY: serve
serve:
	npx servor publish/ --reload