package main

import "log"

//队列、数组
//map

type Lru struct {
	Map  map[int]*Node
	Head *Node
	Tail *Node
	Max  int
}
type Node struct {
	Data int
	Key  int
	Prev *Node
	Next *Node
}

func New(max int) *Lru {
	lru := Lru{
		Max: max,
		Map: make(map[int]*Node),
	}
	return &lru
}
func (l *Lru) Get(key int) int {
	if v, ok := l.Map[key]; ok {
		//退化:断开当前，移到head
		if v == l.Head {

		} else if v == l.Tail {
			v.Prev.Next = nil
			l.Tail = v.Prev

			v.Prev = nil
			v.Next = l.Head
			l.Head.Prev = v
			l.Head = v
		} else {
			v.Prev.Next = v.Next
			v.Next.Prev = v.Prev

			v.Prev = nil
			v.Next = l.Head
			l.Head.Prev = v
			l.Head = v
		}
		//获取值
		return v.Data
	}
	return -1
}
func (l *Lru) Put(key int, value int) {
	if v, ok := l.Map[key]; ok {
		v.Data = value
		//退化:断开当前，移到head
		l.Get(key)
	} else {
		//数据淘汰判断
		if l.Max == len(l.Map) {
			prev := l.Tail.Prev
			if prev != nil {
				prev.Next = nil
				l.Tail.Prev = nil
				delete(l.Map, l.Tail.Key)
				l.Tail = prev
			} else {
				delete(l.Map, l.Tail.Key)
				l.Tail = nil
				l.Head = nil
			}
		}
		//新建节点
		node := &Node{
			Key:  key,
			Data: value,
			Next: l.Head,
			Prev: nil,
		}
		//节点加入
		if l.Head == nil {
			l.Head = node
			l.Tail = node
		} else {
			l.Head.Prev = node
			l.Head = node
		}
		l.Map[key] = node
	}
}
func (l *Lru) Remove(key int) {

}
func (l *Lru) All() {
	log.Println("===========")
	cur := l.Head
	for ; cur != nil; cur = cur.Next {
		log.Println(cur.Key, ":", cur.Data)
	}
}
func main() {
	l := New(5)
	for i := 1; i <= 5; i++ {
		l.Put(i, 'a'+i-1)
	}
	l.All()
	l.Get(4)
	l.All()
	l.Put(6, 'f')
	l.All()
}
