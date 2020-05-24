def intmul(a: int, b: int) -> int:
    def intmul_innter(a: int, b: int, s: int) -> int:
        if b == 0:
            return 0
        s += intmul_innter(a, b-1, s)
        return s+a
    return intmul_innter(a, b, 0)


def intmul2(a: int, b: int) -> int:
    s = 0
    for _ in range(b, 0, -1):
        s += a
    return s


def intmul3(a: int, b: int) -> int:
    def intmul_inner(a: int, b: int) -> int:
        for _ in range(b, 0, -1):
            yield a
    return sum(intmul_inner(a, b))


def intmul4(a: int, b: int) -> int:
    def intmul_inner(smaller: int, bigger: int) -> int:
        if smaller == 0:
            return 0
        if smaller == 1:
            return bigger
        q = smaller >> 1
        side = intmul_inner(q, bigger)
        if smaller % 2 == 1:
            return side + side + bigger
        else:
            return side + side
    smaller, bigger = (a, b) if a < b else (b, a)
    return intmul_inner(smaller, bigger)


for i in range(10):
    print("intmul: ", intmul(2, i))
    print("intmul2:", intmul2(20, i))
    print("intmul3:", intmul3(200, i))
    print("intmul4:", intmul4(2000, i))
