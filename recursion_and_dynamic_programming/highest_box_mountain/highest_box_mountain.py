from __future__ import annotations

from collections import defaultdict, deque
from sys import stdin
from typing import DefaultDict, Dict, List, Set, Tuple

input = stdin.readline


class Box:
    def __init__(self, width: int, height: int, depth: int):
        self.width = width
        self.height = height
        self.depth = depth

    def can_be_above(self, box: Box):
        return (
            self.width < box.width
            and self.height < box.height
            and self.depth < box.depth
        )


def max_box_mountain_height(boxes: List[Box]) -> int:
    def inner(boxes: List[Box], bottomIndex: int) -> int:
        bottom = boxes[bottomIndex]
        max_height = 0
        for i in range(bottomIndex + 1, len(boxes)):
            if boxes[i].can_be_above(bottom):
                max_height = max(max_height, inner(boxes, i))
        max_height += bottom.height
        return max_height

    boxes.sort(key=lambda x: x.width, reverse=True)
    max_height = 0
    for i in range(len(boxes)):
        max_height = max(max_height, inner(boxes, i))
    return max_height


def max_box_mountain_height(boxes: List[Box]) -> int:
    def inner(boxes: List[Box], bottomIndex: int, max_heights: Dict[int, int]) -> int:
        if bottomIndex < len(boxes) and max_heights[bottomIndex] > 0:
            return max_heights[bottomIndex]

        bottom = boxes[bottomIndex]
        max_height = 0
        for i in range(bottomIndex + 1, len(boxes)):
            if boxes[i].can_be_above(bottom):
                max_height = max(max_height, inner(boxes, i, max_heights))
        max_height += bottom.height
        max_heights[bottomIndex] = max_height
        return max_height

    boxes.sort(key=lambda x: x.height, reverse=True)
    max_height = 0
    max_heights = defaultdict(int)
    for i in range(len(boxes)):
        max_height = max(max_height, inner(boxes, i, max_heights))
    return max_height


def max_box_mountain_height(boxes: List[Box]) -> int:
    def inner(
        boxes: List[Box], current_bottom: Box, offset: int, max_heights: Dict[int, int],
    ) -> int:
        if offset >= len(boxes):
            return 0

        # height with bottom
        new_bottom = boxes[offset]
        height_with_bottom = 0
        if current_bottom is None or new_bottom.can_be_above(current_bottom):
            if max_heights[offset] == 0:
                max_heights[offset] = inner(boxes, new_bottom, offset + 1, max_heights)
                max_heights[offset] += new_bottom.height
            height_with_bottom = max_heights[offset]

        # height_without_bottom
        height_without_bottom = inner(boxes, current_bottom, offset + 1, max_heights)

        # compare with and without
        return max(height_with_bottom, height_without_bottom)

    boxes.sort(key=lambda x: x.height, reverse=True)
    max_heights = defaultdict(int)
    return inner(boxes, None, 0, max_heights)


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        print(input().strip())

        n = int(input())
        boxes = []
        for i in range(n):
            w, h, d = (int(v) for v in input().strip().split())
            boxes.append(Box(w, h, d))

        input()

        ans = max_box_mountain_height(boxes)
        print(ans)
