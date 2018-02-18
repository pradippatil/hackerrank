#!/usr/bin/env pyhton3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-18 14:45:42
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 14:47:41
# @Description: https://www.hackerrank.com/challenges/text-wrap/tutorial

import textwrap

def wrap(string, max_width):
    return textwrap.fill(string, max_width)

if __name__ == '__main__':
    string, max_width = input(), int(input())
    result = wrap(string, max_width)
    print(result)

