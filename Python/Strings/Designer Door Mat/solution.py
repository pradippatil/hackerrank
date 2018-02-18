#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-18 15:33:54
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 15:53:43
# @Description: https://www.hackerrank.com/challenges/designer-door-mat/problem


if __name__ == '__main__':
    N, M = map(int,input().split()) # More than 6 lines of code will result in 0 score. Blank lines are not counted.
    for i in range(1,N,2):
        print((".|."*i).center(M,'-'))
    print('WELCOME'.center(M,'-'))
    for i in range(N-2,-1,-2):
        print((".|."*i).center(M,'-'))

