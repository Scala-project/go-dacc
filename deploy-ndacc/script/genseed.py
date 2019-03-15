#!/usr/bin/env python
# coding=utf-8

output = open('config/seeds.json', 'w')
nodes = []
for line in open("temp/nodes.txt","r"):
	info = line.replace('\n', '').replace('\r', '')
	n = '\"'+info+'\"'
	nodes.append(n)
ret = ','.join(nodes)
output.write('[')
output.write(ret)
output.write(']')
output.close()