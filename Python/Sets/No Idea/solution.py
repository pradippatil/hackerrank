#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-19 22:08:57
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-19 22:34:23
# @Description: https://www.hackerrank.com/challenges/no-idea/problem

if __name__ == '__main__':
    n, m = map(int, input().split())
    N = list(map(int, input().split()))
    A = set(map(int, input().split()))
    B = set(map(int, input().split()))

    happiness = 0

    for e in N:
        if e in A:
            happiness += 1
        if e in B:
            happiness -= 1

    print(happiness)
