package partitions

import (
	"testing"
)

var sinkErr error

func BenchmarkSetInt_UniqueKeys(b *testing.B) {
	p := &Partition{
		Schema:  INT,
		IntData: make(map[string]int64, b.N),
	}

	keys := makeByteKeys(b.N)
	val := []byte("123")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkErr = p.Set(keys[i], val)
	}
}

func BenchmarkSetInt_HotKeys(b *testing.B) {
	const hot = 1024 // power of two

	p := &Partition{
		Schema:  INT,
		IntData: make(map[string]int64, hot),
	}

	keys := makeHotByteKeys(hot)
	val := []byte("123")

	// warm: pre-fill so we measure overwrite more than growth
	for i := 0; i < hot; i++ {
		_ = p.Set(keys[i], val)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkErr = p.Set(keys[i&(hot-1)], val)
	}
}

func BenchmarkSetString_UniqueKeys(b *testing.B) {
	p := &Partition{
		Schema:     STRING,
		StringData: make(map[string]string, b.N),
	}

	keys := makeByteKeys(b.N)
	val := []byte("hello")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkErr = p.Set(keys[i], val)
	}
}

func BenchmarkSetString_HotKeys(b *testing.B) {
	const hot = 1024

	p := &Partition{
		Schema:     STRING,
		StringData: make(map[string]string, hot),
	}

	keys := makeHotByteKeys(hot)
	val := []byte("hello")

	for i := 0; i < hot; i++ {
		_ = p.Set(keys[i], val)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkErr = p.Set(keys[i&(hot-1)], val)
	}
}

var sinkI64 int64

func BenchmarkBulkDel_Int(b *testing.B) {
	const (
		total = 1 << 16
		bulk  = 128
	)

	p := &Partition{
		Schema:  INT,
		IntData: make(map[string]int64, total),
	}

	keysAll := makeByteKeys(total)
	val := []byte("123")

	// fill all keys once
	for i := 0; i < total; i++ {
		_ = p.Set(keysAll[i], val)
	}

	// bulk window keys (reuse)
	bulkKeys := keysAll[:bulk]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// delete same bulk set (after first iter they're missing)
		// If you want steady behavior, re-insert before deleting (more realistic but adds cost).
		sinkI64 = p.BulkDel(bulkKeys)
	}
}

func BenchmarkExists_Int(b *testing.B) {
	const (
		total = 1 << 16
		bulk  = 128
	)

	p := &Partition{
		Schema:  INT,
		IntData: make(map[string]int64, total),
	}

	keysAll := makeByteKeys(total)
	val := []byte("123")

	// insert half so EXISTS returns ~64/128
	for i := 0; i < total; i += 2 {
		_ = p.Set(keysAll[i], val)
	}

	bulkKeys := keysAll[:bulk]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkI64 = p.Exists(bulkKeys)
	}
}

func BenchmarkSetInt_HotKeys_Parallel(b *testing.B) {
	const hot = 1024
	p := &Partition{Schema: INT, IntData: make(map[string]int64, hot)}
	keys := makeHotByteKeys(hot)
	val := []byte("123")

	for i := 0; i < hot; i++ {
		_ = p.Set(keys[i], val)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = p.Set(keys[i&(hot-1)], val)
			i++
		}
	})
}
