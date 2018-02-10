#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-10 21:48:06
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-10 22:05:31
# @Description: https://www.hackerrank.com/challenges/python-loops/problem

if __name__ == '__main__':
    n = int(input())
    i = 0
    while i < n:
        print(i * i)
        i += 1

    # Another way using list comprehenstions
    # print(*[i**2 for i in range(n)], sep='\n')
