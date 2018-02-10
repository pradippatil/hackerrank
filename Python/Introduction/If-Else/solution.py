#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-10 21:04:45
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-10 21:12:31
# @Description: https://www.hackerrank.com/challenges/py-if-else/problem

if __name__ == '__main__':
    n = int(input())

    if n % 2 != 0:
        print("Weird")
    else:
        if n >= 2 and n <= 5:
            print("Not Weird")
        elif n >= 6 and n <= 20:
            print("Weird")
        elif n > 20:
            print("Not Weird")
