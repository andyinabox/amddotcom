
.PHONY: build
build: clean
	go run . build -v

.PHONY: clean
clean:
	rm -rf publish

.PHONY: serve
serve: 
	go run . serve -v

# .PHONY: watch
# watch: build
# 	reflex -G 'publish/*' make

# .PHONY: serve
# serve:
# 	npx servor publish/ --reload
