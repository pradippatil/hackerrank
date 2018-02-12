#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-12 23:10:31
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-12 23:24:28
# @Description: https://www.hackerrank.com/challenges/python-tuples/problem

if __name__ == '__main__':
    n = int(input())
    # split space separated input and convert in int using map function
    int_list = [int(i) for i in input().split()[:n]]
    t = tuple(int_list)
    print(hash(t))
