#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-26 23:07:46
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-26 23:08:48
# @Description: https://www.hackerrank.com/challenges/py-check-subset/problem

if __name__ == '__main__':
    # More than 4 lines will result in 0 score. Blank lines won't be counted.
    for i in range(int(input())):
        a = int(input())
        A = set(input().split())
        b = int(input())
        B = set(input().split())
        print(A <= B)
