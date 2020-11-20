#!/usr/bin/env python3
# -*- coding: utf-8 -*-
if __name__=="__main__":
    f = open('my_test.txt', "r")
    f.readlines()
    total_count = 0
    total_sum = 0
    for line in f.readlines():
        total_count+=1
        print(line)
        total_sum += line
    print("finish reading")
    print("avg is " + total_sum/total_count)
    f.close
