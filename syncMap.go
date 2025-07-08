package typeSyncMap

import "sync"

type TypeSyncMap[K any, V comparable] struct {
	data sync.Map
}

func (t *TypeSyncMap[K, V]) Load(key K) (value V, ok bool) {
	val, ok := t.data.Load(key)
	if !ok {
		return value, false
	}
	value, ok = val.(V)
	return value, ok
}

func (t *TypeSyncMap[K, V]) Store(key K, value V) {
	t.data.Store(key, value)
}

func (t *TypeSyncMap[K, V]) Clear() {
	t.data.Clear()
}

func (t *TypeSyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	val, loaded := t.data.LoadOrStore(key, value)
	if loaded {
		actual, _ = val.(V)
	} else {
		actual = value
	}
	return actual, loaded
}

func (t *TypeSyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	val, loaded := t.data.LoadAndDelete(key)
	if !loaded {
		return value, false
	}
	value, _ = val.(V)
	return value, true
}

func (t *TypeSyncMap[K, V]) Delete(key K) {
	t.data.Delete(key)
}

func (t *TypeSyncMap[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) {
	val, loaded := t.data.Load(key)
	if !loaded {
		return false
	}
	value, _ := val.(V)
	if value != old {
		return false
	}
	t.data.Store(key, new)
	return true
}

func (t *TypeSyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	val, loaded := t.data.Load(key)
	if !loaded {
		return false
	}
	value, _ := val.(V)
	if value != old {
		return false
	}
	t.data.Delete(key)
	return true
}

func (t *TypeSyncMap[K, V]) Range(f func(key K, value V) bool) {
	t.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func NewTypeSyncMap[K any, V comparable]() *TypeSyncMap[K, V] {
	return &TypeSyncMap[K, V]{
		data: sync.Map{},
	}
}
