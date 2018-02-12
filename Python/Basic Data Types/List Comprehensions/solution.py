#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-12 23:27:17
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-12 23:51:58
# @Description: https://www.hackerrank.com/challenges/list-comprehensions/problem

if __name__ == '__main__':
    x, y, z, n = (int(input()) for _ in range(4))

    print([[i, j, k] for i in range(x + 1) for j in range(y + 1)
           for k in range(z + 1) if i + j + k != n])
