def solvePell(D):
    t = D ** 0.5     # square root
    if (a := int(t)) == t: return 0     # D is square, no solution
    b, c, r = a, 1, a << 1
    e1, e2 = 1, 0
    f1, f2 = 0, 1
    while True:
        b = r * c - b
        c = (D - b * b) // c
        r = (a + b) // c
        print(b,c,r)
        e1, e2 = e2, e1 + e2 * r
        f1, f2 = f2, f1 + f2 * r
        x, y = f2 * a + e2, f2
        print(x, y, x * x - D * y * y)
        if x * x - D * y * y == 1: return x

print(max(range(61, 62), key=solvePell))  # 661