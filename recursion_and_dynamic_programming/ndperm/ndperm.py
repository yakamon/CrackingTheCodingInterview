from typing import List
from sys import stdin
input = stdin.readline


def find_permutations(s: str) -> List[str]:
    return [l for l in find_permutations_inner(list(s))]


def find_permutations_inner(chars: List[chr]):
    if len(chars) <= 1:
        return chars
    perms = []
    for i in range(len(chars)):
        child_perms = find_permutations_inner(chars[:i] + chars[i+1:])
        for j in range(len(child_perms)):
            perms.append(chars[i] + child_perms[j])
    return perms


def find_permutations(s: str) -> List[str]:
    permutations = []
    if len(s) == 0:
        permutations.append('')
        return permutations
    first = s[0]
    remainder = s[1:]
    words = find_permutations(remainder)
    for w in words:
        for i in range(len(w)+1):
            permutations.append(w[:i] + first + w[i:])
    return permutations


def find_permutations(s: str) -> List[str]:
    length = len(s)
    permutations = []
    if length == 0:
        permutations.append('')
        return permutations

    for i in range(length):
        before = s[:i]
        after = s[i+1:]
        partials = find_permutations(before + after)
        for w in partials:
            permutations.append(s[i] + w)
    return permutations


if __name__ == '__main__':
    n = int(input())
    for i in range(n):
        perms = find_permutations(input().strip())
        print(len(perms), perms)
