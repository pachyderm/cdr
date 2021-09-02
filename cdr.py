from typing import Iterator
from cdr_pb2 import Ref


class Resolver:

    def deref(self, ref: Ref) -> Iterator:
        """ NOTE: I would implement this as a context manager, but this is a simple stand in. """
        raise NotImplementedError


def is_immutable(ref: Ref) -> bool:
    raise NotImplementedError


def min_size(ref: Ref) -> int:
    raise NotImplementedError


def max_size(ref: Ref) -> int:
    raise NotImplementedError
