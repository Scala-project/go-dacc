#!/usr/bin/python
# -*- coding: UTF-8 -*-
import urllib2
import json
url = ' http://localhost:8615/v1/admin/nodeinfo'
response = urllib2.urlopen(url)
info = response.read().decode()
jd = json.loads(info)
print(jd['result']['id'])