from abc import ABC, abstractmethod
from hashlib import blake2b
from hmac import compare_digest
from typing import Callable, Dict, Optional, Type

from cdr_pb2 import Ref, ContentHash, HashAlgo


class ContentHashMiddleware(ABC):
    """ Abstract base class for content hash middleware. """

    # The algorithm corresponding to the middleware. Must be defined.
    algorithm: HashAlgo = NotImplemented

    _subclasses_registry: Dict["HashAlgo", Type["ContentHashMiddleware"]] = dict()

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
    def hash(data: bytes) -> bytes:
        """ Performs the hashing on input data. """
        ...

    @classmethod
    def ref_maker(cls, data: bytes) -> Callable[[Ref], Ref]:
        """ Creates a function that will create a Ref object. """
        signature = cls.hash(data)

        def inner(inner_ref: Ref):
            return Ref(
                content_hash=ContentHash(
                    inner=inner_ref,
                    algo=cls.algorithm,
                    hash=signature,
                )
            )
        return inner

    @classmethod
    def select(cls, algorithm: HashAlgo) -> Optional[Type["ContentHashMiddleware"]]:
        """ Given a hashing algorithm, return the corresponding middleware class. """
        return cls._subclasses_registry.get(algorithm, None)

    @classmethod
    def verify(cls, data: bytes, expected: bytes) -> bool:
        """ Verify a content hash. """
        signature = cls.hash(data)
        return compare_digest(signature, expected)


class HashBlake2b256(ContentHashMiddleware):

    algorithm = HashAlgo.BLAKE2b_256

    @staticmethod
    def hash(data: bytes) -> bytes:
        """ Performs the hashing on input data. """

        hasher = blake2b(data, digest_size=32)
        return hasher.hexdigest().encode('utf-8')
