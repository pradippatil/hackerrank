#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-13 22:22:19
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-13 23:38:18
# @Description: https://www.hackerrank.com/challenges/nested-list/problem

if __name__ == '__main__':
    students = []
    for _ in range(int(input())):
        students.append([input(), float(input())])

    for i in sorted([x[0] for x in students if x[1] == sorted(set([y[1] for y in students]))[1]]):
        print(i)
