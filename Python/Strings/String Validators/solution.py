#!/usr/bin/env python3

# -*- coding:utf-8 -*-
# @Script: solution.py
# @Author: Pradip Patil
# @Contact: @pradip__patil
# @Created: 2018-02-18 13:57:07
# @Last Modified By: Pradip Patil
# @Last Modified: 2018-02-18 14:18:49
# @Description: https://www.hackerrank.com/challenges/string-validators/problem

if __name__ == '__main__':
    s = input()

    # Use builtin: any(iterable)
    # Return True if any element of the iterable is true. If the iterable is empty, return False. Equivalent to:
    # https://docs.python.org/3/library/functions.html#any

    # In the first line, print True if  has any alphanumeric characters. Otherwise, print False.
    alphanumeric = any([c.isalnum() for c in s])

    # In the second line, print True if  has any alphabetical characters. Otherwise, print False.
    alphabetical = any([c.isalpha() for c in s])

    # In the third line, print True if  has any digits. Otherwise, print False.
    digits = any([c.isdigit() for c in s])

    # In the fourth line, print True if  has any lowercase characters. Otherwise, print False.
    lowercase = any([c.islower() for c in s])

    # In the fifth line, print True if  has any uppercase characters. Otherwise, print False.
    uppercase = any([c.isupper() for c in s])

    print("{}\n{}\n{}\n{}\n{}".format(alphanumeric,
                                      alphabetical, digits, lowercase, uppercase))
