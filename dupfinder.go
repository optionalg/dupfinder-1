package main

import "container/vector"

type DuplicateMap map[string] *vector.StringVector

func FindDuplicates(p Path) DuplicateMap {
	duplicates := make(DuplicateMap)
	
	for fd := range ChecksumIterator(p).Iter() {
		vec, ok := duplicates[fd.Hash]
		if !ok {
			vec = new(vector.StringVector)
			duplicates[fd.Hash] = vec
		}
		vec.Push(fd.Name)
	}
	return duplicates
}
