
.PHONY: build
build: clean
	go run . build

.PHONY: clean
clean:
	rm -rf publish

.PHONY: serve
serve: 
	go run . serve -v
