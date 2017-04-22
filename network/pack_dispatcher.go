package network

//PackDispatcherInf 消息分发接口
type PackDispatcherInf interface {
	PostData(pack PackInf, session *TCPSession)
	FetchData() *PackCache
	FetchDataList(size int) []*PackCache
	FetchAllData() []*PackCache
	Dispose()
}

//PackDispatcher cls
type PackDispatcher struct {
	queue *PackQueue
}

//NewPackDispatcher new
func NewPackDispatcher() *PackDispatcher {
	return &PackDispatcher{queue: NewPackQueue(4096)}
}

//PostData post data to queue
func (d *PackDispatcher) PostData(pack PackInf, session *TCPSession) {
	item := &PackCache{Pack: pack, Session: session}
	d.queue.Push(item)
}

//FetchDataList 获取数据
func (d *PackDispatcher) FetchDataList(size int) []*PackCache {
	items := d.queue.PopList(size)
	return items
}

//FetchAllData fetch all
func (d *PackDispatcher) FetchAllData() []*PackCache {
	return d.queue.PopAll()
}

//FetchData fetch data
func (d *PackDispatcher) FetchData() *PackCache {
	return d.queue.Pop()
}

//Dispose disponse dispatcher data
func (d *PackDispatcher) Dispose() {
	d.queue.Dispose()
}
