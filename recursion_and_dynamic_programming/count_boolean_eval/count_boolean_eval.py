from collections import defaultdict, deque
from sys import stdin
from typing import DefaultDict, Dict, List, Set, Tuple

input = stdin.readline


def count_boolean_eval(exp: str, ret: bool) -> int:
    if len(exp) == 0:
        return 0
    if len(exp) == 1:
        return bool(exp)

    ways = 0
    for i in range(1, len(exp), 2):
        c = exp[i]
        left = exp[:i]
        right = exp[i + 1 :]

        left_true = count_boolean_eval(left, True)
        left_false = count_boolean_eval(left, False)
        right_true = count_boolean_eval(right, True)
        right_false = count_boolean_eval(right, False)
        total = (left_true + left_false) * (right_false + right_true)

        total_true = 0
        if c == "^":
            total_true = left_true * right_false + left_false * right_true
        elif c == "&":
            total_true = left_true * right_true
        elif c == "|":
            total_true = (
                left_true * right_true
                + left_false * right_true
                + left_true * right_false
            )

        sub_ways = total_true if ret else total - total_true
        ways += sub_ways

    return ways


def count_boolean_eval(exp: str, ret: bool) -> int:
    def inner(exp: str, ret: bool, cache: Dict[Tuple[str, bool], int]) -> int:
        if len(exp) == 0:
            return 0
        if len(exp) == 1:
            return bool(exp)
        if (exp, ret) in cache:
            return cache[(exp, ret)]

        ways = 0
        for i in range(1, len(exp), 2):
            c = exp[i]
            left = exp[:i]
            right = exp[i + 1 :]

            lt = inner(left, True, cache)
            lf = inner(left, False, cache)
            rt = inner(right, True, cache)
            rf = inner(right, False, cache)
            total = (lt + lf) * (rt + rf)

            total_true = 0
            if c == "^":
                total_true = lt * rf + lf * rt
            elif c == "&":
                total_true = lt * rt
            elif c == "|":
                total_true = lt * rt + lt * rf + lf * rt

            sub = total_true if ret else total - total_true
            ways += sub
        cache[(exp, ret)] = ways
        return cache[(exp, ret)]

    return inner(exp, ret, {})


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        row = input().strip().split(" ")
        exp = row[0]
        ret = True if row[1] == "true" else False
        ans = count_boolean_eval(exp, ret)
        print(ans)
