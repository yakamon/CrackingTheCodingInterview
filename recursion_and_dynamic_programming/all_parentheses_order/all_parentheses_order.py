from typing import List, Tuple, Set, Dict, DefaultDict
from sys import stdin

input = stdin.readline


def all_parentheses_order(n: int) -> Set[str]:
    result = set()
    if n == 0:
        result.add("")
        return result
    prev = all_parentheses_order(n - 1)
    for s in prev:
        for i in range(len(s)):
            if s[i] == "(":
                result.add(s[: i + 1] + "()" + s[i + 1 :])
        result.add("()" + s)
    return result


def all_parentheses_order(n: int) -> Set[str]:
    def add_parentheses(result: List[str], left: int, right: int, s: str, index: int):
        if left < 0 or right < left:
            return
        if left == 0 and right == 0:
            result.append(s)
        else:
            s = s[:index] + '(' + s[index+1:]
            add_parentheses(result, left - 1, right, s, index + 1)

            s = s[:index] + ')' + s[index+1:]
            add_parentheses(result, left, right - 1, s, index + 1)

    result = []
    add_parentheses(result, n, n, "", 0)
    return result


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        n = int(input())
        ans = all_parentheses_order(n)
        print(ans)
