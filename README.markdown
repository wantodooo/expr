Basic expression evaluation using go/parser.


Implemented operations:

 - +x: nop
 - -x: arithmetic negate
 - ^x: bitwise not
 - x+y: sum
 - x-y: difference
 - x*y: product
 - x/y: quotient
 - x%y: remainder of division
 - x^y: bitwise exclusive or
 - x&y: bitwise and
 - x|y: bitwise inclusive or
 - x&^y: bit clear
 - x<<y: bitwise shift left
 - x>>y: bitwise (arithmetic) shift right
 - abs(x): absolute value
 - binomial(x, y): binomial coefficient (x choose y)
 - lb(x): position of highest set bit
 - pow(x, y[, m]): exponentiation base x of y, optionally mod m
 - gcd(x, y): greatest common denominator
 - modinv(x, y): multiplicative inverse of x in the group Z/yZ where y is assumed prime
 - factorial(x=1, y): y!/x!
 - rand(x=0, y): random number in [x, y)

