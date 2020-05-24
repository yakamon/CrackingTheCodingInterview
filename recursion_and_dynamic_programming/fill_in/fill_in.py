from collections import deque
from sys import stdin
from typing import DefaultDict, Dict, List, Set, Tuple

input = stdin.readline


def fill_in(screen: List[List[int]], c: int, r: int, new_color: int) -> bool:
    """
    Depth-First Search
    """

    def fill_in_inner(
        screen: List[List[int]], c: int, r: int, old_col: int, new_col: int
    ) -> bool:
        if c < 0 or len(screen[0]) - 1 < c or r < 0 or len(screen) - 1 < r:
            return False
        if screen[r][c] == old_col:
            screen[r][c] = new_col
            fill_in_inner(screen, c - 1, r, old_col, new_col)
            fill_in_inner(screen, c + 1, r, old_col, new_col)
            fill_in_inner(screen, c, r - 1, old_col, new_col)
            fill_in_inner(screen, c, r + 1, old_col, new_col)
        return True

    if screen[r][c] == new_color:
        return False
    return fill_in_inner(screen, c, r, screen[r][c], new_color)


def fill_in(screen: List[List[int]], c: int, r: int, new_col: int) -> bool:
    """
    Breadth-First Search
    """

    def is_valid_loc(c: int, r: int):
        return 0 <= r and r < len(screen) and 0 <= c and c < len(screen[0])

    def adjacents(dot: Tuple[int]) -> List[Tuple[int]]:
        c, r = dot
        ret = []
        return [(c - 1, r), (c + 1, r), (c, r - 1), (c, r + 1)]

    if screen[r][c] == new_col:
        return False

    old_col = screen[r][c]
    q = deque()
    if is_valid_loc(c, r):
        screen[r][c]
    q.append((c, r))
    while len(q) > 0:
        root = q.popleft()
        for c, r in adjacents(root):
            if is_valid_loc(c, r) and screen[r][c] == old_col:
                screen[r][c] = new_col
                q.append((c, r))
    return True


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        H, W, = [int(v) for v in input().strip().split()]
        x, y, = [int(v) for v in input().strip().split()]
        color = int(input())
        print(H, W, x, y, color)
        screen = [] * H
        for r in range(H):
            screen.append([int(c) for c in input().strip().split(" ")])

        [print(row) for row in screen]

        ans = fill_in(screen, x, y, color)
        print(ans)

        [print(row) for row in screen]
