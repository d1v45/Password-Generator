######################################################################################################
# Title: PASSWORD GENERATOR                                                                          #
# Author: DIVAS A S                                                                                  #
# Github : https://github.com/d1v45                                                                  #
######################################################################################################

import random
import string

def banner():
    print('''\n
\t     ▒░█▀▀▄ █▀▀█ █▀▀▀ █▀▀▀   █▀▀█ █▀▀▀ █▄  █░▒       
\t     ▒░█▄▄▀░█▄▄█ ▀▀▀█ ▀▀▀█   █ ▄▄ █▀▀░ █▀█▄█░▒       
\t     ▒░█░▒  ▀░░▀ ▀░░▀ ▀░░▀   █▄▄█ ▀░░▀ ▀░░░▀░▒       
\n''')
    

def pass_gen():
    upper = string.ascii_uppercase
    lower = string.ascii_lowercase
    num = string.digits
    symbol = string.punctuation

    all = upper + lower + num + symbol
    #print(all)
    
    length = int(input("Enter the length of the password: "))
    amount = int(input("Enter the number of passwords you need: "))

    for x in range(amount):
        password = "".join(random.sample(all, length))
        print(password)

if __name__ == '__main__':
    banner()
    pass_gen()