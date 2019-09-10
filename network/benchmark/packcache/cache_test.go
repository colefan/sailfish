package packcache

import (
	"testing"

	"sync"

	"github.com/colefan/sailfish/network"
)

type testLoop struct {
	queue *network.PackQueue
}

func newTestLoop() *testLoop {
	t := &testLoop{}
	t.queue = network.NewPackQueue(4096)
	return t

}

func (t *testLoop) loop() {
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	go func() {
		for i := 0; i < 1000; i++ {
			item := new(network.Message)
			item.SetCmd(int32(i))
			cache := new(network.PackCache)
			cache.Pack = item
			t.queue.Push(cache)
			//wg.Add(1)
		}
	}()

	go func() {
		for {
			items := t.queue.PopList(100)
			for _, item := range items {
				item.Pack = nil
				wg.Done()
			}
		}
	}()

	wg.Wait()

}

func BenchmarkPackQueue(b *testing.B) {
	q := new(network.PackQueue)
	for i := 0; i < b.N; i++ {
		pack := new(network.Message)
		pack.SetCmd(int32(i))
		cache := new(network.PackCache)
		cache.Pack = pack
		q.Push(cache)
		q.PopList(64)
	}

}

func BenchmarkPackQueueP(b *testing.B) {
	q := new(network.PackQueue)
	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			pack := new(network.Message)
			cache := new(network.PackCache)
			cache.Pack = pack
			q.Push(cache)
			q.PopList(100)

		}
	})
}
