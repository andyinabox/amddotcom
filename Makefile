
.PHONY: build
build: clean
	go run cmd/build/main.go

.PHONY: clean
clean:
	rm -rf publish

# publish/content.json:
# 	go run cmd/parse-content/main.go -i content

