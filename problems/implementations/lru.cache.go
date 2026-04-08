package implementations

import "fmt"

type Node struct {
	Next     *Node
	Prev     *Node
	Key, Val int
}

func NewNode(key, val int) *Node {
	return &Node{
		Next: nil,
		Prev: nil,
		Key:  key,
		Val:  val,
	}
}
func (head *Node) AddNode(newNode *Node) {
	// you have to insert this node after head node
	temp := head.Next
	newNode.Next = temp
	newNode.Prev = head
	temp.Prev = newNode
	head.Next = newNode
}
func (head *Node) DeleteNode(node *Node) {
	tempPrev := node.Prev
	tempNext := node.Next
	tempPrev.Next = tempNext
	tempNext.Prev = tempPrev
}

type LRUCache struct {
	Head     *Node
	Tail     *Node
	HashMap  map[int]*Node
	Capacity int
}

func NewLruCacheWithCapacity(capacity int) *LRUCache {
	head := NewNode(-1, -1)
	tail := NewNode(-1, -1)
	head.Next = tail
	tail.Prev = head
	return &LRUCache{
		Capacity: capacity,
		HashMap:  make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
}
func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.HashMap[key]; ok == false {
		fmt.Println("key not found")
		return -1
	}
	resNode := cache.HashMap[key]
	fmt.Println("key found", resNode.Val)
	res := resNode.Val
	cache.Head.DeleteNode(resNode)
	cache.Head.AddNode(resNode)
	cache.HashMap[key] = cache.Head.Next
	return res
}

func (cache *LRUCache) Put(key, val int) {
	if _, ok := cache.HashMap[key]; ok {
		existing := cache.HashMap[key]
		delete(cache.HashMap, key)
		cache.Head.DeleteNode(existing)
	}
	if len(cache.HashMap) == cache.Capacity {
		delete(cache.HashMap, cache.Tail.Prev.Key)
		cache.Head.DeleteNode(cache.Tail.Prev)
	}
	newnode := NewNode(key, val)
	cache.Head.AddNode(newnode)
	cache.HashMap[key] = cache.Head.Next

}
