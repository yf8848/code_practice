#!/bin/python3
 
from array import array
from random import random

floats = array('d', [random() for i in range(10**7)])
print(floats[-1])

fp = open('array.bin', 'wb')
floats.tofile(fp)
fp.close()

floats2 = array('d')
fp = open('array.bin', 'rb')
floats2.fromfile(fp, 10**7)
fp.close()
print(floats2[-1])

if floats == floats2:
    print("two array equal.")
