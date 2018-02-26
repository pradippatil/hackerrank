#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-26 23:14:51
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-26 23:22:44
# @Description: https://www.hackerrank.com/challenges/py-check-strict-superset/problem

if __name__ == '__main__':
    A = set(input().split())
    strict_superset = True
    for _ in range(int(input())):
        strict_superset = A > set(input().split())
        if not strict_superset:
            break
    print(strict_superset)
