from typing import Union

import cdr_pb2 as pb

COMMON_DATA_REF = Union[pb.Cipher, pb.Concat, pb.ContentHash, pb.HTTP, pb.SizeLimits, pb.Slice]
