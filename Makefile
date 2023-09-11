build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gitlab-jenkins-runner .
	GOOS=linux GOARCH=arm64 go build -o ./bin/gitlab-jenkins-runner-arm64 .
gzip:
	gzip ./bin/gitlab-jenkins-runner-arm64
	gzip ./bin/gitlab-jenkins-runner

.phony: build gzip
