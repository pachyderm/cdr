from collections import deque
from functools import partial
from io import BytesIO
from typing import Optional

from .cipher import CipherMiddleware
from .compress import CompressionMiddleware
from .content_hash import ContentHashMiddleware
from .http import HTTPMiddleware
from cdr_pb2 import Ref, Concat


def create_ref(
    data: bytes,
    destination: HTTPMiddleware,
    cipher: Optional[CipherMiddleware] = None,
    compression: Optional[CompressionMiddleware] = None,
    content_hash: Optional[ContentHashMiddleware] = None,
    chunk_size: Optional[int] = None
) -> Ref:
    """ Create a Ref by storing data at the specified destination.

    Args:
        data: The data to be stored and referenced.
        destination: An IO middleware component for storing data and
            creating the base Ref for the CDR chain.
        cipher: The (optional) cipher middleware.
        compression: The (optional) compression middleware.
        content_hash: The (optional) hashing middleware.
        chunk_size: If provided, the data will be chunked into segments
            of, at most, the specified size (in bytes).

    Returns:
        A complete CDR pointing to the stored object.
    """
    ref_maker_deque = deque([])

    def transform_data(data_: bytes) -> bytes:
        if compression is not None:
            data_ = compression.compress(data_)
            ref_maker_deque.append(compression.ref_maker())
        if cipher is not None:
            data_ = cipher.encrypt(data_)
            ref_maker_deque.append(cipher.ref_maker())
        if content_hash is not None:
            ref_maker_deque.append(content_hash.ref_maker(data_))
        return data_

    def construct_ref(r: Ref) -> Ref:
        while ref_maker_deque:
            ref_maker = ref_maker_deque.pop()
            r = ref_maker(r)
        return r

    if chunk_size and len(data) > chunk_size:
        chunker = partial(BytesIO(data).read, chunk_size)
        original_key = destination.key
        refs = []
        for index, chunk in enumerate(iter(chunker, b'')):
            chunk = transform_data(chunk)
            destination.key = f"{original_key}-{index:06d}"
            ref = destination(chunk)
            refs.append(construct_ref(ref))
        return Ref(concat=Concat(refs=refs))
    else:
        data = transform_data(data)
        ref = destination(data)
        return construct_ref(ref)
