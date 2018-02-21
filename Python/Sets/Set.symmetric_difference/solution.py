#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-20 22:12:46
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-20 22:31:27
# @Description: https://www.hackerrank.com/challenges/py-set-symmetric-difference-operation/problem


if __name__ == '__main__':
    _, E = input(), set(input().split())
    _, F = input(), set(input().split())
    print(len(E ^ F))
    # Or
    # print(len(E.symmetric_difference(F)))
