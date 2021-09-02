from typing import Iterator, List, Optional, Type, TypeVar
from cdr_pb2 import Ref, Cipher, Concat, ContentHash, HTTP, SizeLimits, Slice

M = TypeVar("M")


class Resolver:

    def deref(self, ref: Ref) -> Iterator:
        """ NOTE: I would implement this as a context manager, but this is a simple stand in. """
        raise NotImplementedError


def is_immutable(ref: Ref) -> bool:
    raise NotImplementedError


def min_size(ref: Ref) -> Optional[int]:
    size_limit_refs = _collect_all_size_limit_refs(ref)
    return sum(getattr(size_limit_ref, "min", 0) for size_limit_ref in size_limit_refs)


def max_size(ref: Ref) -> Optional[int]:
    size_limit_refs = _collect_all_size_limit_refs(ref)
    return sum(getattr(size_limit_ref, "max", 0) for size_limit_ref in size_limit_refs)


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


def _find_first_ref(ref: Ref, message_type: Type[M]) -> Optional[M]:
    """ Finds and returns (if exists) the first instance of the specified message type
    within the specified Ref.
    """
    # Get the body message, if it exists
    field = ref.WhichOneof("body")
    if field is None:
        return None
    body = getattr(ref, field)

    if isinstance(body, message_type):
        return ref
    if isinstance(body, (Concat, HTTP)):
        # If the body message is a Concat or HTTP message stop recursion.
        return None
    return _find_first_ref(body.inner, message_type)
