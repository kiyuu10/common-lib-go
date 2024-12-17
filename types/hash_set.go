package types

import (
	"encoding/json"
)

type HashSet[T comparable] map[T]struct{}

func NewHashSet[T comparable](items ...T) HashSet[T] {
	return NewHashSetFromList(items)
}

func NewHashSetFromList[T comparable](items []T) HashSet[T] {
	hashSet := make(HashSet[T], len(items))
	for i := 0; i < len(items); i++ {
		hashSet.Add(items[i])
	}
	return hashSet
}

func (s HashSet[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}
	s[v] = struct{}{}
	return true
}

func (s HashSet[T]) Remove(v T) bool {
	if !s.Contains(v) {
		return false
	}
	delete(s, v)
	return true
}

func (s HashSet[T]) Contains(v T) bool {
	_, exists := s[v]
	return exists
}

func (s HashSet[T]) AsList() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s HashSet[T]) Intersect(that HashSet[T]) HashSet[T] {
	commonSet := make(HashSet[T])
	for item := range s {
		if that.Contains(item) {
			commonSet.Add(item)
		}
	}
	for item := range that {
		if s.Contains(item) {
			commonSet.Add(item)
		}
	}
	return commonSet
}

func (s HashSet[T]) Union(that HashSet[T]) HashSet[T] {
	maxLen := len(s)
	if maxLen < len(that) {
		maxLen = len(that)
	}
	unionSet := make(HashSet[T], maxLen)
	for item := range s {
		unionSet.Add(item)
	}
	for item := range that {
		unionSet.Add(item)
	}
	return unionSet
}

func (s HashSet[T]) Diff(that HashSet[T]) HashSet[T] {
	diffSet := make(HashSet[T])
	for item := range s {
		if !that.Contains(item) {
			diffSet.Add(item)
		}
	}
	return diffSet
}

func (s HashSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.AsList())
}

func (s *HashSet[T]) UnmarshalJSON(data []byte) error {
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	*s = make(HashSet[T], len(items))
	for _, item := range items {
		s.Add(item)
	}

	return nil
}
