
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

.PHONY: deploy
deploy: publish/_redirects
	rsync -rv publish/ pgs.sh:/andydaytondotcom/

# .PHONY: deploy
# deploy: build
# 	aws --endpoint-url=https://s3.pub1.infomaniak.cloud s3 sync publish s3://andydaytondotcom


# non-phony targets below

publish: build

publish/_redirects: publish
	echo "/*  /index.html  200" > publish/_redirects