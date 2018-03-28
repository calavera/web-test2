build:
	echo "${GO_IMPORT_PATH}"
	mkdir -p functions
	go get -v ./...
	go build -o functions/hello-buildbot cmd/hello-buildbot/main.go
	go build -o functions/submission-created cmd/submission-created/main.go
	go build -o functions/deploy-succeeded cmd/deploy-succeeded/main.go
	echo "done"
