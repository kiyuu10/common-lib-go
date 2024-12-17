package types

import (
	"encoding/json"
)

type Slice[T comparable] []T

func NewSlice[T comparable](items ...T) Slice[T] {
	return NewSliceFromList(items)
}

func NewSliceFromList[T comparable](items []T) Slice[T] {
	slice := make(Slice[T], 0)
	for _, v := range items {
		slice.Add(v)
	}
	return slice
}

func (s *Slice[T]) Add(v T) {
	*s = append(*s, v)
}

func (s *Slice[T]) Remove(v T) bool {
	items := *s
	if !s.Contains(v) {
		return false
	}
	for i, _v := range items {
		if _v == v {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	*s = items
	return true
}

func (s Slice[T]) Contains(v T) bool {
	for _, _v := range s {
		if _v == v {
			return true
		}
	}
	return false
}

func (s Slice[T]) Intersect(that Slice[T]) Slice[T] {
	commonSlice := make(Slice[T], 0)
	for _, item := range s {
		if that.Contains(item) {
			commonSlice.Add(item)
		}
	}
	for _, item := range that {
		if s.Contains(item) {
			commonSlice.Add(item)
		}
	}
	return commonSlice
}

func (s Slice[T]) Union(that Slice[T]) Slice[T] {
	unionSlice := make(Slice[T], 0)
	for _, item := range s {
		unionSlice.Add(item)
	}
	for _, item := range that {
		unionSlice.Add(item)
	}
	return unionSlice
}

func (s Slice[T]) Diff(that Slice[T]) Slice[T] {
	diffSlice := make(Slice[T], 0)
	for _, item := range s {
		if !that.Contains(item) {
			diffSlice.Add(item)
		}
	}
	return diffSlice
}

func (s Slice[T]) Unique() Slice[T] {
	uniqueSlice := make([]T, 0, len(s))
	seen := make(map[T]bool, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			uniqueSlice = append(uniqueSlice, v)
			seen[v] = true
		}
	}
	return uniqueSlice
}

func (s Slice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Slice[T]) UnmarshalJSON(data []byte) error {
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	*s = make(Slice[T], 0)
	for _, item := range items {
		s.Add(item)
	}

	return nil
}
