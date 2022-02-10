# Common Data Refs

**NOTE:** This is not an officially supported Pachyderm product.

Common Data Refs are a standard intermediary format for passing references to data between systems.

## CLI
There is a command line tool `cdr` which can be built with `make targets/cdr`.
It allows users to inpsect and dereference CDRs.
Run it with no arguments to see the help message.

## Contributing
To run the tests use `make run_minio` to run a minio container.
Then run the tests with `make test`

To create a test reference with arbitrary data using minio run `./targets/create-test-ref.`
The result is base64 encoded
```
make targets/create-test-ref
echo "test data" | ./targets/create-test-ref 127.0.0.1:9000
```
