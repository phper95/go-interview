package main

import (
	"log"
	"sync"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// 集合结构体
type Set struct {
	m            map[int]struct{} // 用字典来实现，因为字段键不能重复
	sync.RWMutex                  // 锁，实现并发安全
}

// 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
}

// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()

	// 集合没元素直接返回
	if s.IsEmpty() {
		return
	}

	delete(s.m, item) // 实际从字典删除这个键
}

// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.m)
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
}

// 集合是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.Len())
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	// 空结构体不占空间
	a := struct{}{}
	b := struct{}{}
	log.Printf("a.addr:%p;b.addr:%p", &a, &b)
	if a == b {
		log.Println("equal")
	}
	log.Println(unsafe.Sizeof(a))
	return
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)
	log.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		log.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		log.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	log.Println("list of all items", s.List())
}
