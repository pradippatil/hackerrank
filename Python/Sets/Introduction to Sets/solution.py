#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-18 20:52:05
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 20:55:31
# @Description: https://www.hackerrank.com/challenges/py-introduction-to-sets/problem

def average(array):
    return sum(set(array))/len(set(array))


if __name__ == '__main__':
    n = int(input())
    arr = list(map(int, input().split()))
    result = average(arr)
    print(result)
