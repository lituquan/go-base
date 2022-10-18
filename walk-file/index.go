package main

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

// index is an inverted index. It maps tokens to document IDs.
type index struct {
	contain *cache.Cache
}

var locks sync.Mutex

func New() index {
	// 创建缓存
	var codeIndex = cache.New(1000*time.Hour, 1000*time.Hour)
	return index{codeIndex}
}

// add adds documents to the index.
func (idx index) add(doc File) {
	locks.Lock()
	defer locks.Unlock()
	for _, token := range analyze(doc.Content) {
		ids, ok := idx.contain.Get(token)
		v := []int{}
		if ok {
			if v = ids.([]int); v[len(v)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
		}
		v = append(v, doc.ID)
		idx.contain.Set(token, v, cache.NoExpiration)
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
		ids, ok := idx.contain.Get(token)
		v := []int{}
		if ok {
			if v, ok = ids.([]int); ok {
				if r == nil {
					r = v
				} else {
					r = intersection(r, v)
				}
			} else {
				// Token doesn't exist.
				return nil
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

func (idx index) loadFromMap(kv map[string][]int) {
	for k, v := range kv {
		codeIndex.contain.Set(k, v, cache.NoExpiration)
	}
}
