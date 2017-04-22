package config
type Yl_limit struct{
	Items []*Yl_limit_Item
}

type Yl_limit_Item struct{
	Id int
	Level int
	Type int
	Time int
	Vip int
	Rmb int
	Rewards [][]int
}

var InstYl_limit *Yl_limit

func init(){
	item := &Yl_limit_Item{}
	item.Id = 1
	item.Level = 30
	item.Type = 1
	item.Time = 86400
	item.Vip = 0
	item.Rmb = 30
	item.Rewards = [][]int{{201,431001,1},{201,406002,1},{2,2,2000}}
	InstYl_limit.Items = append(InstYl_limit.Items,item)
	item = &Yl_limit_Item{}
	item.Id = 2
	item.Level = 30
	item.Type = 1
	item.Time = 86400
	item.Vip = 0
	item.Rmb = 98
	item.Rewards = [][]int{{201,431001,1},{201,406002,1},{2,2,2000}}
	InstYl_limit.Items = append(InstYl_limit.Items,item)
	item = &Yl_limit_Item{}
	item.Id = 3
	item.Level = 55
	item.Type = 1
	item.Time = 86400
	item.Vip = 0
	item.Rmb = 30
	item.Rewards = [][]int{{201,431001,1},{201,406002,1},{2,2,2000}}
	InstYl_limit.Items = append(InstYl_limit.Items,item)
	item = &Yl_limit_Item{}
	item.Id = 4
	item.Level = 55
	item.Type = 1
	item.Time = 86400
	item.Vip = 0
	item.Rmb = 98
	item.Rewards = [][]int{{201,431001,1},{201,406002,1},{2,2,2000}}
	InstYl_limit.Items = append(InstYl_limit.Items,item)
}
