
.PHONY: build
build: clean publish/content.json

.PHONY: clean
clean:
	rm -rf publish

publish/content.json:
	go run cmd/parse-content/main.go -i content

