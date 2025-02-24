from decimal import Decimal, localcontext

def prob80():
    total = 0
    for x in range(1,100):
        with localcontext() as ctx:
            ctx.prec = 105
            if len(str(Decimal(x).sqrt())) == 1:
                total+=0
            else:
                a = sum([int(i) for i in str(Decimal(x).sqrt())[2:101]])+int(str(Decimal(x).sqrt())[0])
                total+=a
    return total
print(prob80())