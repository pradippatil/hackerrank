#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-13 23:44:25
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-14 00:08:53
# @Description: https://www.hackerrank.com/challenges/finding-the-percentage/problem

if __name__ == '__main__':
    n = int(input())
    marks = dict()
    for _ in range(n):
        s = input().split()
        marks[s[0]] = list(map(float, s[1:]))

    print("{:.2f}".format(sum(marks[input()]) / 3))
