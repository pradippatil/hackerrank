# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-12-24 22:22:19
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-12-24 23:32:05
# @Description: https://www.hackerrank.com/challenges/alphabet-rangoli/problem


def print_rangoli(size):
    import string
    alphabets = string.ascii_lowercase
    rows = (size*2) - 1
    width = (rows*2) - 1
    s = alphabets[:size]
    for i in range(1, size+1):
        r = s[::-1][:i]
        print('-'.join(r+r[:i-1][::-1]).center(width, '-'))

    for i in range(size-1, 0, -1):
        r = s[::-1][:i]
        print('-'.join(r+r[:i-1][::-1]).center(width, '-'))


if __name__ == '__main__':
    n = int(input())
    print_rangoli(n)
