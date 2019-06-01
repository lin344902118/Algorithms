# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/1
"""

import random

def quick_sort(arr: list):
    if len(arr) < 2:
        return arr
    num = arr[0]
    less = []
    greater = []
    for i in range(1, len(arr)):
        if arr[i] <= num:
            less.append(arr[i])
        else:
            greater.append(arr[i])
    return quick_sort(less) + [num] + quick_sort(greater)

def main():
    nums = [i for i in range(100)]
    for i in range(5):
        arr = []
        for i in range(10):
            arr.append(random.choice(nums))
        print(quick_sort(arr))

if __name__ == '__main__':
    main()

