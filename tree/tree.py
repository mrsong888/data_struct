#!/usr/bin/env python
# *-* coding:utf-8 *-*
import json
from collections import defaultdict, OrderedDict

class TreeNode(object):

    def __init__(self, height=None):
        self._height = height
        self._root = defaultdict(list)
        self._tmp = self._root
        self._floor = 0

    def insert(self, key, val):
        if key == None:
            self._floor += 1
            return
        if self._floor + 1 == self._height:
            if isinstance(self._tmp, dict):
                self._tmp[key].append(val)
            else:
                self._tmp.append(val)
            self._floor += 1
        else:
            tmp = {}
            tmp[val] = []
            if isinstance(self._tmp, list):
                self._tmp.append(tmp)
                self._tmp = tmp[val]
            else:
                self._tmp[key].append(tmp)
                self._tmp = tmp
            self._floor += 1

    def search(self, key):
        if isinstance(self._tmp, dict):
            if key in self._tmp and len(self._tmp[key]) > 0:
                self._tmp = self._tmp[key]
                self._floor += 1
                return True
            else:
                return False
        elif isinstance(self._tmp, list):
            for d in self._tmp:
                if isinstance(d, dict):
                    if key in d:
                        self._tmp = d[key]
                        self._floor += 1
                        return True
                else:
                    if d == key:
                        self._floor += 1
                        return True
            return False

    def delete(self, l):
        d = OrderedDict()
        for key in l:
            if self.search(key):
                d[key] = len(self._tmp)
        t = d.items()
        if len(t) == len(l):
            #node leaf
            if t[-2][1] > 1:
                self._tmp.remove(t[-1][0])
                return True
            else:
                print t
                index = None
                for i, k in enumerate(t):
                    if k[1] > 1 and i != len(t)-1:
                        index = i
                print "gggggggg"
                print index
                if index:
                    self.reload()
                    r_d = None
                    for key in t[:index+1]:
                        self.search(key[0])
                    print t[:index+1]
                    print self._tmp
                    if isinstance(self._tmp, list):
                        for d in self._tmp:
                            if t[index+1][0] in d:
                                r_d = d
                        print "r_d", r_d
                        if r_d:
                            self._tmp.remove(r_d)

                else:
                    # all node 1 and node leaf 1
                    self._root.pop(l[0], None)

    @property
    def floor(self):
        return self._floor

    @property
    def root(self):
        return self._root

    def reload(self):
        self._tmp = self._root
        self._floor = 0

if __name__ == "__main__":
    l = [1, 2, 4, 5]
    n = TreeNode(height=len(l))
    pre = None
    for i in l:
        n.insert(pre, i)
        pre = i
    n.reload()
    print n.root
    l1 = [1, 2, 4, 6]
    pre = None
    for j in l1:
        if n.search(j):
            pre = j
        else:
            n.insert(pre, j)
            pre = j

    # print json.dumps(dict(n.root))
    n.reload()
    l2 = [1, 2, 5, 6]
    for j in l2:
        if n.search(j):
            pre = j
        else:
            n.insert(pre, j)
            pre = j

    n.reload()
    l3 = l2
    # for j in l3:
    #     n.search(j)
    #     print len(n._tmp)
    #     print n._tmp
    print n.delete(l3)
    print n.root
    print n._floor