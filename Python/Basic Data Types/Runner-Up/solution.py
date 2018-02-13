#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-12 23:58:20
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-13 00:15:16
# @Description: https://www.hackerrank.com/challenges/find-second-maximum-number-in-a-list/problem

if __name__ == '__main__':
    n = int(input())

    # 1. create list fom space separated input
    # 2. remove duplicates by creating set from list
    # 3. sort the set and print second last element
    print(sorted(set([int(i) for i in input().split()]))[-2])
