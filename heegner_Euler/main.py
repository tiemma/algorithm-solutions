import math
from decimal import *

# https://projecteuler.net/problem=heegner
# https://docs.python.org/3/library/decimal.html#decimal-recipes

getcontext().prec = 100

def cos(x):
    """Return the cosine of x as measured in radians.

    The Taylor series approximation works best for a small value of x.
    For larger values, first compute x = x % (2 * pi).

    >>> print(cos(Decimal('0.5')))
    0.8775825618903727161162815826
    >>> print(cos(0.5))
    0.87758256189
    >>> print(cos(0.5+0j))
    (0.87758256189+0j)

    """
    getcontext().prec += 2
    i, lasts, s, fact, num, sign = 0, 0, 1, 1, 1, 1
    while s != lasts:
        lasts = s
        i += 2
        fact *= i * (i-1)
        num *= x * x
        sign *= -1
        s += num / fact * sign
    getcontext().prec -= 2
    return +s

def pi():
    """Compute Pi to the current precision.

    >>> print(pi())
    3.141592653589793238462643383

    """
    getcontext().prec += 2  # extra digits for intermediate steps
    three = Decimal(3)      # substitute "three=3.0" for regular floats
    lasts, t, s, n, na, d, da = 0, three, 3, 1, 0, 0, 24
    while s != lasts:
        lasts = s
        n, na = n+na, na+8
        d, da = d+da, da+32
        t = (t * n) / d
        s += t
    getcontext().prec -= 2
    return +s               # unary plus applies the new precision

def exp(x):
    """Return e raised to the power of x.  Result type matches input type.

    >>> print(exp(Decimal(1)))
    2.718281828459045235360287471
    >>> print(exp(Decimal(2)))
    7.389056098930650227230427461
    >>> print(exp(2.0))
    7.38905609893
    >>> print(exp(2+0j))
    (7.38905609893+0j)

    """
    getcontext().prec += 2
    i, lasts, s, fact, num = 0, 0, 1, 1, 1
    while s != lasts:
        lasts = s
        i += 1
        fact *= i
        num *= x
        s += num / fact
    getcontext().prec -= 2
    return +s

def cosh(x):
   return ( x.exp() + (-x).exp() ) / 2

lowestDelta = 1
answer = 0

for i in range(-1000, 1001):
    if math.sqrt(abs(i)).is_integer():
        continue

    val = cos(Decimal(abs(i)).sqrt() * pi())
    if i < 0:
        # cosh(ix) = cos(x) and cos(ix) = cosh(-x) = cosh(x)
        # cos(-x) = cos(x) so the negative case is the same as the positive case
        # In the negative case of the range, we need to use cosh instead of cos
        # cos(x) = (e^ix + e-ix) / 2 where i = sqrt(-1) following euler's formula
        # cosh(ix) = (e^x + e-x) / 2 once the i powers are squared giving us cos(x)
        # Since the sqrt of n being negative gives i in the expression cos(pi * sqrt(n)), this can be simplified using cosh(x) instead of cos(ix)
        val = cosh(Decimal(abs(i)).sqrt() * pi())
    delta = val - math.floor(val)
    delta  = min(delta, 1-delta)
    print(i, delta)

    if delta == 0:
        continue

    if delta < lowestDelta:
        lowestDelta = min(delta, 1-delta)
        answer = i

print("----")
print(answer, lowestDelta)