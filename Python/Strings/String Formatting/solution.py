#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-18 16:32:38
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 17:07:34
# @Description: https://www.hackerrank.com/challenges/python-string-formatting/problem

def print_formatted(number):
    for i in range(1, number+1):
        print("{0:>{w}d} {0:>{w}o} {0:>{w}X} {0:>{w}b}".format(i, w=len(bin(number)[2:])))


if __name__ == '__main__':
    n = int(input())
    print_formatted(n)
