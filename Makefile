
.PHONY: proto run_minio test

test:
	go test -v ./...

proto:
	./etc/build_protobuf.sh

run-minio:
	./etc/run_minio.sh

targets/create-test-ref: etc/create-test-ref/*
	go build -o ./targets/create-test-ref ./etc/create-test-ref
