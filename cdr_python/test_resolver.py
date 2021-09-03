import os
from datetime import timedelta

import pytest
from minio import Minio
from minio.credentials import StaticProvider

from .cipher import ChaCha20
from .compress import GZip
from .content_hash import HashBlake2b256
from .creator import create_ref
from .http import MinioMiddleware
from .resolver import Resolver


@pytest.fixture(name="bucket_name")
def _bucket_name(request: pytest.FixtureRequest) -> str:
    return (
        request.node.name
        .replace("_", "-")
        .replace("[", ".")
        .replace("]", ".")
    )


@pytest.fixture(name="client")
def _minio_client(bucket_name: str) -> Minio:

    def safe_delete_bucket(c: Minio) -> None:
        if c.bucket_exists(bucket_name):
            for obj in client.list_objects(bucket_name):
                client.remove_object(bucket_name, obj.object_name)
            client.remove_bucket(bucket_name)

    client = Minio(
        endpoint=f"127.0.0.1:9000",
        credentials=StaticProvider("minioadmin", "minioadmin"),
        secure=False
    )
    safe_delete_bucket(client)
    client.make_bucket(bucket_name)
    try:
        yield client
    finally:
        safe_delete_bucket(client)


def test_resolver(client: Minio, bucket_name: str) -> None:
    # Arrange
    original = b"testing..."
    ref = create_ref(
        data=original,
        destination=MinioMiddleware(client, bucket_name, "test-key", timedelta(minutes=5)),
        cipher=ChaCha20(key=os.urandom(32), nonce=os.urandom(16)),
        compression=GZip(),
        content_hash=HashBlake2b256()
    )
    resolver = Resolver()

    # Act
    result = resolver.dereference(ref)

    # Assert
    assert result == original
