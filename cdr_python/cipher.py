from abc import ABC, abstractmethod
from typing import Callable, Dict, Optional, Type

from cryptography.hazmat.primitives.ciphers import Cipher, algorithms

from cdr_pb2 import Ref, Cipher as CipherRef, EncAlgo


class CipherMiddleware(ABC):
    """ Abstract base class for cipher middleware. """

    # The algorithm corresponding to the middleware. Must be defined.
    algorithm: EncAlgo = NotImplemented

    def __init__(self, key: bytes, nonce: bytes):
        self.key = key
        self.nonce = nonce

    _subclasses_registry: Dict["EncAlgo", Type["CipherMiddleware"]] = dict()

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

    @abstractmethod
    def encrypt(self, data: bytes) -> bytes:
        """ Encrypt the input data. """
        ...

    @abstractmethod
    def decrypt(self, data: bytes) -> bytes:
        """ Decrypt the input data. """
        ...

    def ref_maker(self) -> Callable[[Ref], Ref]:
        """ Creates a function that will create a Ref object. """

        def inner(inner_ref: Ref):
            return Ref(
                cipher=CipherRef(
                    inner=inner_ref,
                    algo=self.algorithm,
                    key=self.key,
                    nonce=self.nonce,
                )
            )

        return inner

    @classmethod
    def select(cls, algorithm: EncAlgo) -> Optional[Type["CipherMiddleware"]]:
        """ Given a hashing algorithm, return the corresponding middleware class. """
        return cls._subclasses_registry.get(algorithm, None)


class ChaCha20(CipherMiddleware):
    algorithm = EncAlgo.CHACHA20

    def encrypt(self, data: bytes) -> bytes:
        """ Encrypt the input data. """
        cipher = Cipher(algorithms.ChaCha20(self.key, self.nonce), mode=None)
        return cipher.encryptor().update(data)

    def decrypt(self, data: bytes) -> bytes:
        """ Decrypt the input data. """
        cipher = Cipher(algorithms.ChaCha20(self.key, self.nonce), mode=None)
        return cipher.decryptor().update(data)
