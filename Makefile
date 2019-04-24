build:
	dep ensure -v
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda1/main pkg/lambda1/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda2/main pkg/lambda2/main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
