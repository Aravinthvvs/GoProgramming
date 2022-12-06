
def patterns(n):
    asciiNumber = 65
    for i in range(0,n):
        for j in range(i+1):
            character = chr(asciiNumber)
            print(character,end=" ")
            asciiNumber +=1
        print("")

def wordPattern():
    word = "Python"
    out =""
    for i in range(len(word)):
        out += word[i]
        print(out)  

def triangle(n):
    for i in range(n,0,-1):
        c = n-(i-1)
        print((n-cw)*" ",end="")
        for j in range(c):
           print("*",end="")
        print("")

if __name__ == "__main__":
    triangle(5)