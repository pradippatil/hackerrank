#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-20 22:43:59
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-20 22:49:12
# @Description: https://www.hackerrank.com/challenges/py-set-mutations/problem

if __name__ == '__main__':
    _, A = input(), set(input().split())
    N = int(input())
    for _ in range(N):
        getattr(A, input().split()[0])(set(input().split()))

    print(sum(map(int, A)))
