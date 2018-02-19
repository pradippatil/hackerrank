#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-19 22:53:17
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-19 23:26:23
# @Description: https://www.hackerrank.com/challenges/py-set-discard-remove-pop/problem

if __name__ == '__main__':
    n = int(input())
    s = set(map(int, input().split()))

    # Simple way
    # for _ in range(int(input())):
    #     a = input().split()
    #     if a[0] == 'pop':
    #         s.pop()
    #     elif a[0] == 'remove':
    #         s.remove(int(a[1]))
    #     elif a[0] == 'discard':
    #         s.discard(int(a[1]))

    # More efficient way using 'getattr'
    for _ in range(int(input())):
        method, *args = input().split()
        getattr(s, method)(*map(int, args))

    print(sum(s))
