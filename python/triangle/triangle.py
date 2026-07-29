def equilateral(sides):
    a, b, c = sides
    if a == 0 and b == 0  and c == 0:
        return False
    if a == b and a == c:
        return True
    return False

def verify_triangle(a,b,c):
    return a + b >= c and b + c >= a and a + c >= b

def isosceles(sides):
    a, b, c = sides
    if verify_triangle(a,b,c):
        if (a == b or a == c) or (b == a or b == c) or (c == b or c == a):
            return True
    return False

def scalene(sides):
    a, b, c = sides
    if verify_triangle(a,b,c):
        if a != b and a != c and c != b:
            return True
    return False

