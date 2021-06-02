package lru

type LRUCache struct {
	Head, Tail *Node
	Capacity   int
	Cache      map[int]*Node
}

type Node struct {
	Prev, Next *Node
	Key, Val   int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Prev = tail, head
	return LRUCache{
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
		Cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.Cache[key]
	if !ok {
		return -1
	}
	n := &Node{
		Prev: nil,
		Next: nil,
		Key:  key,
		Val:  res.Val,
	}
	this.remove(res)
	this.insert(n)
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	// check if key exists, if so, remove and insert
	if n, ok := this.Cache[key]; ok {
		this.remove(n)
	}
	// check if cache is at capacity, if so remove head
	if len(this.Cache) == this.Capacity {
		this.remove(this.Head.Next)
	}

	n := &Node{
		Prev: nil,
		Next: nil,
		Key:  key,
		Val:  value,
	}

	this.insert(n)
}

func (this *LRUCache) insert(n *Node) {
	this.Cache[n.Key] = n
	// insert at tail
	this.Tail.Prev.Next = n
	n.Prev = this.Tail.Prev
	n.Next = this.Tail
	this.Tail.Prev = n
}

func (this *LRUCache) remove(n *Node) {
	// remove from map
	delete(this.Cache, n.Key)
	// remove from list
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
