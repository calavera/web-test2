build:
	ls -la .
	go get -v ./...
	go build -o /tmp/foo ./...
	echo "done"