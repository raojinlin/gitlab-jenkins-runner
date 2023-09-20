build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gitlab-jenkins-runner .
	GOOS=linux GOARCH=arm64 go build -o ./bin/gitlab-jenkins-runner-arm64 .
gzip:
	gzip -f ./bin/gitlab-jenkins-runner-arm64
	gzip -f ./bin/gitlab-jenkins-runner

all: build gzip
.phony: all
