from abc import ABC, abstractmethod
from datetime import timedelta
from io import BytesIO
from typing import Any, Optional

from minio import Minio

from cdr_pb2 import Ref, HTTP


class HTTPMiddleware(ABC):
    """ Abstract base class for http middleware. """

    def __init__(
        self,
        client: Any,
        bucket: str,
        key: str,
        expires: Optional[timedelta]
    ):
        self.client = client
        self.bucket = bucket
        self.key = key
        self.expires = expires

    @abstractmethod
    def _post(self, data: bytes) -> None:
        ...

    @abstractmethod
    def _make_ref(self) -> HTTP:
        ...

    def __call__(self, data: bytes) -> Ref:
        """ This method posts the specified data and returns a Ref to that posted location. """
        self._post(data)
        body = self._make_ref()
        return Ref(http=body)


class MinioMiddleware(HTTPMiddleware):

    def __init__(
        self,
        client: Minio,
        bucket: str,
        key: str,
        expires: Optional[timedelta]
    ):
        super().__init__(client, bucket, key, expires)

    def _post(self, data: bytes) -> None:
        """ Performs the hashing on input data. """
        stream = BytesIO(data)
        self.client.put_object(self.bucket, self.key, stream, len(data))

    def _make_ref(self) -> HTTP:
        url = self.client.presigned_get_object(self.bucket, self.key, self.expires)
        return HTTP(url=url)
