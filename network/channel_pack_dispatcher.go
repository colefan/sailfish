package network

const (
	defaultChannelPackBuffer = 40960
)

//ChannelPackDispatcher dispatcher
type ChannelPackDispatcher struct {
	packBuffer chan *PackCache
	bufferSize int
}

//NewChannelPackDispatcher creator
func NewChannelPackDispatcher(size int) *ChannelPackDispatcher {
	d := &ChannelPackDispatcher{}
	if size <= 1020 {
		size = defaultChannelPackBuffer
	}
	d.bufferSize = size
	d.packBuffer = make(chan *PackCache, d.bufferSize)
	return d
}

//PostData post
func (d *ChannelPackDispatcher) PostData(pack PackInf, session *TCPSession) {
	item := &PackCache{Pack: pack, Session: session}
	d.packBuffer <- item
}

func (d *ChannelPackDispatcher) FetchData() *PackCache {
	item := <-d.packBuffer
	return item
}

func (d *ChannelPackDispatcher) FetchDataList(size int) []*PackCache {
	return nil
}

func (d *ChannelPackDispatcher) FetchAllData() []*PackCache {
	return nil
}

func (d *ChannelPackDispatcher) Dispose() {
	close(d.packBuffer)
}

func (c *ChannelPackDispatcher) IsChannelDispatcher() bool {
	return true
}
