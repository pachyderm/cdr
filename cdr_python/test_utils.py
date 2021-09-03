import pytest

from cdr_pb2 import Ref, Cipher, Concat, ContentHash, HTTP, SizeLimits, Slice
from . import utils


@pytest.mark.parametrize(
    'ref,expected',
    [(Ref(http=HTTP()), None),
     (Ref(size_limits=SizeLimits(max=128, inner=Ref(http=HTTP()))), 128),
     (Ref(concat=Concat(refs=[
          Ref(size_limits=SizeLimits(max=128)),
          Ref(size_limits=SizeLimits(max=128)),
      ])), 256),
     ]
)
def test_max_size(ref: Ref, expected: int) -> None:
    assert utils.max_size(ref) == expected


@pytest.mark.parametrize(
    'ref,expected',
    [(Ref(http=HTTP()), None),
     (Ref(size_limits=SizeLimits(min=12, inner=Ref(http=HTTP()))), 12),
     (Ref(concat=Concat(refs=[
         Ref(size_limits=SizeLimits(min=12)),
         Ref(size_limits=SizeLimits(min=12)),
     ])), 24),
     ]
)
def test_min_size(ref: Ref, expected: int) -> None:
    assert utils.min_size(ref) == expected