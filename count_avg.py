#!/usr/bin/env python3
# -*- coding: utf-8 -*-
if __name__=="__main__":
    f = open('my_test.txt', "r")
    total_count = 0
    total_sum = 0
    for line in f.readlines():
        print(line)
        if total_count >0 :
            total_sum += float(line)
        total_count+=1
    print("finish reading")
    print("avg is " + str(total_sum/total_count))
    f.close
