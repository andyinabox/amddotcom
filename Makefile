
.PHONY: build
build: clean
	go run .

.PHONY: clean
clean:
	rm -rf publish

.PHONY: watch
watch: build
	reflex -G 'publish/*' make

.PHONY: serve
serve:
	npx servor publish/ --reload
