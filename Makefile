
.PHONY: proto run_minio test

test:
	go test -v ./...

proto:
	./etc/build_protobuf.sh

run-minio:
	./etc/run_minio.sh

targets/create-test-ref: etc/create-test-ref/*
	go build -o ./targets/create-test-ref ./etc/create-test-ref
	chmod +x ./targets/create-test-ref

targets/cdr: cmd/cdr/*
	go build -o ./targets/cdr ./cmd/cdr
	chmod +x ./targets/cdr
