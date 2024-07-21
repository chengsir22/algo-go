package main

import "sort"

func groupAnagrams(strs []string) (res [][]string) {
	mp := make(map[string][]string)
	for _, str := range strs {
		t := []byte(str)
		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		k := string(t)
		mp[k] = append(mp[k], str)
	}
	for _, v := range mp {
		res = append(res, v)
	}
	return
}
