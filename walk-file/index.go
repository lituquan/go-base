package main

import "sync"

// index is an inverted index. It maps tokens to document IDs.
type index map[string][]int

var locks sync.Mutex

// add adds documents to the index.
func (idx index) add(doc File) {
	locks.Lock()
	defer locks.Unlock()
	for _, token := range analyze(doc.Content) {
		ids := idx[token]
		if ids != nil && ids[len(ids)-1] == doc.ID {
			// Don't add same ID twice.
			continue
		}
		idx[token] = append(ids, doc.ID)
	}
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// search queries the index for the given text.
func (idx index) search(text string) []File {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	var fileSearch = []File{}
	for _, v := range r {
		file := File{files[v].ID, files[v].Path, ""}
		//file := files[v]
		fileSearch = append(fileSearch, file)
		if len(fileSearch) > 20 {
			break
		}
	}
	return fileSearch
}
