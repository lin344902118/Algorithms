# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/9
"""

import random

def merge(arr1: list, arr2: list):
    tmp = []
    i, j = 0, 0
    while i < len(arr1) and j < len(arr2):
        if arr1[i] < arr2[j]:
            tmp.append(arr1[i])
            i += 1
        else:
            tmp.append(arr2[j])
            j += 1
    if i < len(arr1):
        tmp.extend(arr1[i:])
    if j < len(arr2):
        tmp.extend(arr2[j:])
    return tmp


def merge_sort(arr: list):
    if len(arr) <= 2:
        if len(arr) == 1:
            return arr
        if arr[0] > arr[1]:
            arr[0], arr[1] = arr[1], arr[0]
        return arr
    mid = len(arr) // 2
    if len(arr) % 2 != 0:
        mid += 1
    font = merge_sort(arr[:mid])
    after = merge_sort(arr[mid:])
    merged_arr = merge(font, after)
    return merged_arr

def main():
    nums = [i for i in range(100)]
    for i in range(5):
        arr = []
        for i in range(10):
            arr.append(random.choice(nums))
        print(merge_sort(arr))

if __name__ == '__main__':
    main()