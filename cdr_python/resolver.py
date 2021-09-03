import requests

from cdr_pb2 import Ref, Cipher, Compress, Concat, ContentHash, HTTP, SizeLimits, Slice
from .cipher import CipherMiddleware
from .compress import CompressionMiddleware
from .content_hash import ContentHashMiddleware
from .utils import _get_ref_body


class Resolver:

    def dereference(self, ref: Ref) -> bytes:
        """ NOTE: I would implement this as a context manager, but this is a simple stand in. """
        body = _get_ref_body(ref)
        if isinstance(body, HTTP):
            return self._dereference_http(body)
        elif isinstance(body, Cipher):
            return self._dereference_cipher(body)
        elif isinstance(body, ContentHash):
            return self._dereference_content_hash(body)
        elif isinstance(body, Compress):
            return self._dereference_compress(body)
        else:
            raise ValueError(f"unsupported Ref variant: {body}")

    @staticmethod
    def _dereference_http(body: HTTP) -> bytes:
        response = requests.get(url=body.url, headers=body.headers)
        response.raise_for_status()
        return response.content

    def _dereference_cipher(self, body: Cipher) -> bytes:
        inner_data = self.dereference(body.inner)
        middleware_type = CipherMiddleware.select(body.algo)
        if middleware_type is None:
            raise ValueError(f"unrecognized cipher algorithm: {body.algo}")

        middleware = middleware_type(body.key, body.nonce)
        return middleware.decrypt(inner_data)

    def _dereference_compress(self, body: Compress) -> bytes:
        inner_data = self.dereference(body.inner)
        middleware = CompressionMiddleware.select(body.algo)
        if middleware is None:
            raise ValueError(f"unrecognized compression algorithm: {body.algo}")
        return middleware.decompress(inner_data)

    def _dereference_content_hash(self, body: ContentHash) -> bytes:
        inner_data = self.dereference(body.inner)
        middleware = ContentHashMiddleware.select(body.algo)
        if middleware is None:
            raise ValueError(f"unrecognized hash algorithm: {body.algo}")

        if not middleware.verify(inner_data, body.hash):
            raise ValueError(
                f"content failed hash check. "
                f"HAVE: {middleware.hash(inner_data)} "
                f"GOT: {body.hash}"
            )
        return inner_data
