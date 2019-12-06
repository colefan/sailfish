package network

import "sync"

var packPool = sync.Pool{
	New: func() interface{} {
		pack := NewMessage()
		return pack
	},
}

// GetPooledPack get pooled pack
func GetPooledPack() PackInf {
	return NewMessage()
	//return packPool.Get().(*Message)
}

// FreePack free package data
func FreePack(pack PackInf) {
	// if pack != nil {
	// 	pack.Reset()
	// 	packPool.Put(pack)
	// }
}

func GetPackPool() *sync.Pool {
	return &packPool
}

func SetPackPool(pool sync.Pool) {
	packPool = pool
}
