#!/usr/bin/env python
# coding=utf-8

output = open('config/seeds.json', 'w')
seeds = []
for line in open("temp/seeds.txt","r"):
	info = line.replace('\n', '').replace('\r', '')
	n = '\"'+info+'\"'
	seeds.append(n)
ret = ','.join(seeds)
output.write('[')
output.write(ret)
output.write(']')
output.close()