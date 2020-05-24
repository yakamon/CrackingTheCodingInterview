from collections import defaultdict, deque
from sys import stdin
from typing import DefaultDict, Dict, List, Set, Tuple

input = stdin.readline


def place_queens(n: int) -> List[List[int]]:
    def inner(n: int, cols: List[int], row: int, results: List[List[int]]):
        if row == n:
            results.append(cols.copy())
            return
        for col in range(n):
            if is_valid(cols, row, col):
                cols[row] = col
                inner(n, cols, row + 1, results)
        return

    def is_valid(cols: List[int], row: int, col: int) -> bool:
        for r in range(row):
            c = cols[r]
            if c == col:
                return False
            if abs(col - c) == row - r:
                return False
        return True

    results = []
    inner(n, [0] * n, 0, results)
    return results


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        n = int(input())
        ans = place_queens(n)

        print(f"Case{i}:")
        for i, a in enumerate(ans):
            print(f"Answer{i}:")
            for row in a:
                s = [" "] * n
                s[row] = "q"
                print(s)
