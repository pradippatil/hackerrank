#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-10 22:07:03
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-10 22:55:44
# @Description: https://www.hackerrank.com/challenges/write-a-function/problem

# shortest way
# year is a leap year if it is:
#   divisible by 4 and non-divisible by 100
#   or
#   divisble by 4 and divisible by 400


def is_leap(year):
    return year % 4 == 0 and (year % 100 != 0 and year % 400 == 0)

# verbose
# def is_leap(year):
#    leap = False
#    if year % 4 == 0:
#        leap = True
#        if year % 100 == 0:
#            if year % 400 == 0:
#                leap = True
#            else:
#                leap = False
#    return leap


# if __name__ == '__main__':
    # year = int(input())
    # print(is_leap(year))
