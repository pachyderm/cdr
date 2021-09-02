# Common Data Refs

## Contributing
To run the tests use `make run_minio` to run a minio container.
Then run the tests with `make test`

To create a test reference with arbitrary data using minio run `./targets/create-test-ref.`
```
make targets/create-test-ref
echo "test data" | ./targets/create-test-ref 127.0.0.1:9000 | base64
```
