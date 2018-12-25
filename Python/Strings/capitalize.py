# -*- coding:utf-8 -*-
# @Script: capitalize.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-12-24 23:37:37
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-12-25 00:08:14
# @Description: https://www.hackerrank.com/challenges/capitalize/problem


def solve(s):
    return ' '.join([word.capitalize() for word in s.split(' ')])


if __name__ == '__main__':
    s = input()
    result = solve(s)
    print(result)
