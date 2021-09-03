from datetime import timedelta
from os import urandom
from typing import Optional

import pytest
from minio import Minio
from minio.credentials import StaticProvider

from .cipher import CipherMiddleware, ChaCha20
from .compress import CompressionMiddleware, GZip
from .content_hash import ContentHashMiddleware, HashBlake2b256
from .creator import create_ref
from .http import MinioMiddleware
from .resolver import Resolver


@pytest.fixture(name="bucket_name")
def _bucket_name(request: pytest.FixtureRequest) -> str:
    return request.node.originalname.replace("_", "-")


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


##### Middleware Parametrizers ############################

@pytest.fixture(name="cipher", params=[None, ChaCha20(key=urandom(32), nonce=urandom(24))])
def _cipher(request) -> Optional[CipherMiddleware]:
    return request.param


@pytest.fixture(name="compression", params=[None, GZip()])
def _compression(request) -> Optional[CompressionMiddleware]:
    return request.param


@pytest.fixture(name="content_hash", params=[None, HashBlake2b256()])
def _content_hash(request) -> Optional[ContentHashMiddleware]:
    return request.param


@pytest.mark.parametrize("chunk_size", [None, 5, 100])
def test_resolver(
    client: Minio,
    bucket_name: str,
    cipher: Optional[CipherMiddleware],
    compression: Optional[CompressionMiddleware],
    content_hash: Optional[ContentHashMiddleware],
    chunk_size: Optional[int]
) -> None:
    # Arrange
    original = b"testing..." * 3
    ref = create_ref(
        data=original,
        destination=MinioMiddleware(
            client=client,
            bucket=bucket_name,
            key="test-key",
            expires=timedelta(minutes=5)
        ),
        cipher=cipher,
        compression=compression,
        content_hash=content_hash,
        chunk_size=chunk_size
    )
    resolver = Resolver()

    # Act
    result = resolver.dereference(ref)

    # Assert
    assert result == original
