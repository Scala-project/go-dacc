#!/usr/bin/env python
# coding=utf-8

output = open('config/static-nodes.json', 'w')
nodes = []
for line in open("temp/static-nodes.txt","r"):
	info = line.replace('\n', '').replace('\r', '')
	s1 = info.split(' ')
	s2 = s1[0].split('[::]')
	n = '  \"'+s2[0] + s1[1] + s2[1] +'\"'
	nodes.append(n)	
	print n
ret = ',\n'.join(nodes)
output.write('[\n')
output.write(ret)
output.write('\n]')
output.close()	