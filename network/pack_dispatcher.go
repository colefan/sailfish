package network

//PackDispatcher cls
type PackDispatcher struct {
	queue *PackQueue
}

func newPackDispatcher() *PackDispatcher {
	return &PackDispatcher{queue: NewPackQueue(4096)}
}

//PostData post data to queue
func (d *PackDispatcher) PostData(pack PackInf, session *TCPSession) {
	item := &PackCache{Pack: pack, Session: session}
	d.queue.Push(item)
}

//FetchDataList 获取数据
func (d *PackDispatcher) FetchDataList() []*PackCache {
	items := d.queue.PopList(100)
	return items
}

//FetchData fetch data
func (d *PackDispatcher) FetchData() *PackCache {
	return d.queue.Pop()
}

//Dispose disponse dispatcher data
func (d *PackDispatcher) Dispose() {
	d.queue.Dispose()
}
