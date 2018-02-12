#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-11 12:56:58
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-12 00:23:57
# @Description: https://www.hackerrank.com/challenges/python-lists/problem

if __name__ == '__main__':
    n = int(input())
    l = []
    while n > 0:
        i = input().split()

        if i[0] == 'insert':
            l.insert(int(i[1]), int(i[2]))
        elif i[0] == 'remove':
            l.remove(int(i[1]))
        elif i[0] == 'append':
            l.append(int(i[1]))
        elif i[0] == 'sort':
            l.sort()
        elif i[0] == 'pop':
            l.pop()
        elif i[0] == 'reverse':
            l.reverse()
        elif i[0] == 'print':
            print(l)

        n -= 1
