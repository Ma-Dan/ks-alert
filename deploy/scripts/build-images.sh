# !/bin/bash
echo "build binary..."
CGO_ENABLED=0 go build -o dispatcher ../../cmd/alert-dispatcher/main.go
CGO_ENABLED=0 go build -o executor ../../cmd/alert-executor/main.go
echo "Building images..."
docker build -t carmanzhang/alerting-dev:latest -f ./Dockerfile.dev .
echo "Built successfully"
docker push carmanzhang/alerting-dev:latest
echo "Push successfully"
