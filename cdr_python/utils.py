from typing import List, Optional, Type, TypeVar

from cdr_pb2 import Ref, Cipher, Compress, Concat, ContentHash, HTTP, SizeLimits, Slice
from . import COMMON_DATA_REF

B = TypeVar("B", bound=COMMON_DATA_REF)


def is_immutable(ref: Ref) -> bool:
    """ Determine whether a reference is immutable (data has corresponding hash). """
    body = _get_ref_body(ref)
    if isinstance(body, ContentHash) or body is None:
        return True
    if isinstance(body, Concat):
        return all(map(is_immutable, body.refs))
    if isinstance(body, (Cipher, Compress, SizeLimits, Slice)):
        return is_immutable(body.inner)
    return False


def min_size(ref: Ref) -> Optional[int]:
    """ Determine the minimum size of the data specified by the CDR. """
    size_limit_refs = _collect_all_size_limit_refs(ref)
    return sum(getattr(r, "min", 0) for r in size_limit_refs) or None


def max_size(ref: Ref) -> Optional[int]:
    """ Determine the maximum size of the data specified by the CDR. """
    size_limit_refs = _collect_all_size_limit_refs(ref)
    return sum(getattr(r, "max", 0) for r in size_limit_refs) or None


def _collect_all_size_limit_refs(ref: Ref) -> List[Optional[SizeLimits]]:
    """ Collect and return a list of all of the SizeLimits ref messages within the specified Ref. """
    size_limit_ref = _find_first_ref(ref, SizeLimits)
    if size_limit_ref:
        return [size_limit_ref]

    # SizeLimit refs might exist within a Concat ref.
    size_limit_refs = []
    concat_message = _find_first_ref(ref, Concat)
    if concat_message:
        for inner_ref in concat_message.refs:
            size_limit_refs.append(_find_first_ref(inner_ref, SizeLimits))
    return size_limit_refs


def _find_first_ref(ref: Ref, message_type: Type[B]) -> Optional[B]:
    """ Finds and returns (if exists) the first instance of the specified message type
    within the specified Ref.
    """
    # Get the body message, if it exists
    body = _get_ref_body(ref)
    if isinstance(body, message_type):
        return body
    if not body or isinstance(body, (Concat, HTTP)):
        # If the body message is a Concat or HTTP message stop recursion.
        return None
    return _find_first_ref(body.inner, message_type)


def _get_ref_body(ref: Ref) -> Optional[COMMON_DATA_REF]:
    field = ref.WhichOneof("body")
    if field is None:
        return None
    body = getattr(ref, field)
    return body
