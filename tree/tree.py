#!/usr/bin/env python
# *-* coding:utf-8 *-*
import json
from collections import defaultdict, OrderedDict

class TreeNode(object):
    """
    search tree
    """

    def __init__(self, height=None):
        self._height = height
        self._root = defaultdict(list)
        self._tmp = self._root
        self._floor = 0

    def _insert(self, key, val):
        """
        insert operate
        :param key:
        :param val:
        :return:
        """
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

    def insert(self, l):
        """
        signle insert
        :param l: list
        :return:
        """
        pre = None
        for j in l:
            if n.search(j):
                pre = j
            else:
                n._insert(pre, j)
                pre = j
        self._reload()

    def batchinsert(self, l):
        """
        batch insert
        :param l: list:list
        :return:
        """
        for j in l:
            self.insert(j)

    def search(self, key):
        """
        key contains in tree or not
        :param key:
        :return: bool
        """
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
        """
        delete node
        :param l: list
        :return: bool
        """
        try:
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
                    index = None
                    for i, k in enumerate(t):
                        if k[1] > 1 and i != len(t)-1:
                            index = i
                    if index:
                        self._reload()
                        r_d = None
                        for key in t[:index+1]:
                            self.search(key[0])
                        if isinstance(self._tmp, list):
                            for d in self._tmp:
                                if t[index+1][0] in d:
                                    r_d = d
                            if r_d:
                                self._tmp.remove(r_d)
                                return True

                    else:
                        # all node 1 and node leaf 1
                        self._root.pop(l[0], None)
                        return True
            else:
                return False
        finally:
            self._reload()

    def modify(self, orign, new):
        """

        :param orign: origin load
        :param new: modify new load
        :return: bool
        """
        #first: delete origin load second: insert new load
        if self.delete(orign):
            self.insert(new)

    @property
    def floor(self):
        return self._floor

    @property
    def root(self):
        return self._root

    def _reload(self):
        self._tmp = self._root
        self._floor = 0

if __name__ == "__main__":
    l = [1, 2, 4, 5]
    n = TreeNode(height=len(l))
    l1 = [1, 2, 4, 6]
    l2 = [1, 2, 5, 6]
    all = [l, l2]
    n.batchinsert(all)
    n.modify(l1, l2)
    # l3 = l2
    # print n.delete(l1)
    print n.root
    print n._floor