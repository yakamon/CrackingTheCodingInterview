from typing import *


def power_set(s: List[int]) -> List[List[int]]:
    return power_set_inner(s, 0, [])


def power_set_inner(s: List[int], index: int, memo: List[List[int]]) -> List[List[int]]:
    if index == len(s):
        return memo + [[]]
    n = s[index]
    memo = power_set_inner(s, index+1, memo)
    size = len(memo)
    for i in range(0, size):
        memo += [memo[i] + [n]]
    return memo


def power_set2(set: List[int]) -> List[List[int]]:
    subsets = []
    max_num = 1 << len(set)
    for k in range(max_num):
        subset = [set[i] for i in range(len(set)) if (k & (1 << i)) != 0]
        subsets.append(subset)
    return subsets


if __name__ == '__main__':
    for i in range(6):
        print(power_set(range(i)))
        print(power_set2(range(i)))
