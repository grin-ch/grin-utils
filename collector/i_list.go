package collector

type List[T comparable] interface {
	Collector[T]
	Iterator[T]

	// Insert 在指定位置 i 插入元素, i 是该元素插入后的索引
	// i<0时与i=0等效, i> list.Len() 时, 与 i=list.Len() 等效
	Insert(T, int)

	// Delete 删除指定位置 i 的元素, T 是被删除的元素
	// 当 i>=list.Len() 或 i < 0 时, 不会删除元素, 此时放回该类型零值
	Delete(int) T
}
