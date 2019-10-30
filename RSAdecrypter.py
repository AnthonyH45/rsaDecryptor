#!/usr/bin/python3
import math
'''
Written by Anthony Hallak for CS111 HW2
Given N (or p&q) and e, this program will
attempt to crack it and find d to decrypt
the messge
No license given, feel free to take, change, 
and spread as you would like

KNOWN ISSUES:
    - not a lot of santizing user input and 
        breaks when types are mismatched
    - requires the message to be a str pasted 
        into after it finds d
    - have not tested it with large values, use 
        with caution
    - havnet tested with lowercase or symbols
        - Assumes the user can paste in the message after 
        it finds a value for d (which it might not if
        n is stupid large, again I haven't tested much)
    - breaks if e > n
    - probably more, haven't done a lot of testing, 
        this was made to do some hw for me lol

'''
def find_pq(n: int):
    for i in range(2, int(math.sqrt(n))):
        if n % i == 0:
            p = i
            q = n / i
            print("\tFOUND P :", p, "\n\tFOUND Q :", q)
            return p, q;
    return 1, n

'''
taken from
https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
thanks to them for the mod_inverse function!
I DID NOT WRITE THIS FUNCTION, I DO NOT CLAIM ANYTHING

EDIT: October 29, 2019
Danial Beg showed us that we could improve this function
as Prof. Elena suggested by changing our if statement to the following
if (((phi * i)+1) % e) == 0
'''
def mod_inv(a: int, m: int):
# Iterative Python 3 program to find 
# modular inverse using extended 
# Euclid algorithm 
  
# Returns modulo inverse of a with 
# respect to m using extended Euclid 
# Algorithm Assumption: a and m are 
# coprimes, i.e., gcd(a, m) = 1 
    m0 = m 
    y = 0
    x = 1
  
    if (m == 1) : 
        return 0
  
    while (a > 1) : 
  
        # q is quotient 
        q = a // m 
  
        t = m 
  
        # m is remainder now, process 
        # same as Euclid's algo 
        m = a % m 
        a = t 
        t = y 
  
        # Update x and y 
        y = x - q * y 
        x = t 
  
  
    # Make x positive 
    if (x < 0) : 
        x = x + m0 
  
    return x 
  
  
# This code is contributed by Nikita tiwari. 
'''
    for i in range(1, int(phi)):
        if e * i % phi - 1 == 0:
            return i
    return 1
'''

def decrypt(d: int, n: int):
    alpha = {2: 'A', 3: 'B', 4: 'C', 5: 'D', 6: 'E', 7: 'F', 8: 'G', 9: 'H', 10: 'I', 11: 'J', 12: 'K', 13: 'L', 14: 'M', 15: 'N', 16: 'O', 17: 'P', 18: 'Q', 19: 'R', 20: 'S', 21: 'T', 22: 'U', 23: 'V', 24: 'W', 25: 'X', 26: 'Y', 27: 'Z', 28: ' ' }
    msg = str(input("Paste message : "))
    Splits = msg.split( )
    to_print = ""
    for s in Splits:
        print(str(s))
        # s^d mod n = dmsg
        # s is what we want to decrypt, d is the key
        dmsg = (int(s) ** d) % (n)
        print(str(dmsg))
        to_print += str(alpha.get(dmsg))
    print(to_print)

def find_d(p, q, e):
    print("\nUsing",p,",",q,", &",e,"to find d via the following formula: de = 1 mod((p-1)(q-1))")
    phi = (p-1) * (q-1)
    # de = 1 mod((p-1)(q-1))
    d = mod_inv(e, phi)
    print("(",p-1,")(",q-1,") : ", phi)
    print("d = ", d)
    decrypt(d, (p*q))    
    
def main():
    print("This program assumes that A->2,B->3,...,Z->27,\' \'->28\nIf this is not the case, please modify the program accordingly")
    print("This program also requires you to paste the message in after \'d\' has been found")
    know_pq = str(input("Do you know p&q? (y/n) "))
    if know_pq == "y" or know_pq == "Y" or know_pq == "yes" or know_pq == "Yes":
        p = int(input("Enter p : "))
        q = int(input("Enter q : "))
    else:
        n = int(input("Enter n : "))
        p, q = find_pq(n)
    e = int(input("Enter e : "))
    find_d(p,q,e)

main()

