from collections import defaultdict, deque
from sys import stdin
from typing import DefaultDict, Dict, List, Set, Tuple

input = stdin.readline


def make_coin_change(n: int) -> int:
    def make_coin_change_inner(amount: int, coins: List[int], index: int) -> int:
        if index >= len(coins) - 1:
            return 1
        coinAmount = coins[index]
        ways = 0
        for i in range(amount // coinAmount + 1):
            amountRemaining = amount - i * coinAmount
            ways += make_coin_change_inner(amountRemaining, coins, index + 1)
        return ways

    coins = [25, 10, 5, 1]
    return make_coin_change_inner(n, coins, 0)


def make_coin_change(amount: int) -> int:
    def inner(
        amount: int, coins: List[int], index: int, m: DefaultDict[Tuple[int], int]
    ) -> int:
        if m[(amount, index)] > 0:
            return m[(amount, index)]
        if index >= len(coins) - 1:
            return 1
        coinAmount = coins[index]
        m[(amount, index)] = sum(
            [
                inner(amount - i * coinAmount, coins, index + 1, m)
                for i in range(amount // coinAmount + 1)
            ]
        )
        return m[(amount, index)]

    return inner(amount, [25, 10, 5, 1], 0, defaultdict(int))


if __name__ == "__main__":
    N = int(input())
    for i in range(N):
        n = int(input())
        answer = make_coin_change(n)
        print(answer)
