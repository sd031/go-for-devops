.PHONY: build-GoLambdaFunction clean

build-GoLambdaFunction:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	cp ./bootstrap $(ARTIFACTS_DIR)/.

clean:
	rm -f bootstrap
