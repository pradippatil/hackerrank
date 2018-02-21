#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-20 23:00:39
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-21 00:12:14
# @Description: https://www.hackerrank.com/challenges/py-the-captains-room/problem

if __name__ == '__main__':
    K, R = int(input()), list(input().split())
    f, s = set(), set()
    for x in R:
        if x in f:
            s.add(x)
        else:
            f.add(x)

    print((f - s).pop())
    # l = len(R)
    # print((set(R[l // 2:]) - set(R[:l // 2])).pop())
