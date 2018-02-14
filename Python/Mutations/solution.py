#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-14 23:26:27
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-14 23:31:21
# @Description: https://www.hackerrank.com/challenges/python-mutations/problem


def mutate_string(string, position, character):
    return string[:position] + character + string[position + 1:]


if __name__ == '__main__':
    s = input()
    i, c = input().split()
    s_new = mutate_string(s, int(i), c)
    print(s_new)
