package config
type Yl_wish struct{
	Items []*Yl_wish_Item
}

type Yl_wish_Item struct{
	Id int
	Day int
	Choose int
	Vip int
	Times int
	Rewards [][]int
}

var InstYl_wish *Yl_wish

func init(){
	item := &Yl_wish_Item{}
	item.Id = 110001
	item.Day = 1
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,405002,5},{201,405001,10},{201,431001,5},{201,406002,5},{2,2,50000},{3,3,20000},{201,432003,1}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
	item = &Yl_wish_Item{}
	item.Id = 110002
	item.Day = 2
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,406003,3},{201,405001,10},{201,431001,5},{201,406002,10},{2,2,50000},{3,3,20000},{201,432003,1}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
	item = &Yl_wish_Item{}
	item.Id = 110003
	item.Day = 3
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,405002,5},{201,406002,10},{201,431001,5},{201,435001,5},{2,2,50000},{3,3,20000},{201,432003,1}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
	item = &Yl_wish_Item{}
	item.Id = 110004
	item.Day = 4
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,405002,5},{201,406003,5},{201,431001,5},{201,435001,5},{2,2,50000},{3,3,20000},{201,310042,5}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
	item = &Yl_wish_Item{}
	item.Id = 110005
	item.Day = 5
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,405002,5},{201,405001,10},{201,435002,5},{201,435001,5},{2,2,50000},{3,3,20000},{201,310016,5}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
	item = &Yl_wish_Item{}
	item.Id = 110006
	item.Day = 6
	item.Choose = 3
	item.Vip = 0
	item.Times = 1
	item.Rewards = [][]int{{1,1,100},{2,2,20000},{3,3,50000},{201,405002,5},{201,405001,10},{201,435002,5},{201,406003,5},{2,2,50000},{3,3,20000},{201,310049,5}}
	InstYl_wish.Items = append(InstYl_wish.Items,item)
}
