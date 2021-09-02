
.PHONY: proto run_minio test

test:
	go test -v ./...

proto:
	./build_protobuf.sh

run_minio:
	./run_minio.sh


