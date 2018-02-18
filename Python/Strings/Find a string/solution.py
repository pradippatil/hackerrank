#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-17 14:56:02
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 13:51:34
# @Description: https://www.hackerrank.com/challenges/find-a-string/problem

def count_substring(string, sub_string):
    start, count = 0, 0
    n = string.find(sub_string)
    while n >= 0:
        count += 1
        start += n+1
        n = string[start:].find(sub_string)
        if n == -1:
            break
    return count


if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()
    
    count = count_substring(string, sub_string)
    print(count)
