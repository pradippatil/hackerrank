#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-19 22:43:37
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-19 22:50:17
# @Description: https://www.hackerrank.com/challenges/py-set-add/problem

if __name__ == '__main__':
    C = set()
    for _ in range(int(input())):
        C.add(input().strip())

    print(len(C))
