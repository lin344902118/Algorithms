# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/1
"""

import random

def find_smallest_num(arr: list):
    smallest_num = arr[0]
    for num in arr:
        if num < smallest_num:
            smallest_num = num
    return smallest_num

def select_sort(arr: list):
    new_arr = []
    for i in range(len(arr)):
        smallest_num = find_smallest_num(arr)
        new_arr.append(smallest_num)
        arr.remove(smallest_num)
    return new_arr

def main():
    nums = [i for i in range(100)]
    for i in range(5):
        arr = []
        for i in range(10):
            arr.append(random.choice(nums))
        print(select_sort(arr))

if __name__ == '__main__':
    main()