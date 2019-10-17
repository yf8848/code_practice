#!/bin/.python3

from collections import deque

dq = deque(range(10), maxlen=10)
print(dq)
dq.rotate(5)
print(dq)

