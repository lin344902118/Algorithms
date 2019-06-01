# -*- encoding:utf-8 -*-
"""
    author: nick
    date: 2019/6/1
"""

import random
import unittest

from python.selectSort import select_sort
from python.quickSort import quick_sort
from python.bubbleSort import bubble_sort

def get_random_arr():
    nums = [i for i in range(100)]
    test_arr = []
    for i in range(10):
        test_arr.append(random.choice(nums))
    return test_arr


class SortTest(unittest.TestCase):

    def sort_test(self, sorted_arr: list):
        for i in range(len(sorted_arr)-1):
            for j in range(i+1, len(sorted_arr)):
                self.assertLessEqual(sorted_arr[i], sorted_arr[j])

    def test_bubble_sort(self):
        test_arr = get_random_arr()
        sorted_arr = bubble_sort(test_arr)
        print(sorted_arr)
        self.sort_test(sorted_arr)

    def test_quick_sort(self):
        test_arr = get_random_arr()
        sorted_arr = quick_sort(test_arr)
        print(sorted_arr)
        self.sort_test(sorted_arr)

    def test_select_sort(self):
        test_arr = get_random_arr()
        sorted_arr = select_sort(test_arr)
        print(sorted_arr)
        self.sort_test(sorted_arr)


if __name__ == '__main__':
    unittest.main()