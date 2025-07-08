package typeSyncMap

type ISyncMap[K any, V comparable] interface {
	Load(key K) (value V, ok bool)
	Store(key K, value V)
	Clear()
	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
	Delete(key K)
	CompareAndSwap(key K, old V, new V) (swapped bool)
	CompareAndDelete(key K, old V) (deleted bool)
	Range(f func(key K, value V) bool)
}
