package config
type Limithd_zuiqiangzhenrong struct{
	Items []*Limithd_zuiqiangzhenrong_Item
}

type Limithd_zuiqiangzhenrong_Item struct{
	Id int
	Number int
	Star int
	Rewards [][]int
}

var InstLimithd_zuiqiangzhenrong *Limithd_zuiqiangzhenrong

func init(){
	item := &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20001
	item.Number = 3
	item.Star = 2
	item.Rewards = [][]int{{2,2,2000,},{201,405001,2,},{201,406002,2}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
	item = &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20002
	item.Number = 5
	item.Star = 2
	item.Rewards = [][]int{{2,2,3000,},{201,405001,2,},{201,406002,2}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
	item = &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20003
	item.Number = 1
	item.Star = 3
	item.Rewards = [][]int{{2,2,5000,},{201,405001,3,},{201,406002,2}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
	item = &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20004
	item.Number = 2
	item.Star = 3
	item.Rewards = [][]int{{2,2,10000,},{201,405001,3,},{201,406002,2}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
	item = &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20005
	item.Number = 3
	item.Star = 3
	item.Rewards = [][]int{{2,2,15000,},{201,405002,2,},{201,310017,10}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
	item = &Limithd_zuiqiangzhenrong_Item{}
	item.Id = 20006
	item.Number = 1
	item.Star = 4
	item.Rewards = [][]int{{2,2,20000,},{201,310017,10,},{201,406003,2}}
	InstLimithd_zuiqiangzhenrong.Items = append(InstLimithd_zuiqiangzhenrong.Items,item)
}
