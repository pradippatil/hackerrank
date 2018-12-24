#!/bin/python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-11-28 23:18:18
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-11-28 23:38:09
# @Description: https://www.hackerrank.com/challenges/counting-valleys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup

import math
import os
import random
import re
import sys


def countingValleys(n, s):
    level = 0
    valley_count = 0
    hike = {'U': 1, 'D': -1}

    for step in s:
        level += hike[step]

        if step == 'U' and level == 0:
            valley_count += 1

    return valley_count


if __name__ == '__main__':
    n = int(input())
    s = input()

    result = countingValleys(n, s)
    print(result)
