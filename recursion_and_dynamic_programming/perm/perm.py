from typing import List, DefaultDict
from sys import stdin
from collections import defaultdict
import time
input = stdin.readline


def find_permutations(s: str) -> List[str]:
    if len(s) <= 1:
        return s
    charmap = 0
    permutations = []
    for i in range(len(s)):
        c = s[i]
        charmap_key = 1 << (ord(c)-ord('a'))
        if charmap_key & charmap != 0:
            continue
        charmap |= charmap_key
        remainder = s[:i] + s[i+1:]
        words = find_permutations(remainder)
        permutations.extend([c + w for w in words])
    return permutations


def find_permutations(s: str) -> List[str]:
    def make_char_freq_map(s: str) -> DefaultDict[chr, int]:
        m = defaultdict(int)
        for c in s:
            m[c] += 1
        return m
    permutations = []
    m = make_char_freq_map(s)
    find_permutations_inner(m, '', len(s), permutations)
    return permutations


def find_permutations_inner(m: DefaultDict[chr, int], prefix: str, remaining: int, results: List[str]):
    if remaining == 0:
        results.append(prefix)
        return
    for c, cnt in m.items():
        if cnt > 0:
            m[c] -= 1
            find_permutations_inner(m, prefix + c, remaining-1, results)
            m[c] += 1


if __name__ == '__main__':
    times = []
    n = int(input())
    for i in range(n):
        s = input().strip()
        start = time.time()
        permutations = find_permutations(s)
        end = time.time()
        times.append(end-start)
        print(len(permutations), permutations)
    print(times)
