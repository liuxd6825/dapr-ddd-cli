export GOPATH ?= /Users/lxd/go
build:
	CGO_ENABLED=0  GOOS=darwin   GOARCH=arm64  go build  -o ./dist/darwin_arm64/dapr-ddd-cli ./main.go
	CGO_ENABLED=0  GOOS=darwin   GOARCH=amd64  go build  -o ./dist/darwin_amd64/dapr-ddd-cli ./main.go
	CGO_ENABLED=0  GOOS=windows  GOARCH=amd64  go build  -o ./dist/window_amd64/dapr-ddd-cli ./main.go
	CGO_ENABLED=0  GOOS=linux    GOARCH=amd64  go build  -o ./dist/linux_amd64/dapr-ddd-cli ./main.go
	CGO_ENABLED=0  GOOS=darwin   GOARCH=arm64  go build  -o ${GOPATH}/bin/dapr-ddd-cli ./main.go