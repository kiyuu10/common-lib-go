package types

import (
	"fmt"
	"sync"
)

type Singleton[T any] struct {
	mux      sync.Mutex
	isLoaded bool
	instance T
	loader   func() T
}

func NewSingleton[T any](loader func() T) *Singleton[T] {
	return &Singleton[T]{
		isLoaded: false,
		loader:   loader,
	}
}

func (s *Singleton[T]) Get() T {
	if s.isLoaded {
		return s.instance
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	if !s.isLoaded {
		s.instance = s.loader()
		s.isLoaded = true
	}
	return s.instance
}

type SingletonMap[T any] struct {
	mux         sync.Mutex
	instanceMap map[string]T
	loader      func(key fmt.Stringer) T
}

func NewSingletonMap[T any](loader func(key fmt.Stringer) T) *SingletonMap[T] {
	return &SingletonMap[T]{
		loader:      loader,
		instanceMap: make(map[string]T),
	}
}

func (s *SingletonMap[T]) Get(key fmt.Stringer) T {
	textKey := key.String()
	if value, ok := s.instanceMap[textKey]; ok {
		return value
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	value, ok := s.instanceMap[textKey]
	if !ok {
		value = s.loader(key)
		s.instanceMap[textKey] = value
	}
	return value
}
