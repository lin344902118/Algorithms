# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/1
"""

import random

def bubble_sort(arr: list):
    for i in range(len(arr)-1):
        for j in range(i+1, len(arr)):
            if arr[i] > arr[j]:
                arr[i], arr[j] = arr[j], arr[i]
    return arr

def main():
    nums = [i for i in range(100)]
    for i in range(5):
        arr = []
        for i in range(10):
            arr.append(random.choice(nums))
        print(bubble_sort(arr))

if __name__ == '__main__':
    main()