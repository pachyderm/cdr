from collections import deque
from typing import Optional

from .cipher import CipherMiddleware
from .compress import CompressionMiddleware
from .content_hash import ContentHashMiddleware
from .http import HTTPMiddleware
from cdr_pb2 import Ref


def create_ref(
    data: bytes,
    destination: HTTPMiddleware,
    cipher: Optional[CipherMiddleware] = None,
    compression: Optional[CompressionMiddleware] = None,
    content_hash: Optional[ContentHashMiddleware] = None,
) -> Ref:
    """ Create a Ref by storing data at the specified destination. """
    ref_maker_deque = deque([])
    if compression is not None:
        data = compression.compress(data)
        ref_maker_deque.appendleft(compression.ref_maker())
    if cipher is not None:
        data = cipher.encrypt(data)
        ref_maker_deque.appendleft(cipher.ref_maker())
    if content_hash is not None:
        ref_maker_deque.appendleft(content_hash.ref_maker(data))

    ref = destination(data)
    for ref_maker in ref_maker_deque:
        ref = ref_maker(ref)

    return ref
