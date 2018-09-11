DEP=dep

build: recompile

recompile: ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	serverless sam export --output ./template.yml

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose

.PHONY: ensure
ensure:
	$(DEP) ensure -vendor-only

supervise:
	./node_modules/.bin/supervisor --no-restart-on exit -e go -i bin -i vendor --exec make -- recompile

watch:
	./scripts/parallel "make supervise" "make start-local"

start-local:
	sam local start-api
