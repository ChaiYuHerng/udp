if __name__=="__main__":
    f = open('my_test.txt','r')
    for line in f.readlines():
        print(line)
    f.close