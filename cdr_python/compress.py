import gzip
from abc import ABC, abstractmethod
from typing import Callable, Dict, Optional, Type

from cdr_pb2 import Ref, Compress, CompressAlgo


class CompressionMiddleware(ABC):
    """ Abstract base class for compression middleware. """

    # The algorithm corresponding to the middleware. Must be defined.
    algorithm: CompressAlgo = NotImplemented

    _subclasses_registry: Dict["CompressAlgo", Type["CompressionMiddleware"]] = dict()

    def __init_subclass__(cls) -> None:
        """ This method is called when subclasses of this object are instantiated.
        Existence and uniqueness of the algorithm class attribute is checked here.
        """
        algorithm = getattr(cls, "algorithm", None)
        if not algorithm:
            raise NotImplementedError(
                f"Class {cls.__qualname__} lacks required `algorithm` class attribute"
            )
        if algorithm in cls._subclasses_registry.keys():
            raise ValueError(f"ContentHashMiddleware already exists for for algorithm: {algorithm}")
        cls._subclasses_registry[algorithm] = cls

    @staticmethod
    @abstractmethod
    def compress(data: bytes) -> bytes:
        """ Compress the input data. """
        ...

    @staticmethod
    @abstractmethod
    def decompress(data: bytes) -> bytes:
        """ Decompress the input data. """
        ...

    @classmethod
    def ref_maker(cls) -> Callable[[Ref], Ref]:
        """ Creates a function that will create a Ref object. """
        def inner(inner_ref: Ref):
            return Ref(
                compress=Compress(
                    inner=inner_ref,
                    algo=cls.algorithm,
                )
            )

        return inner

    @classmethod
    def select(cls, algorithm: CompressAlgo) -> Optional[Type["CompressionMiddleware"]]:
        """ Given a hashing algorithm, return the corresponding middleware class. """
        return cls._subclasses_registry.get(algorithm, None)


class GZip(CompressionMiddleware):
    algorithm = CompressAlgo.GZIP

    @staticmethod
    def compress(data: bytes) -> bytes:
        """ Compress the input data. """
        return gzip.compress(data)

    @staticmethod
    def decompress(data: bytes) -> bytes:
        """ Decompress the input data. """
        return gzip.decompress(data)
