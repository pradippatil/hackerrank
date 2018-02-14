#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-14 23:22:46
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-14 23:24:16
# @Description: https://www.hackerrank.com/challenges/whats-your-name/problem


def print_full_name(a, b):
    print("Hello {} {}! You just delved into python.".format(a, b))


if __name__ == '__main__':
    first_name = input()
    last_name = input()
    print_full_name(first_name, last_name)
