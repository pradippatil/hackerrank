#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-14 23:12:35
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-14 23:14:15
# @Description: https://www.hackerrank.com/challenges/python-string-split-and-join/problem


def split_and_join(line):
    return "-".join(line.split(" "))


if __name__ == '__main__':
    line = input()
    result = split_and_join(line)
    print(result)
