#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-14 23:04:10
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-14 23:09:07
# @Description: https://www.hackerrank.com/challenges/swap-case/problem


def swap_case(s):
    return s.swapcase()


if __name__ == '__main__':
    s = input()
    result = swap_case(s)
    print(result)
