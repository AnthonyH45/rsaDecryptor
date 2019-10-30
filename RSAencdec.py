#!/usr/bin/python3
import math

'''
taken from
https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
thanks to them for the mod_inverse function!
I DID NOT WRITE THIS FUNCTION, I DO NOT CLAIM ANYTHING
'''
def mod_inv(e: int, phi: int):
    e = e % phi;
    for i in range(1, int(phi)):
        # make sure that e does not share any divisors with phi
        if e * i % phi == 1:
            return i
    return 1

'''
From mathexchange, Gauss's algorithm:
a^-1 == a^(phi-1) mod n
'''
def mod_inv(e: int, phi: int, n: int):
    return ((e ** (phi-1)) % n)

def find_pq(n: int):
    for i in range(2, int(math.sqrt(n))):
        if n % i == 0:
            p = i
            q = n / i
            print("\tFOUND P :", p, "\n\tFOUND Q :", q)
            return p, q;
    return 1, n

def decrypt_file(d: int, n: int, file_name: str):
    alpha = {2: 'A', 3: 'B', 4: 'C', 5: 'D', 6: 'E', 7: 'F', 8: 'G', 9: 'H', 10: 'I', 11: 'J', 12: 'K', 13: 'L', 14: 'M', 15: 'N', 16: 'O', 17: 'P', 18: 'Q', 19: 'R', 20: 'S', 21: 'T', 22: 'U', 23: 'V', 24: 'W', 25: 'X', 26: 'Y', 27: 'Z', 28: ' ' }
    msg = open(file_name, 'r')
    msg = str(msg.read())
    Splits = msg.split( )
    to_print = ""
    for s in Splits:
        #print(str(s))
        # s^d mod n = dmsg
        # s is what we want to decrypt, d is the key
        dmsg = (int(s) ** d) % (n)
        #print("dmsg: ", str(dmsg))
        to_print += str(alpha.get(dmsg))
    print(to_print)
    file_name += "_decrypted"
    to_write = open(file_name, "w+")
    to_write.write(to_print)
    print("Written to the file name you gave but added \'_decrypted\' to the end of it")

def encrypt_file(e: int, n: int, file_name: str):
    alpha = {2: 'A', 3: 'B', 4: 'C', 5: 'D', 6: 'E', 7: 'F', 8: 'G', 9: 'H', 10: 'I', 11: 'J', 12: 'K', 13: 'L', 14: 'M', 15: 'N', 16: 'O', 17: 'P', 18: 'Q', 19: 'R', 20: 'S', 21: 'T', 22: 'U', 23: 'V', 24: 'W', 25: 'X', 26: 'Y', 27: 'Z', 28: ' ' }
    # one way to invert the key/values in the dict,
    alpha = {v: k for k, v in alpha.items()}
    # another way,  dict(alpha(reversed, alpha.items()))
    msg = open(file_name, 'r')
    msg = str(msg.read())
    to_print = ""
    char_msg = list(msg)
    for s in char_msg:
        # print(alpha.get(s.upper()))
        emsg = ((alpha.get(s.upper())) ** e) % (n)
        to_print += (str(emsg) + " ")
    print(to_print)
    file_name += "_encrypted"
    to_write = open(file_name, "w+")
    to_write.write(to_print)
    print("Written to the file name you gave but added \'_encrypted\' to the end of it")

def find_d(p, q, e):
    print("\nUsing",p,",",q,", &",e,"to find d via the following formula: de = 1 mod((p-1)(q-1))")
    phi = (p-1) * (q-1)
    # de = 1 mod((p-1)(q-1))
    d = mod_inv(e, phi)
    print("(",p-1,")(",q-1,") : ", phi)
    print("d = ", d)
    return d #decrypt(d, (p*q))

def main():
    print("This program assumes that A->2,B->3,...,Z->27,\' \'->28\nIf this is not the case, please modify the program accordingly")
    user_choice = str(input("Do you want to (e)ncrypt or (d)ecrypt? "))
    if user_choice == "e" or user_choice == "E":
        n_or_pq = str(input("Do you have n (y/n)? "))
        if n_or_pq == "y" or n_or_pq == "Y":
            n = int(input("Please enter n: "))
            p,q = find_pq(n)
        else:
            p = int(input("Please enter p: "))
            q = int(input("Please enter q: "))
            n = p * q
        e = int(input("Please enter e: "))
        file_name = str(input("Please enter the name of the file you would like to encrypt (must be in same directory as this script): "))
        encrypt_file(e, n, file_name)
    else:
        n_or_pq = str(input("Do you have n (y/n)? "))
        if n_or_pq == "y" or n_or_pq == "Y":
            n = int(input("Please enter n: "))
            p,q = find_pq(n)
        else:
            p = int(input("Please enter p: "))
            q = int(input("Please enter q: "))
        e = int(input("Please enter e: "))
        d = find_d(p, q, e)
        file_name = str(input("Please enter the name of the file you would like to decrypt (must be in same directory as this script): "))
        decrypt_file(d, n, file_name)

main()
