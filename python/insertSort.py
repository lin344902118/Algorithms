# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/9
"""

import random

def insert_sort(arr: list):
    sorted_arr = [arr[0]]
    for i in range(1, len(arr)):
        inserted = False
        for j in range(len(sorted_arr)):
            if arr[i] < sorted_arr[j]:
                sorted_arr.insert(j, arr[i])
                inserted = True
                break
        if not inserted:
            sorted_arr.append(arr[i])
    return sorted_arr

def main():
    nums = [i for i in range(100)]
    for i in range(5):
        arr = []
        for i in range(10):
            arr.append(random.choice(nums))
        print(insert_sort(arr))

if __name__ == '__main__':
    main()